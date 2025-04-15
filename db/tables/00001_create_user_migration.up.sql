CREATE TABLE IF NOT EXISTS user (
    id        INT AUTO_INCREMENT PRIMARY KEY,
    profile_pic VARCHAR(200),
    username  VARCHAR(50) NOT NULL,
    email      VARCHAR(50) NOT NULL UNIQUE,
    `password`   VARCHAR(200) NOT NULL
);
