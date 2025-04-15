
CREATE TABLE IF NOT EXISTS Startup (
    id       INT AUTO_INCREMENT PRIMARY KEY,
    title    VARCHAR(100) NOT NULL,
    category VARCHAR(20) NOT NULL,
    pitch    VARCHAR(100) NOT NULL,
    `image`  VARCHAR(200) NOT NULL,
    slug     VARCHAR(20) NOT NULL,
    `desc`   VARCHAR(200) NOT NULL,
    userId   INT NOT NULL,
    FOREIGN KEY (userId) REFERENCES User(id)
)