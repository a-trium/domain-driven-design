-- +migrate Up
INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `category_id`, `image_id`)
VALUES (1, CURRENT_TIMESTAMP(), 'LG Notebook GRAM 13', '1350000', 'OS is not installed', 5, null);
INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `category_id`, `image_id`)
VALUES (2, CURRENT_TIMESTAMP(), 'LG Notebook GRAM 15', '1550000', 'Windows Installed', 5, null);


-- +migrate Down
SELECT 1;

