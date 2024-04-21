CREATE TABLE IF NOT EXISTS Vacancies (
    id SERIAL PRIMARY KEY,
    company_name VARCHAR(100),
    employer_id INTEGER REFERENCES Users(id),
    title VARCHAR(100),
    content TEXT,
    country VARCHAR(100),
    city VARCHAR(100),
    is_active BOOLEAN,
    position VARCHAR(100)
);
