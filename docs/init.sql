DROP database if exists face_score;
create database face_score;
use face_score;
CREATE TABLE if not exists user (
  id         int unsigned primary key auto_increment,
  open_id    varchar(255) not null unique,
  user_name  varchar(255) not null,
  created_on int          not null
)
  engine = innodb
  default charset = utf8
  comment = '用户表';
create table if not exists job (
  id          int unsigned primary key auto_increment,
  user_id     int unsigned not null,
  file_id     int unsigned not null,
  score       int unsigned,
  created_on  int          not null,
  finished_on int,
  visible     bool                     default true
)
  engine = innodb
  default charset = utf8
  comment = '任务表';
CREATE TABLE if not exists comment (
  id         int unsigned primary key auto_increment,
  user_id    int unsigned  not null,
  job_id     int           not null,
  content    varchar(1023) not null,
  created_on int           not null,
  reply_for  int
)
  engine = innodb
  default charset = utf8
  comment = '评论表';
create table if not exists file (
  id         int unsigned primary key auto_increment,
  user_id    int unsigned not null,
  name       varchar(255) not null,
  md5        varchar(255) not null,
  uri        varchar(255) not null,
  created_on int          not null
)
  engine = innodb
  default charset = utf8
  comment = '文件信息表';

-- # 插入几条测试数据
-- insert into user values (null, "user1_openid", "user1", unix_timestamp());
-- insert into user values (null, "user2_openid", "user2", unix_timestamp());
-- insert into user values (null, "user3_openid", "user3", unix_timestamp());
--
-- insert into file values (null, 1, "file1_name", "file1_md5", "file1_uri", unix_timestamp());
-- insert into file values (null, 2, "file1_name", "file2_md5", "file2_uri", unix_timestamp());
-- insert into file values (null, 3, "file1_name", "file3_md5", "file3_uri", unix_timestamp());
-- insert into file values (null, 4, "file1_name", "file4_md5", "file4_uri", unix_timestamp());
--
-- insert into job values (null, 1, 1, null, unix_timestamp(), null, true);
-- insert into job values (null, 2, 2, 65, unix_timestamp(), unix_timestamp() + 1200, true);
-- insert into job values (null, 3, 3, 60, unix_timestamp(), unix_timestamp() + 1220, true);
-- insert into job values (null, 1, 4, 90, unix_timestamp(), unix_timestamp() + 1230, true);
--
-- insert into comment values (null, 1, 1, "content", unix_timestamp(), null);
-- insert into comment values (null, 1, 2, "content", unix_timestamp(), 1);
-- insert into comment values (null, 3, 3, "content", unix_timestamp(), 1);
--
--
