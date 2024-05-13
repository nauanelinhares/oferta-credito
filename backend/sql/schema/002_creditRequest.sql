-- +goose up

CREATE TABLE creditRequest (
    id UUID PRIMARY KEY,
    person_id UUID NOT NULL,
    start_amount FLOAT NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    reason TEXT NOT NULL,
    note TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE creditRequest;
