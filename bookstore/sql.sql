-- 数据库中创建users表 --
create table users(
	id int primary key auto_increment,
	username varchar(100) not null unique,
	password varchar(100) not null,email varchar(100)
)CHARSET=utf8;

-- 数据库中创建books表 --
create table books(
	id int primary key auto_increment,
	title varchar(100) not null,
	author varchar(100) not null,
	price double(11,2) not null,
	sales int not null,
	stock int not null,
	img_path varchar(100)
)CHARSET=utf8;

-- 数据库中创建session表 --
create table sessions(
    session_id varchar(100) primary key,    -- 主键id
    username varchar(100) not null,
    user_id int not null,
    foreign key(user_id) references users(id)   -- 
)charset=utf8;

-- 创建carts购物车表
create table carts(
    id varchar(100) primary key,    -- 主键id
    total_count int not null,
    total_amount double(11,2) not null,
    user_id int not null,
    foreign key(user_id) references users(id)   -- 将user_id添加外键关联users表中的id
)charset=utf8;

-- 数据库中创建cart_items购物项表 --
create table cart_items(
    id int primary key auto_increment,  --primary key标识该字段为本表主键唯一标识，auto_increment用作自增
    count int not null,
    amount double(11,2) not null,
    book_id int not null,
    cart_id varchar(100) not null,
    foreign key(book_id) references books(id),  -- 将book_id添加外键关联books表中的id
    foreign key(cart_id) references carts(id)  -- 将cart_id添加外键关联carts表中的id
)charset=utf8;

-- 数据库中创建 orders 订单表 --
create table orders(
    id varchar(100) primary key,
    create_time varchar(100) not null,
    total_count int not null,
    total_amount double(11,2) not null,
    state int not null,
    user_id int,
    foreign key(user_id) references users(id)
)charset=utf8;

-- 数据库中创建 order_items 订单项表 --
create table order_items(
    id int primary key auto_increment,
    count int not null,
    amount double(11,2) not null,
    title varchar(100) not null,
    author varchar(100) not null,
    price double(11,2) not null,
    img_path varchar(100) not null,
    order_id varchar(100) not null,
    foreign key(order_id) references orders(id)
)charset=utf8;