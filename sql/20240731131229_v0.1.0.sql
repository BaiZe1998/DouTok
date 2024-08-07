-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id BIGINT PRIMARY KEY,
    mobile VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(64) NOT NULL,
    salt VARCHAR(64) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    INDEX `account_mobile_idx` (mobile),
    INDEX `account_email_idx` (email)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE `users` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `account_id` bigint DEFAULT NULL,
    `mobile` varchar(20) DEFAULT NULL,
    `email` varchar(50) DEFAULT NULL,
    `name` varchar(50) DEFAULT NULL,
    `avatar` varchar(255) DEFAULT NULL,
    `background_image` varchar(255) DEFAULT NULL,
    `signature` varchar(255) DEFAULT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_users_account_id` (`account_id`) USING BTREE,
    UNIQUE KEY `idx_users_email` (`email`) USING BTREE,
    UNIQUE KEY `idx_users_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7240204103809514232 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
