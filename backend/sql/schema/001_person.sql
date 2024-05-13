-- +goose Up

CREATE TABLE person (
    id UUID PRIMARY KEY,
    fname TEXT NOT NULL,
    lname TEXT NOT NULL,
    age INTEGER NOT NULL,
    email TEXT NOT NULL,
    job TEXT NOT NULL,
    savings FLOAT NOT NULL,
    due FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
); 

-- +goose Down
DROP TABLE person;