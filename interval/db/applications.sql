CREATE TABLE IF NOT EXISTS Applications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Employers(id),
    vacancy_id INTEGER REFERENCES Vacancies(id)
);
