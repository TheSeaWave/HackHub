<script lang="ts">
    import { onMount } from 'svelte';
    
    let firstName: string = 'Не задано';
    let lastName: string = 'Не задано';
    let patronymic: string = 'Не задано';
    let group: string = 'Не задано';
    let telegram: string = 'Не задано';
    let role: string = 'Не задано';
    let techStack: string[] = [];
    let otherTech: string = 'Не задано';
    let about: string = 'Не задано';
    let achievements: string = 'Не задано';
    let isFound: boolean = false;
    let isEditModalOpen: boolean = false;
    const roles = ['Бэкенд', 'ML/DS/AI', 'Фронтенд', 'Дизайн', 'Product/Project', 'Аналитик'];
    const technologies = ['React', 'Vue', 'Svelte', 'Node.js', 'Python', 'Java'];
  
    async function fetchUserProfile() {
        try {
            const response = await fetch('https://your-service.com/api/user-profile');
            if (!response.ok) {
                throw new Error('Error fetching user profile');
            }
  
            const data = await response.json();
            
            firstName = data.firstName || '';
            lastName = data.lastName || '';
            patronymic = data.patronymic || '';
            group = data.group || '';
            telegram = data.telegram || '';
            role = data.role || '';
            techStack = data.techStack || [];
            otherTech = data.otherTech || '';
            about = data.about || '';
            achievements = data.achievements || '';
            isFound = data.isFound || false;
        } catch (error: unknown) {
            if (error instanceof Error) {
                console.error('Failed to fetch user profile:', error.message);
            } else {
                console.error('An unknown error occurred');
            }
        }
    }
  
    onMount(() => {
        fetchUserProfile();
    });
  
    async function submitProfile() {
        const profileData = {
            firstName,
            lastName,
            patronymic,
            group,
            telegram,
            role,
            techStack,
            otherTech,
            about,
            achievements,
            isFound,
        };
  
        try {
            const response = await fetch('/api/update-profile', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(profileData),
            });
  
            if (!response.ok) {
                throw new Error('Error updating profile');
            }
  
            alert('Profile updated successfully!');
        } catch (error: unknown) {
            if (error instanceof Error) {
                alert('Error updating profile: ' + error.message);
            } else {
                alert('Error updating profile: Unknown error');
            }
        }
    }
    function closeEditModal() {
        isEditModalOpen = false;
    }
</script>

<header>

<div class=container2>
    <div><h1><p class=naming1>hack.itam</p></h1></div>
    <div><h2><p class=hrefComands>команды</p></h2></div>
    <div><h2><p class=hrefNews>новости</p></h2></div>
    <div><h2><p class=hrefProfile>профиль</p></h2></div>
    <div><h2><p class=hrefHacks>хакатоны</p></h2></div>
    <div><h2><p class=hrefAnkets>анкеты</p></h2></div>
    <div class=notific><button class=notific1></button>
        </div>
        </div>
    </header>
    <body>
