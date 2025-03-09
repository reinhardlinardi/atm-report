CREATE TABLE `atm_report`.`atm` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `atm_id` VARCHAR(10) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `atm_id_UNIQUE` (`atm_id` ASC) VISIBLE
);
