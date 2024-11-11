-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
	id integer primary key,
	username text unique not null,
	password_hash text not null,
	created_at datetime not null,
	updated_at datetime not null,
	deleted_at datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
