DROP DATABASE IF EXISTS toy;
CREATE DATABASE IF NOT EXISTS toy CHARACTER SET utf8 COLLATE utf8_general_ci;
USE toy

CREATE TABLE user (
    id int,
    name varchar(255)
) CHARACTER SET utf8 COLLATE utf8_general_ci;

INSERT INTO user VALUE (1,"sckshuhari");