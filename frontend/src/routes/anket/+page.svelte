<html lang="js">

<body>

    <div class=container2>
    <div><h1><p class=naming1>hack.itam</p></h1></div>
    <div><h2><p class=hrefComands>команды</p></h2></div>
    <div><h2><p class=hrefNews>новости</p></h2></div>
    <div><h2><p class=hrefProfile>профиль</p></h2></div>
    <div><h2><p class=hrefHacks>хакатоны</p></h2></div>
    <div><h2><p class=hrefAnkets>анкеты</p></h2></div>
    <div class=notific><button class=notific1></button>
    <ul>
        <li>В Вашей команде изменился статус капитана</li>
        <li>Ваша заявка на присоединение к команде отклонена</li>
        <li> Вы стали капитаном команды</li>
        <li>Павел Волков отправил заявку на присоединение к Вашей команде</li>
        <li>Павел Волков приглашает Вас в свою команду</li>
        <li>Павел Волков присоединлся к Вашей команде</li>
    </ul></div>
</div>


    <div class=backgr1>
    <img src="/backgr1.png"alt="" style="max-width: 100%;">
    </div>
    <div class="container1">
    <div class="name">
    <form action="" method="">
    <label class="name_label" for="surname">Фамилия</label>
    <input class="name_input" placeholder="Иванов" type="text">
    </form>
    <form action="" method="">
    <label class="name_label" for="name1">Имя</label>
    <input class="name_input" placeholder="Иван" type="text">
    </form>
    <form action="" method="">
    <label class="name_label" for="patronymic">Отчество</label>
    <input class="name_input" placeholder="Иванович" type="text">
    </form>
    <form action="" method="">
    <label class="name_label" for="academicGroup">Академическая группа</label>
    <input class="name_input" placeholder="БИВТ-10-8" type="text">
    </form>
    <form action="" method="">
    <label class="name_label" for="Telegram">Ник в Telegram (обязательно через @)</label> 
    <input class="name_input" placeholder="@abcdef" type="" name="" pattern="^@" required>
    </form>
    <label class="role_label" for="Role">Желаемая роль</label>
    <select class="role_input" name="" id="">
        <option selected disabled value="">Не выбрано</option>
        <option value="">Бэкенд</option>
        <option value="">ML/DS/AI</option>
        <option value="">Фронтенд</option>
        <option value="">Дизайн</option>
        <option value="">Product/Project</option>
        <option value="">Аналитик</option>
    </select>
    <label class="role_label" for="Role">Стек технологий</label>
    <select class="role_input" name="" id="">
        <option selected disabled value="">Не выбрано</option>
        <option value="">Бэкенд</option>
        <option value="">ML/DS/AI</option>
        <option value="">Фронтенд</option>
        <option value="">Дизайн</option>
        <option value="">Product/Project</option>
        <option value="">Аналитик</option>
    </select>
    <form action="" method="">
    <label class="name_label" for="surname">Другие технологии</label>
    <textarea class="tech_input"></textarea>
    </form>
    <!-- это должен быть мультиселект!!! -->
    <form action="" method="">
    <label class="name_label" for="aboutUser">О себе</label>
    <textarea class="about_text"></textarea>
    </form> 
    <label class="role_label" for="Role">Ищу команду:</label>
    <select class="role_input" name="" id="">
        <option selected disabled value="">Не выбрано</option>
        <option value="">Да</option>
        <option value="">Нет</option>
    </select>
    </div>
    <div class="button11">
    <button class="button12">СОХРАНИТЬ</button>
    </div>
</div>
</body>
</html>

<style>
    .notific ul{
        color: white;
        position: absolute;
        display: none;
        width: 20vw;
        margin-right: 10vw;
    }
    .notific ul li{
        display:block;
    }
.notific1:hover{
    background-color:grey;

}
header .notific:hover ul{
    display: block;
}
.container1{
    position: relative;
    width: 50vw;
    height: 75vw;
    background: rgba(0, 0, 0, 0.3);
    z-index: 2;
    margin-left: 25vw;
    margin-top: 10vw;
    border-radius: 3vw;
    border: 0.3vw solid black;
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
    margin-left: 5vw;
    margin-top: 2vw;
    width: 40vw;
}

.name_label, .role_label{
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
  
.name_input, .role_input, .tech_input, .about_text{
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

.name_input{
    height: 2.75vw;
    
}

.role_input{
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
    z-index: 2;
    margin-top: 68vw;
    margin-left: 5vw;
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