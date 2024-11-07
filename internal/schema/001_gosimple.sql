-- +goose Up
CREATE TABLE products (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created DATETIME NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE products;

