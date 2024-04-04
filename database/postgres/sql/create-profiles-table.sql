CREATE TABLE profiles (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  accountId UUID NOT NULL,
  languageId UUID NOT NULL,
  CONSTRAINT fk_account FOREIGN KEY(accountId) REFERENCES accounts(id),
  CONSTRAINT fk_language FOREIGN KEY(languageId) REFERENCES languages(id)
);