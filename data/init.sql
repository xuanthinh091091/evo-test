CREATE TABLE users (
   id BIGSERIAL PRIMARY KEY NOT NULL,
    user_name TEXT NOT NULL,
    pass_word TEXT NOT NULL,
    role TEXT NOT NULL
);

CREATE UNIQUE INDEX user_name_uindex ON users (user_name);

INSERT INTO users(user_name, pass_word,role)
VALUES
('user','123456','USER'),
('admin','123456','ADMIN');


CREATE TABLE books
(
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    author TEXT NOT NULL,
    publisher TEXT NOT NULL,
    total_page INT  NOT NULL,
    published_date  TIMESTAMP WITH TIME ZONE  NOT NULL, 
    price BIGINT NOT NULL,
    stock INT NOT NULL
);

INSERT INTO books(id, name,author, publisher,total_page, published_date, price, stock)
VALUES
(1,'One Last Stop','Casey McQuiston','Tiki', 100,'2019-05-21 00:00:00', 100000,10),
(2,'Ariadne: A Novel','Masahiro Urushido, Michael Anstendig','publisher1', 70,'2019-05-21 00:00:00', 15000,0),
(3,'The Japanese Art of the Cocktail','P. C. Cast, Kristin Cast','Amazon', 80,'2019-05-21 00:00:00', 10000,5),
(4,'The Blacktongue Thief (Blacktongue, 1)','Christopher Buehlman','Hardcover', 20,'2019-05-21 00:00:00', 50000,7),
(5,'The Very Hungry Caterpillar','Eric Carle','Hardcover', 100,'2019-05-21 00:00:00', 70000,0);


CREATE TABLE orders
(
    id SERIAL PRIMARY KEY NOT NULL,
    users_id BIGINT NOT NULL,
    books_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    shipping_details TEXT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE  NOT NULL, 
    CONSTRAINT "orders_users_id_fk" FOREIGN KEY ("users_id") REFERENCES "users"("id"),
    CONSTRAINT "orders_books_id_fk" FOREIGN KEY ("books_id") REFERENCES "books"("id")
);