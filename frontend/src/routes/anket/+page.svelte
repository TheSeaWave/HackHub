<script lang="ts">
    let firstName = '';
    let lastName = '';
    let patronymic = '';
    let group = '';
    let telegram = '';
    let role = '';
    let techStack: string[] = [];  
    let otherTech = '';
    let about = '';
    let achievements = '';
    let isFound = false;  
  
    const roles = ['Бэкенд', 'ML/DS/AI', 'Фронтенд', 'Дизайн', 'Product/Project', 'Аналитик'];
    const technologies = ['React', 'Vue', 'Svelte', 'Node.js', 'Python', 'Java'];
  
    async function submitForm() {
      const formData = {
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
  
      const response = await fetch('https://your-service.com/api/submit-form', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });
  
      const result = await response.json();
      if (response.ok) {
        alert('Анкета создана');
      } else {
        alert('Ошибка при создании анкеты:' + result.message);
      }
    }
  </script>
  
<div class=container2>
    <div  class="naming1"><a href="/main">hack.itam</a></div>
    <div class=hrefComands><a href="/comands">команды</a></div>
    <div class=hrefNews><a href="/news">новости</a></div>
    <div class=hrefProfile><a href="/profile">профиль</a></div>
    <div class=hrefHacks><a href="/hacks">хакатоны</a></div>
    <div class=hrefAnkets><a href="/users">анкеты</a></div>
    <div class=notific><button class=notific1></button></div>
</div>


<div class=backgr1>
<img src="/backgr1.png"alt="" style="max-width: 100%;">
</div>

<div class="container1">
<div class="name">
<form on:submit|preventDefault={submitForm}>
            
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
    
<div class="button11"><button class="button12" type="submit">Отправить</button></div>
</form>
</div></div>

<div class="footer">
    <div class=containerFooter>
    <div class=footer1><p>© 2024</p></div>
    <div class=footer2><p>hack.itam</p></div>
    </div>
</div>


<style>
.notific1:hover{
    background-color:grey;
} 

.container1{
    position: relative;
    width: 50vw;
    height: 85vw;
    background: rgba(0, 0, 0, 0.3);
    z-index: 2;
    border-radius: 3vw;
    border: 0.3vw solid black;
    margin: 10vw 0 0 25vw;
}

:global(body) {
    display: flex;
    flex-direction: column;
    background-color: #000000;
}

label{
    color: white;
}
.footer{
    background-color: #000000;
    width: 100%;
    margin-top: 30vw;
    height: 1%; 
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

p{
    color: white;
}

.name{
    text-align: center;
    position: absolute;
    z-index: 2;
    width: 40vw;
    margin: 2vw, 0, 0, 2vw;
}

.tech_label, .role_label, .about_label, .found_label, .stack_label, .telegram_label, .lastName_label, .firstName_label, .patronomic_label, .group_label{
    display: block;
    text-align: left;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-weight: 100;
    font-style: normal;
    font-size: 1.5vw;
    margin: 0.5vw 0;

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
    position:absolute;
    margin-top: 3.5vw;
    margin-left: 2vw;
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

</style>