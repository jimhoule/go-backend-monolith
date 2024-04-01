CREATE TABLE profiles (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  accountId UUID NOT NULL,
  CONSTRAINT fk_account FOREIGN KEY(accountId) REFERENCES accounts(id)
);