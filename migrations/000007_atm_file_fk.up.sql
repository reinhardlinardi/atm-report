ALTER TABLE `atm-report`.`atm_file` 
ADD CONSTRAINT `atm_file_id_FK`
  FOREIGN KEY (`atm_id`)
  REFERENCES `atm-report`.`atm` (`atm_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
