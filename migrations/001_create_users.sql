CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

INSERT INTO users (name, email)
VALUES
('Dustin', 'dustin@example.com'),
('Test User', 'test@example.com')
ON CONFLICT DO NOTHING;
