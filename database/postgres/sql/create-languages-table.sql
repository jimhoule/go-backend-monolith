CREATE TABLE languages (
  id UUID NOT NULL,
  code VARCHAR(5) NOT NULL,
  CONSTRAINT pk_language PRIMARY KEY(id),
  CONSTRAINT unique_code UNIQUE(code)
);