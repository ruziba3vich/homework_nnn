CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(50),
    username VARCHAR(64),
    email VARCHAR(64),
    password VARCHAR(64),
    country VARCHAR(50),
    city VARCHAR(50),
    role VARCHAR(8)
);
