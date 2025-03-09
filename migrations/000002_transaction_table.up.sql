CREATE TABLE `atm_report`.`transaction` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `atm_id` VARCHAR(10) NOT NULL,
  `transaction_id` VARCHAR(40) NOT NULL,
  `date` date NOT NULL,
  `type` TINYINT NOT NULL,
  `amount` INT NOT NULL,
  `card_num` CHAR(16) NOT NULL,
  `dest_card_num` CHAR(16) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`transaction_id` ASC) VISIBLE
);
