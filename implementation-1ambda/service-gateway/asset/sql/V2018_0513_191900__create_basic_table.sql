-- +migrate Up
CREATE TABLE `User` (
  -- primary key
  `id`         INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at` TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at` TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at` TIMESTAMP            NULL     DEFAULT NULL,
  KEY `idx_USER_deleted_at` (`deleted_at`),

  -- columns
  `password`   TEXT                 NOT NULL,
  `email`      VARCHAR(50)         NOT NULL,
  `phone`      VARCHAR(50)          NOT NULL,
  `name`       VARCHAR(50)          NOT NULL,
  `birthday`   VARCHAR(20)          NOT NULL,
  `address`    TEXT                 NOT NULL,

  CONSTRAINT `uniq_USER_email` UNIQUE (`email`)

  -- FK columns
);

CREATE TABLE `Category` (
  -- primary key
  `id`          INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`  TIMESTAMP            NULL     DEFAULT NULL,
  KEY `idx_Category_deleted_at` (`deleted_at`),

  -- columns
  `name`        VARCHAR(191)         NOT NULL,
  `description` VARCHAR(191)         NOT NULL,

  CONSTRAINT `uniq_Category_name` UNIQUE (`name`)

  -- FK columns
);


CREATE TABLE `Image` (
  -- primary key
  `id`         INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at` TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at` TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at` TIMESTAMP            NULL     DEFAULT NULL,
  KEY `idx_Image_deleted_at` (`deleted_at`),

  -- columns
  `name`       VARCHAR(191)         NOT NULL,
  `type`       VARCHAR(191)         NOT NULL,
  `path`       TEXT                 NOT NULL

  -- FK columns
);

CREATE TABLE `Product` (
  -- primary key
  `id`          INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`  TIMESTAMP            NULL     DEFAULT NULL,
  KEY `idx_Product_deleted_at` (`deleted_at`),

  -- columns
  `name`        VARCHAR(191)         NOT NULL,
  `price`       INTEGER(10) UNSIGNED NOT NULL,
  `detail`      TEXT                 NOT NULL,

  -- FK columns
  `category_id` INTEGER(10) UNSIGNED NOT NULL,
  `image_id`    INTEGER(10) UNSIGNED NOT NULL,

  CONSTRAINT `fk_Product_category_id`
  FOREIGN KEY (`category_id`) REFERENCES `Category` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Product_image_id`
  FOREIGN KEY (`image_id`) REFERENCES `Image` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE `Order` (
  -- primary key
  `id`                INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`        TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`        TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`        TIMESTAMP            NULL     DEFAULT NULL,
  KEY `idx_Order_deleted_at` (`deleted_at`),

  -- columns
  `state`             VARCHAR(30)          NOT NULL,
  `amount`            INTEGER(10) UNSIGNED NOT NULL,

  `shipping_country`  VARCHAR(50)          NOT NULL,
  `shipping_city`     VARCHAR(50)          NOT NULL,
  `shipping_state`    VARCHAR(50)          NOT NULL,
  `shipping_zipcode`  VARCHAR(20)          NOT NULL,
  `shipping_address1` TEXT                 NOT NULL,
  `shipping_address2` TEXT                 NOT NULL,
  `shipping_message`  TEXT                 NOT NULL,

  `orderer_name`      VARCHAR(50)          NOT NULL,
  `orderer_phone`     VARCHAR(50)          NOT NULL,
  `orderer_email`     VARCHAR(50)          NOT NULL,
  `recipient_name`    VARCHAR(50)          NOT NULL,
  `recipient_phone`   VARCHAR(50)          NOT NULL,
  `recipient_email`   VARCHAR(50)          NOT NULL,

  -- FK columns
  `user_id`           INTEGER(10) UNSIGNED NOT NULL,

  CONSTRAINT `fk_Order_user_id`
  FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE OrderDetail (
  -- primary key
  `id`         INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at` TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at` TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at` TIMESTAMP            NULL     DEFAULT NULL,
  KEY `idx_OrderDetail_deleted_at` (`deleted_at`),

  -- columns
  `index`      INTEGER(10) UNSIGNED NOT NULL,
  `price`      INTEGER(10) UNSIGNED NOT NULL,
  `quantity`   INTEGER(10) UNSIGNED NOT NULL,
  `amount`     INTEGER(10) UNSIGNED NOT NULL,

  -- FK columns
  `order_id`   INTEGER(10) UNSIGNED NOT NULL,
  `product_id` INTEGER(10) UNSIGNED NOT NULL,

  CONSTRAINT `fk_OrderDetail_order_id`
  FOREIGN KEY (`order_id`) REFERENCES `Order` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,

  CONSTRAINT `fk_OrderDetail_product_id`
  FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE `User`;
DROP TABLE `Category`;
DROP TABLE `Image`;
DROP TABLE `Product`;
DROP TABLE `Order`;
DROP TABLE `OrderDetail`;

