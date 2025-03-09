ALTER TABLE `atm_report`.`transaction` 
ADD INDEX `transaction_idx` (`date` ASC, `type` ASC, `amount` DESC) VISIBLE;
