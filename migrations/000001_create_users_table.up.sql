CREATE TABLE IF NOT EXISTS users (
     id BIGSERIAL NOT NULL PRIMARY KEY,
     email VARCHAR (255) NOT NULL UNIQUE,
     name VARCHAR (255) NOT NULL,
     username VARCHAR (255) NOT NULL UNIQUE,
     encrypted_password VARCHAR NOT NULL,
     last_login TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tokens (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    token VARCHAR(255) NOT NULL UNIQUE,
    ttl TIMESTAMP NOT NULL,
    is_alive BOOLEAN NOT NULL DEFAULT FALSE,
    user_id INT,
    CONSTRAINT user_id
        FOREIGN KEY(user_id)
            REFERENCES users(id)
)
-- TODO: created_in/ deleted_in