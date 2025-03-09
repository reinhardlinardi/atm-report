ALTER TABLE `atm-report`.`file` 
ADD INDEX `file_atm_date_idx` (`atm_id` ASC, `file_date` DESC) VISIBLE,
ADD INDEX `file_date_idx` (`file_date` DESC) VISIBLE;
