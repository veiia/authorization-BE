CREATE TABLE users (
    id bigserial not null primary key,
    email varchar(255) not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    encrypted_password varchar not null
)