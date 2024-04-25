CREATE TABLE profiles (
  id UUID NOT NULL,
  name VARCHAR(50) NOT NULL,
  account_id UUID NOT NULL,
  language_id UUID,
  CONSTRAINT pk_profile PRIMARY KEY(id),
  CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES accounts(id) ON DELETE CASCADE,
  CONSTRAINT fk_language FOREIGN KEY(language_id) REFERENCES languages(id) ON DELETE SET NULL
);