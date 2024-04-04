CREATE TABLE translations (
    entityId UUID NOT NULL,
    languageCode VARCHAR(5) NOT NULL,
    text VARCHAR(300) NOT NULL,
    PRIMARY KEY(entityId, languageCode)
);