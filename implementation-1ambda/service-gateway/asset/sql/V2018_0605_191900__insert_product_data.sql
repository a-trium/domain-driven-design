-- +migrate Up
INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (1, CURRENT_TIMESTAMP(), 'LG Notebook GRAM 13', '1350000', 'OS is not installed', 'Y', 5, null);
INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (2, CURRENT_TIMESTAMP(), 'LG Notebook GRAM 15', '1550000', 'Windows Installed', 'N', 5, null);

INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (1, CURRENT_TIMESTAMP(), 'Memory 8 GB+', '160000', '', 'Y', 1, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (2, CURRENT_TIMESTAMP(), 'SSD 256 GiB+', '100000', '', 'Y', 1, null);

INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (3, CURRENT_TIMESTAMP(), 'Memory 4 GB+', '160000', '', 'N', 2, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (4, CURRENT_TIMESTAMP(), 'SSD 128 GiB+', '100000', '', 'Y', 2, null);

-- +migrate Down
SELECT 1;

