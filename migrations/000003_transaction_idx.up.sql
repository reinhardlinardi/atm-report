ALTER TABLE `atm_report`.`transaction` 
ADD INDEX `transaction_date_type_amount_idx` (`date` ASC, `type` ASC, `amount` DESC) VISIBLE,
ADD INDEX `transaction_type_idx` (`type` ASC) VISIBLE;
