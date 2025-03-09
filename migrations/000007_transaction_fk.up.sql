ALTER TABLE `atm_report`.`transaction` 
ADD CONSTRAINT `transaction_atm_id_fk`
  FOREIGN KEY (`atm_id`)
  REFERENCES `atm_report`.`atm` (`atm_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
