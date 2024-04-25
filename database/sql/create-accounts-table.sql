CREATE TABLE accounts (
  id UUID NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password VARCHAR(200) NOT NULL,
  is_memberShip_cancelled BOOLEAN NOT NULL,
  plan_id UUID NOT NULL,
  CONSTRAINT pk_account PRIMARY KEY(id),
  CONSTRAINT fk_plan FOREIGN KEY(plan_id) REFERENCES plans(id),
  CONSTRAINT unique_email UNIQUE(email)
);