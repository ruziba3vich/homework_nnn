CREATE TABLE IF NOT EXISTS Email_box (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    employer_id INTEGER REFERENCES Users(id),
    content TEXT,
    sent_on TIME
);
