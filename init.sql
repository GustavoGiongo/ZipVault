CREATE TABLE IF NOT EXISTS files (
                                     id INT AUTO_INCREMENT PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    folder VARCHAR(255) NOT NULL
    );

DELETE FROM files;

INSERT INTO files (name, date, folder)
VALUES
    ('dummyfile.png', '2025-05-21', '/Pictures/Screenshots'),
    ('dum.png',           '2025-05-21', '/Pictures/Screenshots'),
    ('dummy.png',         '2025-05-21', '/Pictures/Screenshots');
