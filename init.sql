CREATE TABLE IF NOT EXISTS files (
                                       id INT AUTO_INCREMENT PRIMARY KEY,
                                       name VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    location VARCHAR(255) NOT NULL
    );

INSERT INTO files (name, date, location)
VALUES ('meu-arquivo.zip', '2025-05-21', '/arquivos/meu-arquivo.zip');
