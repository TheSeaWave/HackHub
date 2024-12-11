<script lang="ts">
    import { onMount } from 'svelte';

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

    let usersLookingForTeam: User[] = [];  
    let selectedUser: User | null = null;  
    let searchQuery: string = '';  
    let searchResults: User[] = [];  

        async function loadUsers() {
        const response = await fetch('/api/auth/users-looking-for-team');
        if (response.ok) {
            usersLookingForTeam = await response.json();
        } else {
            alert('Не удалось загрузить анкеты участников');
        }
    }

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

    function openUserProfile(user: User) {
        selectedUser = user;
    }

    function closeUserProfile() {
        selectedUser = null;
    }

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


<div class=container2>
    <div  class="naming1"><a href="/">hack.itam</a></div>
    <div class="hrefComands"><a href="/comands">команды</a></div>
    <div class="hrefNews"><a href="/news">новости</a></div>
    <div class="hrefProfile"><a href="/profile">профиль</a></div>
    <div class="hrefHacks"><a href="/hacks">хакатоны</a></div>
    <div class="hrefAnkets"><a href="/users">анкеты</a></div>
    <div class="notific"><button class="notific1"></button></div>
</div>

<div class="container11">
    <div class="ankets1">АНКЕТЫ</div>
    <div class="sortimage"> </div>
<div>
    
<input class="search"
        type="text" 
        id="userSearch" 
        placeholder="Введите ID или имя участника" 
        bind:value={searchQuery} 
        on:input={searchUsers} />
</div>
</div>
<ul>
    {#each searchResults as user (user.id)}
        <li>
            <button type="button" on:click={() => addUserToTeam(user.id)}>
                {user.firstName} {user.lastName}
            </button>
        </li>
    {/each}
</ul>

<div class="ankets">  
{#if usersLookingForTeam.length > 0}
    <div>
        {#each usersLookingForTeam as user (user.id)}
            <div class="usercard" on:click={() => openUserProfile(user)}>
                <div class="mainpart"><div class="mainpart1">{user.firstName} {user.lastName} {user.telegram}</div> 
                <div class="reit1">★{user.rating}</div></div>
                <div class="role1">{user.role}</div>
                <div class="mainpart2">
                <div class="tech1"> {user.techStack.join(', ')}</div>
            <div class="idOn">ID:{user.id}</div></div>
        </div>
        {/each}
    </div>

{:else}
    <p class="nofound">Больше нет участников, которые ищут команду.</p>
{/if}
</div>
{#if selectedUser}
<div class="modal">
<div class="modal-content">
        <h2>{selectedUser.firstName} {selectedUser.lastName}</h2>
        <p><strong>ID:</strong> {selectedUser.id}</p>
        <p><strong>Рейтинг:</strong> {selectedUser.rating}</p>
        <p><strong>Роль:</strong> {selectedUser.role}</p>
        <p><strong>Стек технологий:</strong> {selectedUser.techStack.join(', ')}</p>
        <p><strong>Telegram:</strong> {selectedUser.telegram}</p>
        <p><strong>О себе:</strong> {selectedUser.about}</p>
        <p><strong>Достижения:</strong> {selectedUser.achievements}</p>
        <button class="close-button" on:click={closeUserProfile}>Закрыть</button>
</div></div>
{/if}
<div class="footer">
    <div class=containerFooter>
        <div class=footer1><p>© 2024</p></div>
        <div class=footer2><p>hack.itam</p></div>
        </div>
</div>

<style>
.usercard{
    width: 83vw;
    color: white;
    position: relative;
    min-height: 17vw;
    height: auto;
    background: rgba(0, 0, 0, 0.3);
    margin-left: 7vw;
    margin-top: 5vw;
    border-radius: 2vw;
    border: 0.3vw solid black;
}
    
.footer1{
    font-weight: 100;
    font-size: 2vw;
    margin-left: 5vw;
}
    
.footer2{
    font-weight: 300;
    font-size: 2vw;
    margin-left: auto;
    margin-right: 4vw; 
}

.footer1, .footer2{
    color:white;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    margin-bottom: 4vw;
}
    
:global(body) {
    display: flex;
    flex-direction: column;
    background-image: url("/backgr.png");
    background-repeat: repeat-y;
    background-position: center;
    background-size: 100%;
}

.mainpart, .mainpart2{
    display: flex;
    flex-direction:row;
}

.footer{
    background-color: #000000;
    width: 100%;
    margin-top: 200vw;
}
    
.containerFooter{
    display: flex;
    position: relative;
    margin-bottom: 0;
}

.ankets1{
    margin-top: 4vw;
    margin-left: 12vw;
    margin-bottom: -3vw;
    font-size: 4vw;
    font-family: "Russo One", sans-serif;
    font-weight: 400;
    font-style: normal;
    color:white;
}
    
.mainpart1{
    margin-top: 2vw;
    margin-left: 2vw;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    font-size: 2.5vw;
    font-weight: 200;
}

.tech1{
    margin-top: 1vw;
    margin-left: 2vw;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    font-size: 3.5vw;
    font-weight: 200;
}

.reit1{
    border-radius: 1vw;
    border: 0.3vw solid rgb(255, 255, 255);
    margin-left: auto;
    margin-right: 2vw;
    padding: 0.5vw;
    margin-top: 1.5vw;
    font-size: 2.5vw;
}

.idOn{
    margin-left: auto;
    margin-right: 2vw;
    padding: 0.5vw;
    margin-top: 1.5vw;
    font-size: 2vw;
    color: rgb(190, 181, 181);
}

.role1, .nofound{
    margin-top: 2vw;
    margin-left: 2vw;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    font-weight: 200;
    color: rgb(190, 181, 181);
    font-size: 2.5vw;
}

.modal {
    position: fixed;
    width: 60vw;
    top: 15vw;
    bottom: 10vw;
    left: 22vw;
    border: none;
    padding: 0;
    color: #000000;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 999;
    border-radius: 3vw;
    border: 0.3vw solid black;
    max-height: 100vw; 
    overflow-y: auto; 
}

.modal-content {
    background: rgba(0, 0, 0, 0.8);
    padding: 4vw; 
    width: 100%;
    overflow: hidden;
    max-height: 70vw;
}

.search{
    width: 30vw;
    padding: 0.2vw 0.6vw;
    font-size: 1.25vw;
    color: #212529;
    background-color: #fff;
    background-clip: padding-box;
    border: 0.2vw solid #bdbdbd;
    border-radius: 0.4vw;
    transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-weight: 100;
    font-style: normal;
    height: 2.5vw;
    margin-top: 4vw;
    margin-left: 27vw;
}

.container11{
    display: flex;
    position: relative;
    margin-bottom: 0;
}

.container2{
    display: flex;
    z-index: 2;
}

.naming1, .hrefComands, .hrefNews, .hrefProfile, .hrefHacks, .hrefAnkets{
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    margin-left: 5vw;
}

.hrefComands{
    margin-left: 5vw;
}

.hrefComands, .hrefNews, .hrefProfile, .hrefHacks, .hrefAnkets{
    font-weight: 100;
    font-size: 2.5vw;
    margin-top: 3.5vw;
}

.naming1{
    font-weight: 400;
    font-size: 3vw;
    margin-top: 3vw;
    color:white;
}

.notific{
    margin-left: auto;
    position: relative;
}

.notific1{
    background-image: url('/notific.png');
    margin-top: 3.5vw;
    margin-right: 5vw;
    width: 3vw;
    height: 3vw;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    border: 0;
    background-color: #000000;
    z-index: 2;
}
</style>