CREATE TABLE phone_numbers (
  id SERIAL PRIMARY KEY,
  value TEXT
);

INSERT INTO phone_numbers (value) VALUES ('1234567890');
INSERT INTO phone_numbers (value) VALUES ('123 456 7891');
INSERT INTO phone_numbers (value) VALUES ('(123) 456 7892');
INSERT INTO phone_numbers (value) VALUES ('(123) 456-7893');
INSERT INTO phone_numbers (value) VALUES ('123-456-7894');
INSERT INTO phone_numbers (value) VALUES ('123-456-7890');
INSERT INTO phone_numbers (value) VALUES ('1234567892');
INSERT INTO phone_numbers (value) VALUES ('(123)456-7892');
