ALTER TABLE `atm-report`.`atm_transaction` 
ADD INDEX `atm_id_FK_idx` (`atm_id` ASC) VISIBLE;

ALTER TABLE `atm-report`.`atm_transaction` 
ADD CONSTRAINT `atm_id_FK`
  FOREIGN KEY (`atm_id`)
  REFERENCES `atm-report`.`atm` (`atm_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
