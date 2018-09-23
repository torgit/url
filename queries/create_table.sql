CREATE TABLE url (
    id int NOT NULL AUTO_INCREMENT,
    originalUrl varchar(255),
    shortUrl varchar(255),
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
    PRIMARY KEY (ID)
);