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

<header>


</header>
<body>
   <div class="container11">
    <div class="ankets1">АНКЕТЫ</div>
    <div class="sortimage"> </div>
    <div>
        <input class="search"
            type="text" 
            id="userSearch" 
            placeholder="Введите ID или имя участника" 
            bind:value={searchQuery} 
            on:input={searchUsers} 
        />
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
   </div>
            
            <div class="usercard" on:click={() => openUserProfile}>
                <div class="mainpart"><div class="mainpart1">Имя Фамилия @buklykot</div>  
                <div class="reit1">★ 2,25</div></div>
                <div class="role1">Designer</div>
                <div class="mainpart2">
                <div class="tech1">html, css, figma</div>
                <div class="idOn">ID:544156156</div></div>
            </div>
            <div class="usercard" on:click={() => openUserProfile}>
                <div class="mainpart"><div class="mainpart1">Имя Фамилия @buklykot</div>  
                <div class="reit1">★ 2,25</div></div>
                <div class="role1">Designer</div>
                <div class="mainpart2">
                <div class="tech1">html, css, figma</div>
                <div class="idOn">ID:544156156</div></div>
            </div>
            <div class="usercard" on:click={() => openUserProfile}>
                <div class="mainpart"><div class="mainpart1">Имя Фамилия @buklykot</div>  
                <div class="reit1">★ 2,25</div></div>
                <div class="role1">Designer</div>
                <div class="mainpart2">
                <div class="tech1">html, css, figma</div>
                <div class="idOn">ID:544156156</div></div>
            </div>
   

    <p class="nofound">Больше нет участников, которые ищут команду</p>

{#if selectedUser}
    <div class="modal">
        <div class="modal-content">
            <p>Имя Фамилия</p>
            <p>ID:5512151</p>
            <p><span class="star">★</span><span>2,25</span></p>
            <p>дизайнер</p>
            <p><strong>Стек технологий:</strong> {selectedUser.techStack.join(', ')}</p>
            <p><strong>Telegram:</strong> @buklykot</p>
            <p><strong>О себе:</strong> многа многа букаф</p>
            <p><strong>Достижения:</strong> победил многа хакатонаф</p>
            <button class="close-button" on:click={closeUserProfile}>Закрыть</button>
        </div>
    </div>
{/if}
</body>
<footer>
    <div class=containerFooter>
    <div class=footer1><p>© 2024</p></div>
    <div class=footer2><p>hack.itam</p></div>
    </div>
</footer>
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

body {
    display: flex;
    flex-direction: column;
    background-image: url("/backgr.png");
    background-repeat: repeat-y;
    background-position: center;
    background-size: 100%;
}
header {
    height: 1vw; 
    width: 100%;
    background-color: #000000;
}
.mainpart, .mainpart2{
    display: flex;
    flex-direction:row;
}
footer{
    background-color: #000000;
    width: 100%;
    margin-top: 200vw;
}
.containerHeader{
    display: flex;
    position: relative;
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
</style>