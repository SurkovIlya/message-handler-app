CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    sender VARCHAR(25),
    recipient VARCHAR(25),
    body_message VARCHAR(255),
    processed INT,
);