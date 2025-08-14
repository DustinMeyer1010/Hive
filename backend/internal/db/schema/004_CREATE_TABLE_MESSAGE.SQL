CREATE TABLE message (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    FOREIGN KEY (chat_id) REFERNCES chat(chat_id)
)