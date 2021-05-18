# prrn

<!-- # Short Description -->

Simple migration tool for MySQL

This is a CLI that helps you create a DB migration file. There is no need to write up and down files from scratch anymore.

<!-- # Badges -->

[![Github issues](https://img.shields.io/github/issues/kamijin-fanta/prrn)](https://github.com/kamijin-fanta/prrn/issues)
[![Github forks](https://img.shields.io/github/forks/kamijin-fanta/prrn)](https://github.com/kamijin-fanta/prrn/network/members)
[![Github stars](https://img.shields.io/github/stars/kamijin-fanta/prrn)](https://github.com/kamijin-fanta/prrn/stargazers)
[![Github top language](https://img.shields.io/github/languages/top/kamijin-fanta/prrn)](https://github.com/kamijin-fanta/prrn/)
[![Github license](https://img.shields.io/github/license/kamijin-fanta/prrn)](https://github.com/kamijin-fanta/prrn/)

# Installation

### Download binary

from https://github.com/kamijin-fanta/prrn/releases


### Docker Image

```shell
docker pull docker.pkg.github.com/kamijin-fanta/prrn/prrn
```

### Via go cli

```shell
go install github.com/kamijin-fanta/prrn
```

# Minimal Example

see [./example](./example) dir.

### 1. Initialize project

```shell
$ prrn init
$ tree
.
└── schema       
    ├── main.sql    # An SQL file that defines a declarative schema, describing CREATE TABLE, etc.
    ├── histories   # The main.sql file is copied when the migration is created.
    └── migrations  # Up and down files that can be executed by the migration tool.
```

### 2. Put schema

```sql
-- schema/main.sql
create table article (
  id bigint not null auto_increment primary key,
  content text not null
);
```

### 3. Make first migration

```shell
$ prrn make --name=init
```

```sql
$ cat schema/migrations/000001_init.sql 
-- +migrate Up
SET FOREIGN_KEY_CHECKS = 0;
CREATE TABLE `article` (
`id` BIGINT (20) NOT NULL AUTO_INCREMENT,
`content` TEXT NOT NULL,
PRIMARY KEY (`id`)
);
SET FOREIGN_KEY_CHECKS = 1;


-- +migrate Down
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE `article`;
SET FOREIGN_KEY_CHECKS = 1;
```

### 4. Edit schema

```sql
-- schema/main.sql
create table article (
  id bigint not null auto_increment primary key,
  content text not null,
  original_url varchar(255) not null,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

### 5. Make second migration

```shell
$ prrn make --name=init
```

```sql
$ cat schema/migrations/000002_add-article-fileds.sql 
-- +migrate Up
SET FOREIGN_KEY_CHECKS = 0;
ALTER TABLE `article` ADD COLUMN `original_url` VARCHAR (255) NOT NULL AFTER `content`;
ALTER TABLE `article` ADD COLUMN `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP AFTER `original_url`;
SET FOREIGN_KEY_CHECKS = 1;


-- +migrate Down
SET FOREIGN_KEY_CHECKS = 0;
ALTER TABLE `article` DROP COLUMN `original_url`;
ALTER TABLE `article` DROP COLUMN `created_at`;
SET FOREIGN_KEY_CHECKS = 1;
```

### 6. Exec migration

Run the migration with the tool of your choice.

recommended migrate tool: https://github.com/rubenv/sql-migrate

⚠ Be sure to check the content of the migration to be performed.
