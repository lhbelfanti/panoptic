CREATE TABLE IF NOT EXISTS example_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    data TEXT
);

INSERT INTO example_table(name, data) VALUES ('Example Name 1', 'Example Data 1');
INSERT INTO example_table(name, data) VALUES ('Example Name 2', 'Example Data 2');
