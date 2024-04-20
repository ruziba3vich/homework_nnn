CREATE TABLE IF NOT EXISTS Required_skills (
    id SERIAL PRIMARY KEY,
    vacancy_id INTEGER REFERENCES Vacancies(id),
    skill_id INTEGER REFERENCES Skills(id)
);
