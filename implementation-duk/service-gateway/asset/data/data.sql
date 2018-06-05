-- +migrate Up

-- customer
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (1, '서울', '00001', '김사주', '1991-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (2, '서울', '00002', '대장 가즈아', '1992-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (3, '서울', '00003', '크립토키티', '1993-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (4, '서울', '00004', '고객4', '1994-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (5, '서울', '00005', '고객5', '1995-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (6, '서울', '00006', '고객6', '1996-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (7, '서울', '00007', '고객7', '1997-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (8, '서울', '00008', '고객8', '1998-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (9, '서울', '00009', '고객9', '1999-01-01');
insert into customer(`id`, `address1`, `zip_code`, `name`, `birthday`) values (10, '서울', '00010', '고객10', '1990-01-01');

insert into seller(`id`, `address1`, `zip_code`, `name`) values (1, '서울', '00001', '셀러1');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (2, '서울', '00002', '셀러2');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (3, '서울', '00003', '셀러3');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (4, '서울', '00004', '셀러4');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (5, '서울', '00005', '셀러5');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (100, '판교', '00006', 'SJW');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (101, '??', '00007', '나카모토 사토시');
insert into seller(`id`, `address1`, `zip_code`, `name`) values (102, '러시아', '00008', '비탈릭 부테린');

-- tag
insert into tag(`id`, `name`) values (1, '우리사주');
insert into tag(`id`, `name`) values (2, '휴지');
insert into tag(`id`, `name`) values (3, 'coin');
insert into tag(`id`, `name`) values (4, 'cryptocurrency');
insert into tag(`id`, `name`) values (5, '가즈아');
insert into tag(`id`, `name`) values (6, 'smart contract');

-- product
insert into product(`id`, `name`, `price`, `seller_id`, `image_url`, `on_sale`) values (1, '우리사주', 500, 100, 'http://s3.amazonaws.com/pix.iemoji.com/images/emoji/apple/ios-11/256/pile-of-poo.png', 'Y');
insert into product(`id`, `name`, `price`, `seller_id`, `image_url`, `on_sale`) values (2, 'BTC', 10000000, 101, 'https://cdn.namuwikiusercontent.com/s/92b23f05cdfa39ba1d241ad1177fddcca8000b3ade608f7079251701398a4c7e3ae216e9d8557ce1716cba0d84489f3481ed16475202424aa1612bf4044212abbddaa157c16beaf4b5c5a702aac605f0?e=1530581304&k=X5NLwU0sQEVc44PjXg2i9w', 'Y');
insert into product(`id`, `name`, `price`, `seller_id`, `image_url`, `on_sale`) values (3, 'ETH', 1000000, 102, 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Ethereum_logo_2014.svg/300px-Ethereum_logo_2014.svg.png', 'Y');

insert into product_option(`id`, `product_id`, `name`, `stock`, `price`) values (1, 1, '1차', 800000, 0);
insert into product_option(`id`, `product_id`, `name`, `stock`, `price`) values (2, 1, '2차', 310000, 0);
insert into product_option(`id`, `product_id`, `name`, `stock`, `price`) values (3, 2, 'BTC', 21000000, 0);
insert into product_option(`id`, `product_id`, `name`, `stock`, `price`) values (4, 3, 'ETH', 99235862, 0);

insert into product_tag(`id`, `product_id`, `tag_id`) values (1, 1, 1);
insert into product_tag(`id`, `product_id`, `tag_id`) values (2, 1, 2);
insert into product_tag(`id`, `product_id`, `tag_id`) values (3, 2, 3);
insert into product_tag(`id`, `product_id`, `tag_id`) values (4, 2, 4);
insert into product_tag(`id`, `product_id`, `tag_id`) values (5, 2, 5);
insert into product_tag(`id`, `product_id`, `tag_id`) values (6, 3, 3);
insert into product_tag(`id`, `product_id`, `tag_id`) values (7, 3, 4);
insert into product_tag(`id`, `product_id`, `tag_id`) values (8, 3, 5);
insert into product_tag(`id`, `product_id`, `tag_id`) values (9, 3, 6);

-- cart
insert into cart(`id`, `customer_id`, `option_id`, `quantity`) values (1, 2, 3, 10);
insert into cart(`id`, `customer_id`, `option_id`, `quantity`) values (2, 3, 3, 1);
insert into cart(`id`, `customer_id`, `option_id`, `quantity`) values (3, 3, 4, 100);

-- order
insert into `order`(`id`, `customer_id`) values (1, 1);

insert into order_detail(`id`, `order_id`, `product_id`, `option_id`, `quantity`) values (1, 1, 1, 1, 382);
insert into order_detail(`id`, `order_id`, `product_id`, `option_id`, `quantity`) values (2, 1, 1, 2, 70);

-- discount

-- +migrate Down
SELECT 1;