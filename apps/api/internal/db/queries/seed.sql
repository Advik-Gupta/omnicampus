-- name: AddDummyStudent :exec
INSERT INTO student (
    id,
    name,
    register_number,
    dob,
    email,
    password,
    phone,
    timetable_id,
    courses_ids
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, NULL, $8
)
ON CONFLICT (email) DO NOTHING;
