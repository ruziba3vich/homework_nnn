CREATE TABLE IF NOT EXISTS Employers (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(50),
    username VARCHAR(64),
    email VARCHAR(64),
    password VARCHAR(64)
);
