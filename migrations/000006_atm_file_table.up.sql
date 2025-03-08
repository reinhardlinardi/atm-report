CREATE TABLE `atm-report`.`atm_file` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `atm_id` VARCHAR(10) NOT NULL,
  `last_file_date` DATE,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `atm_id_UNIQUE` (`atm_id` ASC) VISIBLE
);
