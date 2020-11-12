SELECT COUNT(1) AS `total`,COALESCE(SUM(`tmp`.`transfer_amount)`, AS `amount` 
FROM( 
    SELECT `level`.`name` AS `level`,
        `deposit`.`id` AS `id`,
        `deposit`.`member_login` AS `account`,
        `view_member_agents`.`agent` AS `agent`,
        `view_member_agents`.`general_agent` AS `agent`,
        `view_member_agents`.`share_login` AS `login`,
        `deposit`.`card_bank_name` AS `name`,
        `deposit`.`card_name` AS `name`,
        `deposit`.`transfer_method` AS `method`,
        IFNULL(`deposit`.`transfer_amount`, 0) AS `amount`,
        IFNULL(`deposit`.`discount_amount`, 0) AS `amount`,
        IFNULL(`deposit`.`deposit_amount`, 0) AS `total`,
        `deposit`.`account_bank_name` AS `name`,
        `deposit`.`account_name` AS `name`,
        `deposit`.`account_number` AS `number`,
        `deposit`.`accept_login` AS `operator`,
        `deposit`.`audit_login` AS `audit`,
        IFNULL(`deposit`.`status`, -99) AS `status`,
        `deposit`.`added_time` AS `time`,
        `deposit`.`transfer_time` AS `time`,
        `deposit`.`accept_time` AS `time`,
        `deposit`.`audit_time` AS `audit_time` 
    FROM `deposit`,
        `member`,
        `level`,
        `view_member_agents` 
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
        AND `deposit`.`member_login` = `member`.`login` 
        AND `member`.`level_code` = `level`.`code` 
        AND `deposit`.`member_login` = `view_member_agents`.`member` 
        AND `deposit`.`account_number` = {account} 
        AND (IFNULL({account_name},-99)=-99 OR `deposit`.`account_name` = {account_name}) 
        AND (IFNULL({member_agent},-99)=-99 OR `view_member_agents`.`agent` = {member_agent}) 
        AND (IFNULL({member_general_agent},-99)=-99 OR `view_member_agents`.`general_agent` = {member_general_agent}) 
        AND (IFNULL({member_share_login},-99)=-99 OR FIND_IN_SET(`view_member_agents`.`share_login`, {member_share_login})) 
    UNION ALL 
    SELECT `level`.`name` AS `level`,
        `tmp`.`id` AS `id`,
        `tmp`.`member_login` AS `account`,
        `view_member_agents`.`agent` AS `agent`,
        `view_member_agents`.`general_agent` AS `agent`,
        `view_member_agents`.`share_login` AS `login`,
        `member`.`card_bank_name` AS `name`,
        `tmp`.`member_account_name` AS `name`,
        `tmp`.`member_transfer_method` AS `method`,
        IFNULL(`tmp`.`transfer_amount`, 0) AS `amount`,
        IFNULL(`tmp`.`other_discount_amount`, 0) AS `amount`,
        IFNULL(`tmp`.`deposit_amount`, 0) AS `total`,
        `tmp`.`company_bank_name` AS `name`,
        `tmp`.`account_name` AS `name`,
        `tmp`.`account_number` AS `number`,
        `tmp`.`operator` AS `operator`,
        `tmp`.`audit` AS `audit`,
        `tmp`.`status` AS `status`,
        `tmp`.`added_time` AS `time`,
        `tmp`.`transfer_time` AS `time`,
        `tmp`.`accept_time` AS `time`,
        `tmp`.`audit_time` AS `audit_time` 
    FROM( 
        SELECT `mandeposit`.`id` AS `id`,
            IFNULL(`mandeposit`.`action_code`, '0') AS `code`,
            `mandeposit`.`member_login` AS `login`,
            `mandeposit`.`audit_login` AS `name`,
            '人工存入' AS `method`,
            `mandeposit`.`transfer_amount` AS `amount`,
            `mandeposit`.`discount_amount` AS `amount`,
            `mandeposit`.`deposit_amount` AS `amount`,
            `account`.`name` AS `name`,
            `mandeposit`.`account_name` AS `name`,
            `mandeposit`.`account_number` AS `number`,
            `mandeposit`.`accept_login` AS `operator`,
            `mandeposit`.`audit_login` AS `audit`,
            1 AS `status`,
            `mandeposit`.`added_time` AS `time`,
            `mandeposit`.`added_time` AS `time`,
            `mandeposit`.`accept_time` AS `time`,
            `mandeposit`.`audit_time` AS `audit_time` 
        FROM `mandeposit`,`account` 
        WHERE `mandeposit`.`account_number` = `account`.`number` 
    ) AS `tmp`,
        `member`,
        `level`,
        `view_member_agents` 
    WHERE (({choose} = 'added' 
            AND (IFNULL({start_time},-99)=-99 OR `tmp`.`added_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `tmp`.`added_time` < {end_time})) 
        OR ({choose} = 'accept' 
            AND (IFNULL({start_time},-99)=-99 OR `tmp`.`accept_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `tmp`.`accept_time` < {end_time})) 
        OR ({choose} = 'audit' 
            AND (IFNULL({start_time},-99)=-99 OR `tmp`.`audit_time` >= {start_time}) 
            AND (IFNULL({end_time},-99)=-99 OR `tmp`.`audit_time` < {end_time}))) 
        AND `tmp`.`member_login` = `member`.`login` 
        AND `member`.`level_code` = `level`.`code` 
        AND `tmp`.`member_login` = `view_member_agents`.`member` 
        AND `tmp`.`account_number` = {account} 
        AND (IFNULL({account_name},-99)=-99 OR `tmp`.`account_name` = {account_name}) 
        AND (IFNULL({member_agent},-99)=-99 OR `view_member_agents`.`agent` = {member_agent}) 
        AND (IFNULL({member_general_agent},-99)=-99 OR `view_member_agents`.`general_agent` = {member_general_agent}) 
        AND (IFNULL({member_share_login},-99)=-99 OR FIND_IN_SET(`view_member_agents`.`share_login`, {member_share_login})) 
        AND IFNULL(`tmp`.`account_name`, '') != '' 
        AND IFNULL(`tmp`.`account_number`, '') != '' 
) AS `tmp` 
