CREATE TABLE `atm_report`.`history` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `atm_id` VARCHAR(10) NOT NULL,
  `date` DATE NOT NULL,
  `seq` TINYINT NOT NULL,
  PRIMARY KEY (`id`)
);
