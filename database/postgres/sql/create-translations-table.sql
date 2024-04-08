CREATE TABLE translations (
    entity_id UUID NOT NULL,
    language_id UUID NOT NULL,
    type VARCHAR(15) NOT NULL,
    text VARCHAR(300) NOT NULL,
    CONSTRAINT pk_translation PRIMARY KEY(entity_id, language_id, type),
    CONSTRAINT fk_language FOREIGN KEY(language_id) REFERENCES languages(id) ON DELETE CASCADE,
    CONSTRAINT check_type CHECK(type IN('LABEL', 'TITLE', 'DESCRIPTION'))
);