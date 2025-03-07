ALTER TABLE `atm-report`.`atm_transaction` 
DROP FOREIGN KEY `atm_id_FK`;

ALTER TABLE `atm-report`.`atm_transaction` 
DROP INDEX `atm_id_FK_idx` ;
