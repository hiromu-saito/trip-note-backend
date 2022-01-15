CREATE TABLE IF NOT EXISTS users (
  id        bigint(20)     PRIMARY KEY          comment 'ユーザーID'
 ,email     varchar(200)   DEFAULT NULL UNIQUE  comment 'メールアドレス'
 ,password  longblob                             comment 'パスワード'
);

create table if not exists `trip-note`.`memories` (
    id int primary key comment 'ID'
  , user_id int not null comment 'ユーザID'
  , hotel_name varchar(30) not null comment 'ホテル名称'
  , hotel_image varchar(100) not null comment 'ホテル写真'
  , impression varchar(30) comment '感想'
  , accommodation_date date not null comment '宿泊日'
  , delete_flag int default 0 not null comment '削除フラグ:0:有効,1:無効'
)