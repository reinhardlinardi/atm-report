ALTER TABLE `atm-report`.`file` 
ADD CONSTRAINT `file_atm_id_FK`
  FOREIGN KEY (`atm_id`)
  REFERENCES `atm-report`.`atm` (`atm_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
