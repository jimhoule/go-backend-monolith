CREATE TABLE profiles (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  account_id UUID NOT NULL,
  language_id UUID NOT NULL,
  CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES accounts(id),
  CONSTRAINT fk_language FOREIGN KEY(language_id) REFERENCES languages(id)
);