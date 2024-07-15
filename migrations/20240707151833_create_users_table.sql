-- +goose Up
CREATE TABLE users (
                       id CHAR(36) PRIMARY KEY,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE users;
