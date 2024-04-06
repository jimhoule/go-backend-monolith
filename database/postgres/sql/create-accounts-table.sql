CREATE TABLE accounts (
  id UUID PRIMARY KEY NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(200) NOT NULL,
  is_memberShip_cancelled BOOLEAN NOT NULL,
  plan_id UUID NOT NULL,
  CONSTRAINT fk_plan FOREIGN KEY(plan_id) REFERENCES plans(id)
);