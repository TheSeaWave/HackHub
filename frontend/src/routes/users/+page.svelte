<script lang="ts">
    import { onMount } from 'svelte';

    // Интерфейс для пользователя
    interface User {
        id: number;
        firstName: string;
        lastName: string;
        rating: number;
        role: string;
        techStack: string[];
        telegram: string;
        about: string;
        achievements: string;
        isFound: boolean;
    }

    let usersLookingForTeam: User[] = [];  // Массив пользователей, ищущих команду
    let selectedUser: User | null = null;  // Выбранный пользователь для поп-апа
    let searchQuery: string = '';  // Строка поиска
    let searchResults: User[] = [];  // Результаты поиска

    // Функция для загрузки пользователей, которые ищут команду
    async function loadUsers() {
        const response = await fetch('/api/auth/users-looking-for-team');
        if (response.ok) {
            usersLookingForTeam = await response.json();
        } else {
            alert('Не удалось загрузить анкеты участников');
        }
    }

    // Функция для поиска пользователей по запросу
    function searchUsers() {
        if (searchQuery.trim() === '') {
            searchResults = [];
        } else {
            searchResults = usersLookingForTeam.filter(user =>
                user.firstName.toLowerCase().includes(searchQuery.toLowerCase()) ||
                user.lastName.toLowerCase().includes(searchQuery.toLowerCase()) ||
                user.id.toString().includes(searchQuery)
            );
        }
    }

    // Функция для открытия поп-апа с информацией о пользователе
    function openUserProfile(user: User) {
        selectedUser = user;
    }

    // Функция для закрытия поп-апа
    function closeUserProfile() {
        selectedUser = null;
    }

    // Функция для добавления пользователя в команду (пример)
    async function addUserToTeam(userId: number) {
        const response = await fetch(`/api/auth/add-to-team`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ userId })
        });
        if (response.ok) {
            alert('Пользователь добавлен в команду');
        } else {
            alert('Ошибка при добавлении пользователя в команду');
        }
    }

    onMount(() => {
        loadUsers();
    });
</script>


<div>
    <label for="userSearch">Поиск участников:</label>
    <input 
        type="text" 
        id="userSearch" 
        placeholder="Введите ID или имя участника" 
        bind:value={searchQuery} 
        on:input={searchUsers} 
    />
</div>

<!-- Список результатов поиска -->
<ul>
    {#each searchResults as user (user.id)}
        <li>
            <button type="button" on:click={() => addUserToTeam(user.id)}>
                {user.firstName} {user.lastName}
            </button>
        </li>
    {/each}
</ul>

<h1>Участники, ищущие команду</h1>

{#if usersLookingForTeam.length > 0}
    <div>
        {#each usersLookingForTeam as user (user.id)}
            <div class="user-card" on:click={() => openUserProfile(user)}>
                <h3>{user.firstName} {user.lastName}</h3>
                <p><strong>ID:</strong> {user.id}</p>
                <p><strong>Рейтинг:</strong> {user.rating}</p>
                <p><strong>Роль:</strong> {user.role}</p>
                <p><strong>Технологии:</strong> {user.techStack.join(', ')}</p>
            </div>
        {/each}
    </div>
{:else}
    <p>Нет участников, которые ищут команду.</p>
{/if}

{#if selectedUser}
    <div class="popup">
        <div class="popup-content">
            <h2>{selectedUser.firstName} {selectedUser.lastName}</h2>
            <p><strong>ID:</strong> {selectedUser.id}</p>
            <p><strong>Рейтинг:</strong> {selectedUser.rating}</p>
            <p><strong>Роль:</strong> {selectedUser.role}</p>
            <p><strong>Стек технологий:</strong> {selectedUser.techStack.join(', ')}</p>
            <p><strong>Telegram:</strong> {selectedUser.telegram}</p>
            <p><strong>О себе:</strong> {selectedUser.about}</p>
            <p><strong>Достижения:</strong> {selectedUser.achievements}</p>
            <button class="close-button" on:click={closeUserProfile}>Закрыть</button>
        </div>
    </div>
{/if}
