-- name: CreateUser :exec
insert into users (username) values ($1);

-- name: ListUsers :many
select * from users;
