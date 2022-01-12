CREATE TABLE users (
  id        bigint(20)     PRIMARY KEY,
  email     varchar(200)   DEFAULT NULL UNIQUE,
  password longblob
);