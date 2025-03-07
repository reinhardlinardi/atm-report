ALTER TABLE `atm-report`.`atm_transaction` 
ADD INDEX `transaction_date_type_idx` (`transaction_date` DESC, `transaction_time` ASC) VISIBLE;
