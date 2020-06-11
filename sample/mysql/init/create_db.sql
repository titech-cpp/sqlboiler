DROP DATABASE IF EXISTS sqlboiler_sample;
CREATE DATABASE sqlboiler_sample;
USE sqlboiler_sample;

CREATE TABLE IF NOT EXISTS users (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(32) NOT NULL,
    password CHAR(128) NOT NULL,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO users (name,password) VALUES ("test", "test");