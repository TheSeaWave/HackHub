<script>
    import { onMount } from 'svelte';
    
    onMount(() => {
        const token = localStorage.getItem('token');
        if (!token) {
            window.location.href = '/regist';
        }
    });
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

<div class=back1>
<div class=backgr1>
<img src="/backgr1.png"alt="" style="max-width: 100%;">
</div>
<div class=backgr2>
<img src="/backgr3.png"alt="" style="max-width: 100%;">
</div>
    
<div class=text1><p>НАЙДИ СВОЮ КОМАНДУ</p></div>
<div class=mainButton><button class=mainButton1>ЗАПОЛНИТЬ АНКЕТУ</button></div>
<div class=text2><p>ЧТО ТАКОЕ ХАКАТОН?</p></div> 
<div class=text3><p>Хакатон — это соревнование между командами, которые за ограниченное время решают определенную задачу. Термин «хакатон» образован от слов hacker и marathon. Хакер — потому что соревнования обычно проводят в IT-сфере, марафон — потому что команды интенсивно работают несколько дней без перерыва. Разработчики получают задание, два-три дня готовят решение и презентуют программу, приложение или сервис.</p></div> 
<div class=text4><p>ЗАЧЕМ ОН НУЖЕН?</p></div> 
<div class=text5><p>В рамках хакатона команды генерируют идеи. У бизнеса есть возможность найти перспективные проекты, вложиться в них и в будущем монетизировать. Кроме того, организация хакатона — один из способов продвижения компании, в частности HR-бренда.</p></div>    
<div class=text6><p>КАКАЯ ПОЛЬЗА ДЛЯ ТЕБЯ?</p></div> 
<div class=text7><p>Хакатоны помогают участникам развивать навыки, работать в командах и улучшать коммуникацию. Они стимулируют креативное мышление и позволяют экспериментировать с новыми технологиями. Участники создают проекты, которые можно добавить в портфолио, а также налаживают контакты с единомышленниками и работодателями, что способствует личному и профессиональному росту. </p></div>
</div>

<div class="footer">
    <div class=containerFooter>
    <div class=footer1><p>© 2024</p></div>
    <div class=footer2><p>hack.itam</p></div>
    </div>
</div>

<style>
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

.back1{
    position: relative; 
}

.backgr1 {
    margin-top: -50vw;
    width: 97vw;
    margin-left: 0vw;
    margin-right: 0vw;
    position: absolute;
    z-index: 1;
}

.backgr2 {
    width: 97vw;
    margin-left: 0vw;
    margin-right: 50vw;
    position: absolute;
    z-index: 1;
    margin-top: 110vw;
}

.text1{
    position: absolute;
    z-index: 2;
    margin-top: 20vw;
    margin-left: 20vw;
    margin-right: 10vw;
    font-size: 5vw;
    font-family: "Russo One", sans-serif;
    font-weight: 400;
    font-style: normal;
}

.text2, .text4, .text6{
    position: absolute;
    z-index: 2;
    font-size: 4vw;
    font-family: "Russo One", sans-serif;
    font-weight: 50;
    font-style: normal;
}

.text2{
    margin-top: 50vw;
    margin-left: 8vw;
    margin-right: 15vw;
}

.text4{
    margin-top: 103vw;
    margin-left: 40vw;
    margin-right: 8vw;
}

.text6{
    margin-top: 145vw;
    margin-left: 8vw;
    margin-right: 15vw;
    font-size: 4vw;
}

.text3, .text5, .text7{
    position: absolute;
    z-index: 2;
    font-size: 2.7vw;
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-weight: 100;
    font-style: normal;
    font-variation-settings:
      "slnt" 0,
      "wdth" 100,
      "GRAD" -200,
      "XOPQ" 96,
      "XTRA" 468,
      "YOPQ" 79,
      "YTAS" 750,
      "YTDE" -203,
      "YTFI" 738,
      "YTLC" 514,
      "YTUC" 712;
}

.text3{
    margin-top: 57vw;
    margin-left: 8vw;
    margin-right: 40vw; 
}

.text5{
    margin-top: 110vw;
    margin-left: 40vw;
    margin-right: 8vw;
}

.text7{
    margin-top: 152vw;
    margin-left: 8vw;
    margin-right: 40vw;
}

.mainButton{
    position:absolute;
    z-index: 2;
    margin-top: 30vw;
    margin-left: 39vw;
}

:global(body) {
    display: flex;
    flex-direction: column;
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

.mainButton1{
    border: 0;
    color: #000000;
    background-color: #fff;
    padding: 1.5vw 2.5vw;
    border-radius: 4vw;
    font-size: 1.3vw;
    font-family: "Rubik Mono One", monospace;
    font-weight: 400;
    font-style: normal;
}

.notific1{
    background-image: url('/notific.png');
    margin-top: 3vw;
    margin-right: 4vw;
}

.footer{
    background-color: #000000;
    width: 100%;
    margin-top: 200vw;
}

.notific{
    margin-left: auto;
    z-index: 2;
}

.containerFooter{
    display: flex;
    position: relative;
    margin-bottom: 0;
}

.naming1{
    font-family: "Roboto Flex", sans-serif;
    font-optical-sizing: auto;
    font-weight: 200;
    font-style: normal;
    font-size: 2.5vw;
    margin-top: 3vw;
    margin-left: 5vw;
    color:white;
}
</style>