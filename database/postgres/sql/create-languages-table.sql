CREATE TABLE languages (
  id UUID PRIMARY KEY NOT NULL,
  code VARCHAR(5) UNIQUE NOT NULL,
  title VARCHAR(15) NOT NULL
);