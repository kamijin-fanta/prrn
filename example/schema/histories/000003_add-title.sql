create table article (
  id bigint not null auto_increment primary key,
  content text not null,
  title varchar(255) not null,
  original_url varchar(255) not null,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);
