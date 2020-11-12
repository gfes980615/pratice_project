SELECT COUNT(1) AS `total`,COALESCE(SUM(`tmp`.`transfer_amount)`, 0)  AS `amount` 
FROM( 
    SELECT `level`.`name` AS `level`,
        `deposit`.`id` AS `order_id`,
        `deposit`.`member_login` AS `member_account`,
        `view_member_reseller_agents`.`agent` AS `agent`,
        `view_member_reseller_agents`.`general_agent` AS `general_agent`,
        `view_member_reseller_agents`.`share_login` AS `share_login`,
        `deposit`.`card_bank_name` AS `member_bank_name`,
        `deposit`.`card_name` AS `member_account_name`,
        `deposit`.`transfer_method` AS `member_transfer_method`,
        IFNULL(`deposit`.`transfer_amount`, 0) AS `transfer_amount`,
        IFNULL(`deposit`.`discount_amount`, 0) AS `other_discount_amount`,
        IFNULL(`deposit`.`deposit_amount`, 0) AS `total`,
        `deposit`.`account_bank_name` AS `company_bank_name`,
        `deposit`.`account_name` AS `company_account_name`,
        `deposit`.`account_number` AS `company_account_number`,
        `deposit`.`accept_login` AS `operator`,
        `deposit`.`audit_login` AS `audit`,
        IFNULL(`deposit`.`status`, -99) AS `status`,
        `deposit`.`added_time` AS `added_time`,
        `deposit`.`transfer_time` AS `transfer_time`,
        `deposit`.`accept_time` AS `accept_time`,
        `deposit`.`audit_time` AS `audit_time` 
    FROM `deposit` 
    INNER JOIN `member` ON `member`.`login` = `deposit`.`member_login` 
    INNER JOIN `level` ON `level`.`code` = `member`.`level_code` 
    INNER JOIN `view_member_reseller_agents` ON `view_member_reseller_agents`.`member` = `deposit`.`member_login` 
    WHERE (({choose} = 'added' 
            AND (IFNULL({start_time},-99)=-99 OR `deposit`.`added_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `deposit`.`added_time` < {end_time})) 
        OR ({choose} = 'accept' 
            AND (IFNULL({start_time},-99)=-99 OR `deposit`.`accept_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `deposit`.`accept_time` < {end_time})) 
        OR ({choose} = 'audit' 
            AND (IFNULL({start_time},-99)=-99 OR `deposit`.`audit_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `deposit`.`audit_time` < {end_time}))) 
        AND `deposit`.`status` = 1 
        AND `deposit`.`account_number` = {account} 
        AND (IFNULL({account_name},-99)=-99 OR `deposit`.`account_name` = {account_name}) 
        AND (IFNULL({member_agent},-99)=-99 OR `view_member_reseller_agents`.`agent` = {member_agent}) 
        AND (IFNULL({member_general_agent},-99)=-99 OR `view_member_reseller_agents`.`general_agent` = {member_general_agent}) 
        AND (IFNULL({member_share_login},-99)=-99 OR FIND_IN_SET(`view_member_reseller_agents`.`share_login`, {member_share_login})) 
    UNION ALL 
    SELECT `level`.`name` AS `level`,
        `tmp`.`id` AS `order_id`,
        `tmp`.`member_login` AS `member_account`,
        `view_member_reseller_agents`.`agent` AS `agent`,
        `view_member_reseller_agents`.`general_agent` AS `general_agent`,
        `view_member_reseller_agents`.`share_login` AS `share_login`,
        `member`.`card_bank_name` AS `member_bank_name`,
        `tmp`.`member_account_name` AS `member_account_name`,
        `tmp`.`member_transfer_method` AS `member_transfer_method`,
        IFNULL(`tmp`.`transfer_amount`, 0) AS `transfer_amount`,
        IFNULL(`tmp`.`other_discount_amount`, 0) AS `other_discount_amount`,
        IFNULL(`tmp`.`deposit_amount`, 0) AS `total`,
        `tmp`.`company_bank_name` AS `company_bank_name`,
        `tmp`.`account_name` AS `company_account_name`,
        `tmp`.`account_number` AS `company_account_number`,
        `tmp`.`operator` AS `operator`,
        `tmp`.`audit` AS `audit`,
        `tmp`.`status` AS `status`,
        `tmp`.`added_time` AS `added_time`,
        `tmp`.`transfer_time` AS `transfer_time`,
        `tmp`.`accept_time` AS `accept_time`,
        `tmp`.`audit_time` AS `audit_time` 
    FROM( 
        SELECT `mandeposit`.`id` AS `id`,
            IFNULL(`mandeposit`.`action_code`, '0') AS `action_code`,
            `mandeposit`.`member_login` AS `member_login`,
            `mandeposit`.`audit_login` AS `member_account_name`,
            '人工存入' AS `member_transfer_method`,
            `mandeposit`.`transfer_amount` AS `transfer_amount`,
            `mandeposit`.`discount_amount` AS `other_discount_amount`,
            `mandeposit`.`deposit_amount` AS `deposit_amount`,
            `account`.`name` AS `company_bank_name`,
            `mandeposit`.`account_name` AS `account_name`,
            `mandeposit`.`account_number` AS `account_number`,
            `mandeposit`.`accept_login` AS `operator`,
            `mandeposit`.`audit_login` AS `audit`,
            1 AS `status`,
            `mandeposit`.`added_time` AS `added_time`,
            `mandeposit`.`added_time` AS `transfer_time`,
            `mandeposit`.`accept_time` AS `accept_time`,
            `mandeposit`.`audit_time` AS `audit_time` 
        FROM `mandeposit`,`account` 
        WHERE `mandeposit`.`account_number` = `account`.`number` 
    ) AS `tmp` 
    INNER JOIN `member` ON `member`.`login` = `tmp`.`member_login` 
    INNER JOIN `level` ON `level`.`code` = `member`.`level_code` 
    INNER JOIN `view_member_reseller_agents` ON `view_member_reseller_agents`.`member` = `tmp`.`member_login` 
    WHERE (({choose} = 'added' 
            AND (IFNULL({start_time},-99)=-99 OR `tmp`.`added_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `tmp`.`added_time` < {end_time})) 
        OR ({choose} = 'accept' 
            AND (IFNULL({start_time},-99)=-99 OR `tmp`.`accept_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `tmp`.`accept_time` < {end_time})) 
        OR ({choose} = 'audit' 
            AND (IFNULL({start_time},-99)=-99 OR `tmp`.`audit_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `tmp`.`audit_time` < {end_time}))) 
        AND `tmp`.`account_number` = {account} 
        AND (IFNULL({account_name},-99)=-99 OR `tmp`.`account_name` = {account_name}) 
        AND (IFNULL({member_agent},-99)=-99 OR `view_member_reseller_agents`.`agent` = {member_agent}) 
        AND (IFNULL({member_general_agent},-99)=-99 OR `view_member_reseller_agents`.`general_agent` = {member_general_agent}) 
        AND (IFNULL({member_share_login},-99)=-99 OR FIND_IN_SET(`view_member_reseller_agents`.`share_login`, {member_share_login})) 
        AND IFNULL(`tmp`.`account_name`, '') != '' 
        AND IFNULL(`tmp`.`account_number`, '') != '' 
) AS `tmp` 
 
 
 
 
 
