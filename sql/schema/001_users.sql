-- +goose Up
CREATE table users (
  id uuid primary key default gen_random_uuid(),
  name text not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);

-- +goose Down
DROP TABLE users;
