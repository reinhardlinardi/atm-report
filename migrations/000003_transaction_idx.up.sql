ALTER TABLE `atm_report`.`transaction` 
ADD INDEX `transaction_idx` (`transaction_date` DESC, `transaction_type` ASC, `amount` DESC) VISIBLE;
