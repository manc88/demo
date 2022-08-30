CREATE TABLE IF NOT EXISTS users (
    uid INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(255),
    email varchar(255),
    age int,
    deleted bool DEFAULT FALSE
 );
