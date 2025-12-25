-- name: GetUserByEmail :one
SELECT * FROM student WHERE email = $1;

-- name: UserExistsByEmail :one
SELECT EXISTS (
  SELECT 1
  FROM student
  WHERE email = $1
);

-- name: GetUserIDByEmail :one
SELECT id
FROM student
WHERE email = $1;