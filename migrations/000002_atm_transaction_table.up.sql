CREATE TABLE `atm-report`.`atm_transaction` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `atm_id` VARCHAR(10) NOT NULL,
  `transaction_id` VARCHAR(40) NOT NULL,
  `transaction_date` date NOT NULL,
  `transaction_type` TINYINT NOT NULL,
  `amount` INT NOT NULL,
  `card_number` CHAR(16) NOT NULL,
  `destination_card_number` CHAR(16) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`transaction_id` ASC) VISIBLE
);
