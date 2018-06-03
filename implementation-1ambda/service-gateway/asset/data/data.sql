INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (1, null, '/other', 'other', 'Other', '');

INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (2, null, '/gadget', 'gadget', 'Gadget', '');

INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (3, null, '/computer', 'computer', 'Computer', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (4, 3, '/computer/desktop', 'desktop', 'Desktop', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (5, 3, '/computer/laptop', 'laptop', 'Laptop', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (6, 3, '/computer/printer', 'printer', 'Printer', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (7, 3, '/computer/tablet', 'tablet', 'Tablet', '');

INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (8, null, '/software', 'software', 'Software', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (9, 8, '/software/os', 'os', 'OS', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (10, 8, '/software/office_application', 'office_application', 'Office Application', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (11, 8, '/software/house_application', 'house_application', 'House Application', '');

INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (12, null, '/book', 'book', 'Book', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (13, 12, '/book/computer_science', 'computer_science', 'Computer Science', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (14, 12, '/book/biographical', 'biographical', 'Biographical', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (15, 12, '/book/adventure', 'adventure', 'Adventure', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (16, 12, '/book/science_fiction', 'science_fiction', 'Science Fiction', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (17, 12, '/book/romance', 'romance', 'Romance', '');
INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (18, 12, '/book/police', 'police', 'Police', '');

INSERT INTO `Category` (`id`, `parent_category_id`, `path`, `name`, `display_name`, `description`)
VALUES (19, null, '/food', 'food', 'Food', '');


