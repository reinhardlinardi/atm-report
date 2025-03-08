ALTER TABLE `atm-report`.`transaction` 
ADD INDEX `transaction_date_type_idx` (`transaction_date` DESC, `transaction_type` ASC) VISIBLE;
