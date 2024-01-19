create table if not exists roles (
  `name` varchar(25),
  constraint pk_roles primary key (`name`)
);

create table if not exists users (
  id varchar(100) not null,
  username varchar(255) not null,
  email varchar(255) not null,
  full_name varchar(255),
  status_verification boolean default false,
  `password` varchar(255),
  `role` varchar(25),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp default null,
  constraint pk_users primary key (id, username, email),
  constraint fk_users_role foreign key (`role`) references roles(`name`) on delete cascade
);

create table if not exists courses (
  id int not null auto_increment,
  course_title varchar(255),
  course_image longtext,
  course_content longtext,
  is_free boolean default true,
  is_active boolean default true,
  has_buyer boolean default false,
  price int(255) default 0,
  constraint pk_courses primary key (id)
);

create table if not exists user_has_course (
  user_id varchar(100) not null,
  course_id int not null,
  constraint fk_user_has_course_users foreign key (user_id) references users(id) on delete cascade,
  constraint fk_user_has_course_courses foreign key (course_id) references courses(id) on delete cascade
);

create table if not exists user_cart (
  id varchar(100) not null,
  user_id varchar(100) not null,
  created_at timestamp not null,
  deleted_at timestamp,
  constraint pk_user_cart primary key (id),
  constraint fk_user_cart foreign key (user_id) references users(id) on delete cascade
);

create table if not exists cart_item (
  id bigint not null auto_increment,
  cart_id varchar(100) not null,
  course_id int not null,
  constraint pk_user_cart primary key (id),
  constraint fk_user_cart_cart foreign key (cart_id) references user_cart(id),
  constraint fk_user_cart_course foreign key (course_id) references courses(id)
);

create table if not exists user_transaction (
  id varchar(100) not null,
  user_id varchar(100) not null,
  cart_id varchar(100) not null,
  transaction_code varchar(10) not null,
  transaction_currency varchar(4) not null,
  transaction_total decimal(19,4) not null,
  transaction_method varchar(50) not null,
  transaction_status varchar(50) not null,
  transaction_proof longtext,
  created_at timestamp not null,
  updated_at timestamp not null,
  constraint pk_user_transaction primary key (id),
  constraint fk_user_transaction_user foreign key (user_id) references users(id),
  constraint fk_user_transaction_cart foreign key (cart_id) references user_cart(id)
);