-- name: CreatePerson :one

INSERT INTO person (id, fname, lname, age, email, job, savings, due, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetPerson :one

SELECT * FROM person
WHERE id = $1;

-- name: GetPersons :many

SELECT * FROM person;

-- name: DeletePerson :exec
DELETE FROM person WHERE id = $1;


-- name: UpdatePerson :exec
UPDATE person
SET savings = $2 and updated_at = $3
WHERE id = $1
RETURNING *;
-- name: UpdateDue :exec
UPDATE person
SET due = (
    SELECT SUM(amount)
    FROM creditRequest
    WHERE person_id = person.id
    GROUP BY person_id  
)
WHERE EXISTS (
    SELECT 1
    FROM creditRequest
    WHERE person_id = person.id AND status = 'open'
);
