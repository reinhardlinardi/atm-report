ALTER TABLE `atm_report`.`history` 
ADD INDEX `history_atm_date_seq_idx` (`atm_id` ASC, `date` ASC, `seq` ASC) VISIBLE;
