CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    key_id VARCHAR(50),
    body_message VARCHAR(2000),
    processed INT DEFAULT 0,
    is_read INT DEFAULT 0
);