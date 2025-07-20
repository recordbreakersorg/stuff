-- name: GetFileById :one
select * from files
where id = $1;

-- name: CreateFile :one
insert into files (mime, file_size)
values ($1, $2)
returning *;
