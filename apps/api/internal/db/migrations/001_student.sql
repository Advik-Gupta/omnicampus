-- +goose Up
CREATE TABLE student (
  id UUID NOT NULL UNIQUE,
  name TEXT NOT NULL,
  register_number TEXT NOT NULL UNIQUE,
  dob DATE NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  phone TEXT NOT NULL,
  timetable_id UUID,
  courses_ids UUID[] NOT NULL DEFAULT '{}',
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE student;
