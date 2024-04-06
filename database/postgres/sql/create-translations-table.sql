CREATE TABLE translations (
    entity_id UUID NOT NULL,
    language_code VARCHAR(5) NOT NULL,
    text VARCHAR(300) NOT NULL,
    PRIMARY KEY(entity_id, language_code)
);