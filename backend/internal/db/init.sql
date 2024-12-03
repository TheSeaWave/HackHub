-- Создаем таблицу пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,          -- Уникальный идентификатор
    name VARCHAR(100) NOT NULL,     -- Имя пользователя
    email VARCHAR(100) UNIQUE NOT NULL, -- Email пользователя
    created_at TIMESTAMP DEFAULT NOW(), -- Время создания записи
    updated_at TIMESTAMP DEFAULT NOW()  -- Время последнего обновления
);

-- (Опционально) Заполняем тестовыми данными
INSERT INTO users (name, email) VALUES
    ('Alice', 'alice@example.com'),
    ('Bob', 'bob@example.com');
