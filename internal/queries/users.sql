-- name: CreateUser :exec
insert into users (name) values ($1);

-- name: ListUsers :many
select * from users;
