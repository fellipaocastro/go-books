-- ----------------------------
--  Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS books;
CREATE TABLE books (
    id serial PRIMARY KEY,
    title varchar (255) NOT NULL,
    author varchar (40) NOT NULL,
    description text
);

-- ----------------------------
--  Records of books
-- ----------------------------
BEGIN;
INSERT INTO books (title, author, description) VALUES ('Livro de teste', 'Fellipe Castro', 'A young hipster bear seeks his fortune in the wild city of Irvine.');
COMMIT;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar (255) NOT NULL,
    email varchar (255) NOT NULL,
    password varchar (255) NOT NULL
);

BEGIN;
INSERT INTO users(name, email, password) VALUES ('Fellipe', 'contact@fellipecastro.com', 'qwer1234');
COMMIT;
