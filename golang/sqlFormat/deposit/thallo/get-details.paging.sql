SELECT 
    COUNT(1) AS `total`, 
    COALESCE(SUM(`t`.`transfer_amount`),0) AS `amount`
FROM(
    SELECT  
        IFNULL(`deposit`.`transfer_amount`, 0) AS `transfer_amount`
    FROM  `deposit`
    INNER JOIN `view_member_thallo_agents` ON `view_member_thallo_agents`.`member` = `deposit`.`member_login`
    INNER JOIN `level` ON `level`.`code` = `view_member_thallo_agents`.`user_level_code`       
    WHERE
        (IFNULL({start_time},-99)=-99 OR `deposit`.`accept_time` >= {start_time})
        AND (IFNULL({end_time},-99)=-99 OR `deposit`.`accept_time` < {end_time})
        AND `deposit`.`status` = 1
        AND `deposit`.`account_number` = {account}
        AND (IFNULL({account_name},-99)=-99 OR `deposit`.`account_name` = {account_name})
        AND (IFNULL({member_login}, 1) = 1 OR `view_member_thallo_agents`.`member` IN ({member_login_replace}))
    UNION ALL
    SELECT 
        IFNULL(`tmp`.`transfer_amount`, 0) AS `transfer_amount`
    FROM(
        SELECT
            `mandeposit`.`member_login` AS `member_login`,
            `mandeposit`.`accept_time` AS `accept_time`,
            `mandeposit`.`account_name` AS `account_name`,
            `mandeposit`.`account_number` AS `account_number`,
            `mandeposit`.`transfer_amount` AS `transfer_amount`
        FROM `mandeposit`
        INNER JOIN `account` ON `account`.`number` = `mandeposit`.`account_number`
        INNER JOIN `bank` ON `bank`.`code` = `account`.`bank_code`
    ) AS `tmp`
    INNER JOIN `member` ON `member`.login = `tmp`.`member_login`
    INNER JOIN `level` ON `level`.`code` = `member`.`level_code`
    INNER JOIN `view_member_thallo_agents` ON `view_member_thallo_agents`.`member` = `tmp`.`member_login`
    WHERE 
        (IFNULL({start_time},-99)=-99 OR `tmp`.`accept_time` >= {start_time})
        AND (IFNULL({end_time},-99)=-99 OR `tmp`.`accept_time` < {end_time})
        AND `tmp`.`account_number` = {account}
        AND (IFNULL({account_name},-99)=-99 OR `tmp`.`account_name` = {account_name})
        AND IFNULL(`tmp`.`account_name`, '') != ''
        AND IFNULL(`tmp`.`account_number`, '') != ''
        AND (IFNULL({member_login}, 1) = 1 OR `view_member_thallo_agents`.`member` IN ({member_login_replace}))
) AS `t`




