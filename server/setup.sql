CREATE DATABASE Sandbox;

USE Sandbox;

CREATE TABLE Users ( 
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255),
    password VARCHAR(255),
    created DATE
);

CREATE TABLE Articles ( 
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255),
    article LONGTEXT,
    created DATE
);

CREATE TABLE Messages ( 
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    message LONGTEXT,
    created DATE
);