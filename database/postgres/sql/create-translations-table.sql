CREATE TABLE translations (
    entity_id UUID NOT NULL,
    language_id UUID NOT NULL,
    text VARCHAR(300) NOT NULL,
    CONSTRAINT pk_translation PRIMARY KEY(entity_id, language_id),
    CONSTRAINT fk_language FOREIGN KEY(language_id) REFERENCES languages(id) ON DELETE CASCADE
);