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

-- name: UpdateUserPasswordByID :exec
UPDATE student
SET password = $2
WHERE id = $1;  

-- name: GetStudentOnboardingStatusByEmail :one
SELECT is_onboarded
FROM student
WHERE email = $1;

-- name: SetStudentOnboardedByEmail :exec
UPDATE student
SET is_onboarded = TRUE
WHERE email = $1;

-- name: GetStudentByID :one
SELECT *
FROM student
WHERE id = $1;