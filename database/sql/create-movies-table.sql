CREATE TABLE movies(
    id UUID NOT NULL,
    genre_id UUID,
    CONSTRAINT pk_movie PRIMARY KEY(id),
    CONSTRAINT fk_genre FOREIGN KEY(genre_id) REFERENCES genres(id) ON DELETE SET NULL
);