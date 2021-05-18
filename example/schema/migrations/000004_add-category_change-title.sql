-- +migrate Up
SET FOREIGN_KEY_CHECKS = 0;
ALTER TABLE `article` ADD COLUMN `category` VARCHAR (255) NOT NULL AFTER `original_url`;
ALTER TABLE `article` CHANGE COLUMN `title` `title` VARCHAR (100) NOT NULL;
SET FOREIGN_KEY_CHECKS = 1;


-- +migrate Down
SET FOREIGN_KEY_CHECKS = 0;
ALTER TABLE `article` DROP COLUMN `category`;
ALTER TABLE `article` CHANGE COLUMN `title` `title` VARCHAR (255) NOT NULL;
SET FOREIGN_KEY_CHECKS = 1;