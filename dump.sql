CREATE TABLE `gorm`.`users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `gorm`.`comments` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` VARCHAR(45) NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_comments_1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_comments_1`
    FOREIGN KEY (`user_id`)
    REFERENCES `gorm`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

INSERT INTO `gorm`.`users` (`name`, `phone`) VALUES ('u1', 'p1');
INSERT INTO `gorm`.`users` (`name`, `phone`) VALUES ('u2', 'p2');

INSERT INTO `gorm`.`comments` (`content`, `user_id`) VALUES ('u1-c1', '1');
INSERT INTO `gorm`.`comments` (`content`, `user_id`) VALUES ('u1-c2', '1');
INSERT INTO `gorm`.`comments` (`content`, `user_id`) VALUES ('u1-c3', '1');
