-- name: ListAuthors :many
SELECT * FROM authors;

-- name: GetAuthor :one
SELECT * FROM authors 
WHERE id = ?;

-- name: CreateAuthor :exec
INSERT INTO authors (id, name, bio) 
VALUES (?,?,?);

-- name: UpdateAuthor :exec
UPDATE authors SET name = ?, bio = ? 
WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;

-- name: CreateBook :exec
INSERT INTO books (id, name, description, authors_id, price)
VALUES (?,?,?,?,?);

-- name: ListBooks :many
SELECT b.*, a.name as author_name 
FROM books b JOIN authors a ON b.authors_id = a.id;