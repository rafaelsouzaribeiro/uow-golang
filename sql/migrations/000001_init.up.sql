CREATE TABLE authors (
    id   varchar(36)  NOT NULL PRIMARY KEY,
    name text    NOT NULL,
    bio  text
);

CREATE TABLE books (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  authors_id   varchar(36)  NOT NULL,
  name text    NOT NULL,
  description  text,
  price  decimal(10,2)  NOT NULL,
  FOREIGN KEY (authors_id) REFERENCES authors(id)
);



