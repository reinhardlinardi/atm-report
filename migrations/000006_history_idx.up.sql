ALTER TABLE `atm_report`.`history` 
ADD INDEX `history_atm_date_idx` (`atm_id` ASC, `date` ASC) VISIBLE,
ADD INDEX `history_date_idx` (`date` ASC) VISIBLE;
