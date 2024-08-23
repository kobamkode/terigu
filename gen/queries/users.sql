-- name: CreateUser :exec
insert into users (
	name, path
) values (
	$1, $2
);

-- name: ListUsers :many
select * from users;
