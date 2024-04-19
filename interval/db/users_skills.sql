CREATE TABLE IF NOT EXISTS Users_skills (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    skill_id INTEGER REFERENCES Skills(id)
);
