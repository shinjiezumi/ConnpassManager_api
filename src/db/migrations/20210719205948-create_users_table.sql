
-- +migrate Up
CREATE TABLE IF NOT EXISTS `connpass_manager`.`users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar (255) NOT NULL COMMENT 'ユーザー名',
    `email` varchar (255) NOT NULL COMMENT 'メールアドレス',
    `email_verified_at` timestamp NULL COMMENT 'メールアドレス認証日時',
    `password` varchar(255) NOT NULL COMMENT 'パスワード',
    `remember_token` varchar(100) NULL COMMENT 'パスワード再設定トークン',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(id),
    UNIQUE KEY(email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

-- +migrate Down
DROP TABLE IF EXISTS `connpass_manager`.`users`;
