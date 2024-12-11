<script>
    let teamName = '';
  let teamDescription = '';
  let teamMembersInput = '';  
  
  async function submitTeamForm() {
      const teamMembers = teamMembersInput.split(',').map(id => id.trim());
      const teamData = {
          teamName,
          teamDescription,
          teamMembers,  
      };
  
      try {
          const response = await fetch('https://your-api.com/api/teams', {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
              },
              body: JSON.stringify(teamData),
          });
  
          const result = await response.json();
          if (response.ok) {
              alert('Команда создана');
          } else {
              alert('Ошибка при создании команды: ' + result.message);
          }
      } catch (error) {
          console.error(error);
          alert('Ошибка соединения');
      }
  }
  </script>
  
<div class=container2>
    <div  class="naming1"><a href="/main">hack.itam</a></div>
    <div class="hrefComands"><a href="/comands">команды</a></div>
    <div class="hrefNews"><a href="/news">новости</a></div>
    <div class="hrefProfile"><a href="/profile">профиль</a></div>
    <div class="hrefHacks"><a href="/hacks">хакатоны</a></div>
    <div class="hrefAnkets"><a href="/users">анкеты</a></div>
    <div class="notific"><button class="notific1"></button></div>
</div>

<form on:submit|preventDefault={submitTeamForm}>
<div>
    <label for="teamName">Название команды:</label>
    <input id="teamName" bind:value={teamName} required />
</div>

<div>
    <label for="teamDescription">Описание команды:</label>
    <textarea id="teamDescription" bind:value={teamDescription}></textarea>
</div>

<div>
    <label for="teamMembers">Участники команды (по ID):</label>
    <input 
        type="text" 
        id="teamMembers" 
        bind:value={teamMembersInput}
        placeholder="Введите ID участников через запятую" />
</div>

<div class="button11">
    <button type="submit">Создать команду</button>
</div>
</form>

<div class="footer">
    <div class=containerFooter>
    <div class=footer1><p>© 2024</p></div>
    <div class=footer2><p>hack.itam</p></div>
    </div>
</div>

<style>
:global(body) {
    display: flex;
    flex-direction: column;
    background-color: #000000;
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
</style>