<div class="profile">
    <div class=backgr1>
        <img src="/backgr1.png"alt="" style="max-width: 100%;">
        </div>
       
        <div class="title">ВАШ ПРОФИЛЬ</div>
        <div class="container1">
            
        <div class="name">
    <p><span class="all_label">Имя:</span><span class="all_info">{firstName}</span></p>
    <p><span class="all_label">Фамилия:</span> <span class="all_info">{lastName}</span></p>
    <p><span class="all_label">Отчество:</span> <span class="all_info">{patronymic}</span></p>
    <p><span class="all_label">Группа:</span> <span class="all_info">{group}</span></p>
    <p><span class="all_label">Роль:</span><span class="all_info">{role}</span></p>
    <p><span class="all_label">Технологии:</span> <span class="all_info">{techStack.join(', ')}</span></p>
    <p><span class="all_label">Другие технологии:</span> <span class="all_info">{otherTech}</span></p>
    <p><span class="all_label">О себе:</span> <span class="all_info">{about}</span></p>
    <p><span class="all_label">Достижения:</span> <span class="all_info">{achievements}</span></p>
    <p><span class="all_label">Ищу команду:</span> <span class="all_info">{isFound ? 'Да' : 'Нет'}</span></p>

   
        <div class="button11">
            <button class="button12" on:click={() => isEditModalOpen = true}>РЕДАКТИРОВАТЬ АНКЕТУ</button>
        </div>
        {#if isEditModalOpen}
        <div class="modal">
            <div class="modal-content">
                <button class="close" on:click={closeEditModal}>X</button>
                <form on:submit|preventDefault={submitProfile}>  
            
            <div>
                <label class="lastName_label" for="lastName">Фамилия:</label>
                <input class="lastName_input" placeholder="Иванов" type="text" id="lastName" bind:value={lastName} />
            </div>

            <div>
                <label class="firstName_label" for="firstName">Имя:</label>
                <input class="firstName_input" placeholder="Иван" type="text" id="firstName" bind:value={firstName} />
            </div>
            
            <div>
                <label class="patronomic_label" for="patronymic">Отчество:</label>
                <input class="patronomic_input" placeholder="Иванович" type="text" id="patronymic" bind:value={patronymic} />
            </div>

            <div>
                <label class="group_label" for="group">Академическая группа:</label>
                <input class="group_input" placeholder="БИВТ-10-8" type="text" id="group" bind:value={group} />
            </div>

            <div>
                <label class="telegram_label" for="telegram">Ник в Telegram (обязательно через @):</label>
                <input class="telegram_input" placeholder="@abcdef" type="text" pattern="^@" required id="telegram" bind:value={telegram} />
            </div>

            <div>
                <label class="role_label" for="role">Желаемая роль:</label>
                <select class="role_input" id="role" bind:value={role}>
                    {#each roles as roleOption}
                        <option value={roleOption}>{roleOption}</option>
                    {/each}
                </select>
            </div>

            <div>
                <label class="stack_label" for="techStack">Стек технологий:</label>
                <select class="stack_input" id="techStack" bind:value={techStack} multiple>
                    {#each technologies as tech}
                        <option value={tech}>{tech}</option>
                    {/each}
                </select>
            </div>
            
            <div>
                <label class="tech_label" for="otherTech">Другие технологии:</label>
                <textarea class="tech_input" id="otherTech" bind:value={otherTech}></textarea>
            </div>
            
            <div>
                <label class="about_label" for="about">О себе:</label>
                <textarea class="about_text" id="about" bind:value={about}></textarea>
            </div>
            <div>
                <label class="about_label" for="achievements">Достижения:</label>
                <textarea class="tech_input" id="achievements" bind:value={achievements}></textarea>
            </div>
            <div>
                <label class="found_label" for="isFound">Ищу команду:</label>
                <select class="found_input" id="isFound" bind:value={isFound}>
                    <option value={false}>Нет</option>
                    <option value={true}>Да</option>
                </select>
            </div>

    
            <div class="button11">
                <button class="button12" type="submit">Сохранить изменения</button>
            </div>
        </form>
    </div>
</div>
{/if}
</div>
</body>
<footer>
    <div class=containerFooter>
        <div class=footer1><p>© 2024</p></div>
        <div class=footer2><p>hack.itam</p></div>
        </div>
</footer>

<style>
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
  .close {
      position: absolute;
      top: 1vw;
      right: 1.3vw;
      background: rgb(61, 57, 57);
      color: white;
      border: none;
      padding: 0.3vw;
      cursor: pointer;
      width: 3vw;
      height: 3vw;
      text-align: center;
      z-index: 2;
  }
  .title{
    position: absolute;
    z-index: 2;
    font-size: 4vw;
    font-family: "Russo One", sans-serif;
    font-weight: 50;
    font-style: normal;
    color: #ffffff;
    margin-left: 27vw;
    margin-top: 10vw;
  }
header{
    display: block;
}
.container1{
    position: relative;
    width: 50vw;
    background: rgba(0, 0, 0, 0.3);
    z-index: 2;
    margin-left: 25vw;
    margin-top: 17vw;
    border-radius: 3vw;
    border: 0.3vw solid black;
    min-height: 70vw; 
      height: auto; 

}
.all_label{
    font-weight: 300;
}

.all_info, .all_label{
    display: block;
    text-align: left;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-weight: 100;
    font-style: normal;
    margin-bottom: 1vw;
    margin-top: 1vw;
    margin-left: 1vw;
    font-size: 2vw;
}
body {
    display: flex;
    flex-direction: column;
    background-color: #000000;
}

label{
    color: white;
}

.backgr1 {
    width: 100%;
    margin-left: 0%;
    margin-right: 0%;
    position: absolute;
    z-index: 1;
    margin-top: -2vw;

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
}

.notific{
    margin-left: auto;
    position: relative;
}

.notific1{
    background-image: url('/notific.png');
    margin-top: 3.5vw;
    margin-right: 5vw;
}

footer {
    height: 1%; 
    width: 100%;
    background-color: #000000;
    margin-top: 10vw;
}

.notific1{
   width: 3vw;
   height: 3vw;
   background-size: cover;
   background-position: center;
   background-repeat: no-repeat;
   border: 0;
   background-color: #000000;
   z-index: 2;
}

p{
    color: white;
}

.naming1{
    color:white;
}

.name{
    text-align: center;
    position: absolute;
    z-index: 2;
    margin-left: 2vw;
    margin-top: 2vw;
    width: 40vw;
}

.tech_label, .role_label, .about_label, .found_label, .stack_label, .telegram_label, .lastName_label, .firstName_label, .patronomic_label, .group_label{
    display: block;
    text-align: left;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-weight: 100;
    font-style: normal;
    margin-bottom: 0.5vw;
    margin-top: 0.5vw;
    font-size: 1.5vw;

}

.role_input, .found_input, .stack_input, .tech_input, .about_text, .telegram_input, .group_input, .lastName_input, .firstName_input, .patronomic_input{
    display: block;
    width: 100%;
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
}

.lastName_input, .telegram_input, .group_input, .firstName_input, .patronomic_input{
    height: 2.75vw;
    
}

.role_input, .found_input, .stack_input{
    height: 3.5vw;
}

.tech_input{
    height: 3vw;
    resize: none;
}

.about_text{
    height: 7vw;
    resize: none;
}

.button11{    
    margin-left: 2vw;
    margin-top: 3vw;
    margin-bottom: 2vw;
}

.button12{
    border: 0;
    color: #ffffff;
    background: rgba(0, 0, 0, 0.3);
    padding: 1.5vw 4vw;
    border-radius: 1vw;
    border: 0.3vw solid white;
    font-size: 1.3vw;
    font-family: "Rubik Mono One", monospace;
    font-weight: 400;
    font-style: normal;
}
footer{
    background-color: #000000;
    width: 100%;
    margin-top: 30vw;
    
} 
.containerFooter{
    display: flex;
    position: relative;
    margin-bottom: 0;
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

</style>