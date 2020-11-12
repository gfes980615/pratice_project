SELECT COUNT(1) AS `total`,COALESCE(SUM(`tmp`.`transfer_amount)`, 0) AS `amount` 
