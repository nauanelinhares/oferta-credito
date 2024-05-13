-- name: CreateCreditRequest :one

INSERT INTO creditRequest (id, person_id, start_amount, amount, status, reason, note, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetCreditRequest :one

SELECT * FROM creditRequest
WHERE id = $1;

-- name: GetCreditRequests :many
SELECT * FROM creditRequest;

-- name: UpdateCreditRequest :exec
UPDATE creditRequest
SET status = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCreditRequest :exec
DELETE FROM creditRequest WHERE id = $1;