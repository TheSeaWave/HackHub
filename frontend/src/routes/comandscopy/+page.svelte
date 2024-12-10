<body>
 

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
          placeholder="Введите ID участников через запятую" 
      />
  </div>

  <div class="button11">
      <button type="submit">Создать команду</button>
  </div>
</form>

</body>

<style>
  body {
    display: flex;
    flex-direction: column;
    background-color: #000000;
}


</style>

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