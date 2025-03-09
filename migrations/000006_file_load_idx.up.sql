ALTER TABLE `atm_report`.`file_load` 
ADD INDEX `file_load_atm_date_idx` (`atm_id` ASC, `date` DESC) VISIBLE,
ADD INDEX `file_load_date_idx` (`date` DESC) VISIBLE;
