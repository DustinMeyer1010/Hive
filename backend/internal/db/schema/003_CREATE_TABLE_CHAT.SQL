CREATE TABLE chat (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    creation_date DATE NOT NULL,
    FOREIGN KEY (hive_id) REFERNCES hive(hive_id)
)