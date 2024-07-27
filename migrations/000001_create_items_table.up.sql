CREATE TABLE IF NOT EXISTS messages (
    id BIGINT PRIMARY KEY,
    body_message VARCHAR(2000),
    processed BOOLEAN DEFAULT false
);