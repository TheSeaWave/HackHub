-- Таблица пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,       -- Уникальный идентификатор пользователя
    name VARCHAR(255) NOT NULL,  -- Имя пользователя
    email VARCHAR(255) NOT NULL UNIQUE, -- Email (уникальный)
    password_hash TEXT NOT NULL, -- Хэш пароля
    created_at TIMESTAMP DEFAULT NOW() -- Дата создания записи
);

-- Таблица команд
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,                 -- Уникальный идентификатор команды
    name VARCHAR(255) NOT NULL,            -- Название команды
    description TEXT NOT NULL,             -- Описание команды
    captain_id INT NOT NULL,               -- ID капитана команды (FK к users)
    avatar TEXT NOT NULL,                  -- Ссылка на аватар команды
    created_at TIMESTAMP DEFAULT NOW(),    -- Дата создания команды
    FOREIGN KEY (captain_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица анкет
CREATE TABLE surveys (
    id SERIAL PRIMARY KEY,                   -- Уникальный идентификатор анкеты
    full_name VARCHAR(255) NOT NULL,         -- Полное имя
    "group" VARCHAR(255) NOT NULL,           -- Группа (или должность)
    telegram VARCHAR(255) NOT NULL,          -- Telegram
    role VARCHAR(255) NOT NULL,              -- Роль в проекте
    stack TEXT[] NOT NULL,                   -- Массив технологий (массив строк)
    about TEXT NOT NULL,                     -- Информация о пользователе
    achievements TEXT[] NOT NULL,            -- Массив достижений (массив строк)
    status VARCHAR(255) NOT NULL,            -- Статус (например, активный, завершённый)
    portfolio_link TEXT NOT NULL,            -- Ссылка на портфолио
    avatar TEXT NOT NULL,                    -- Ссылка на аватар
    teams TEXT[] NOT NULL,                   -- Список команд (массив строк)
    last_updated TIMESTAMP DEFAULT NOW(),    -- Дата последнего обновления анкеты
    rating INT DEFAULT 0,                    -- Рейтинг
    experience INT DEFAULT 0,                -- Опыт
    "like" INT DEFAULT 0                     -- Количество лайков
);