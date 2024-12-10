CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    profile_picture VARCHAR(255)
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    captain_id INT NOT NULL,
    avatar VARCHAR(255),
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE surveys (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    "group" VARCHAR(255) NOT NULL,
    telegram VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    stack TEXT[] NOT NULL,
    about TEXT NOT NULL,
    achievements TEXT[] NOT NULL,
    status VARCHAR(255) NOT NULL,
    portfolio_link VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    teams TEXT[] NOT NULL,
    last_updated TIMESTAMP NOT NULL,
    rating INT,
    experience INT,
    "like" INT
);
