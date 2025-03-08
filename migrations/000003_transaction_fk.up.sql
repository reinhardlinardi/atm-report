ALTER TABLE `atm-report`.`transaction` 
ADD CONSTRAINT `transaction_atm_id_FK`
  FOREIGN KEY (`atm_id`)
  REFERENCES `atm-report`.`atm` (`atm_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
