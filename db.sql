-- SQL for snippetbox 


-- Create a new UTF-8 snippetbox database 

CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Swith to using the snippetbox database 

USE snippetbox;

-- Create the snippets table

CREATE TABLE snippets(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL, 
    content TEXT NOT NULL,
    created DATETIME NOT NULL, 
    expires DATETIME NOT NULL
    
);

-- Add an index on the created column.

CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records

INSERT INTO snippets (title , content , created , expires ) VALUES (
    'An old silent pond',
    'An old silent pond ...  splash! Silence again',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY ) 
);


INSERT INTO snippets (title , content , created , expires ) VALUES (
    'Over the wintry forest',
    'Over the wintry forest ... la la la ',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY ) 
);


INSERT INTO snippets (title , content , created , expires ) VALUES (
    'First autunm morning',
    'First autunm morning ...  la lallala',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY ) 
);


-- Creating a new user. With SELECT and INSERT privileges only ib the database 

CREATE USER 'web'@'localhost';
GRAND SELCT , INSERT , UPDATE , DELETE ON snippetbox.* TO 'web'@'localhost';


-- import: Make sure to swap 'pass' with password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

