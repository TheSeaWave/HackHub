<script lang="ts">
    import { v4 as uuidv4 } from 'uuid';  // Для генерации уникальных ID

    let email = '';
    let password = '';
    let message = '';
    const userId = uuidv4();  // Генерация уникального ID для нового пользователя

    // Регистрация пользователя
    async function register() {
        const response = await fetch('/api/auth/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });
        const data = await response.json();
        message = data.message;
    }

    async function login() {
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });
        const data = await response.json();
        if (data.token) {
            localStorage.setItem('token', data.token);  
            message = 'Вы успешно вошли!';
            window.location.href = '/main';
        } else {
            message = data.message;
        }
    }
</script>


<body>
<div class="signin-container">
    <p class="signin">Вход</p>
    <div class="signin-label">E-mail:</div>
    <input class="signin-email" 
        type="email" 
        bind:value={email} 
        placeholder="Email" 
        required 
        aria-label="Email" 
    />
    <div class="signin-password-label">Пароль:</div>
    <input class="signin-password"
        type="password" 
        bind:value={password} 
        placeholder="Пароль" 
        required 
        aria-label="Пароль" 
    />
    <button class="signin-button" on:click={login} aria-label="Войти">
        Войти
    </button>
</div>

<div class="login-container">
    <p class="login">Регистрация</p>
    <div class="login-label">E-mail:</div>
    <input class="login-email"
        type="email" 
        bind:value={email} 
        placeholder="Email" 
        required 
        aria-label="Email" 
    />
    <div class="login-password-label">Пароль:</div>
    <input class="login-password"
        type="password" 
        bind:value={password} 
        placeholder="Пароль" 
        required 
        aria-label="Пароль" 
    />
    <button class="login-button" on:click={register} aria-label="Зарегистрироваться">
        Зарегистрироваться
    </button>
</div>

<p>{message}</p>
</body>
<footer>
    <div class=containerFooter>
    <div class=footer1><p>© 2024</p></div>
    <div class=footer2><p>hack.itam</p></div>
    </div>
</footer>
<style>
.login-container, .signin-container{
    width: 55vw;
    color: white;
    position: relative;
    height: 35vw;
    background: rgba(0, 0, 0, 0.3);
    margin-left: 20vw;
    margin-top: 5vw;
    border-radius: 2vw;
    border: 0.3vw solid black;
}
body{
    display: flex;
    flex-direction: column;
    background-image: url("/backgr.png");
    background-position: center;
    background-size: 100%;
}
.signin, .login{
    margin-top: 4vw;
    margin-left: 2vw;
    font-size: 4vw;
    font-family: "Russo One", sans-serif;
    font-weight: 400;
    font-style: normal;
    
    color: rgb(255, 255, 255);
}
.signin-label, .signin-password-label, .login-label, .login-password-label {
    margin-top: 2vw;
    margin-left: 2vw;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    font-weight: 200;
    color: rgb(255, 255, 255);
    font-size: 2.5vw;
}
.signin-email, .signin-password, .login-email, .login-password{
    display: block;
    width: 90%;
    height: 3vw;
    margin-left: 1vw;
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
    margin-top: 1vw;
}
.signin-button, .login-button{
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
    margin-top: 2vw;
    margin-left: 2vw;
}
footer{
    background: rgba(0, 0, 0, 0.3);
    width: 100%;
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
    margin-top: 3vw;
    color:white;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-style: normal;
    margin-bottom: 4vw;
}
</style>