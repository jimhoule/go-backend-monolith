CREATE TABLE plans (
  id UUID NOT NULL,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(100) NOT NULL,
  price FLOAT NOT NULL,
  CONSTRAINT pk_plan PRIMARY KEY(id)
);