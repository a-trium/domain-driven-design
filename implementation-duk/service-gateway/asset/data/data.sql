-- +migrate Up
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (1, '서울', '', '고객1', '1991-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (2, '서울', '', '고객2', '1992-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (3, '서울', '', '고객3', '1993-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (4, '서울', '', '고객4', '1994-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (5, '서울', '', '고객5', '1995-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (6, '서울', '', '고객6', '1996-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (7, '서울', '', '고객7', '1997-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (8, '서울', '', '고객8', '1998-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (9, '서울', '', '고객9', '1999-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (10, '서울', '', '고객10', '1990-01-01');

-- +migrate Down
SELECT 1;