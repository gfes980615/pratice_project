SELECT  
    `statistics_financial`.`code` AS `account_number`, 
    `statistics_financial`.`name` AS `account_name`, 
    SUM(`statistics_financial`.`amount)` AS `amount`,
    COUNT(DISTINCT(`statistics_financial`.`member_login))` AS `people_count` 
FROM 
    `statistics_financial`,`view_member_agents` 
WHERE (({choose} = 'added'  
        AND (IFNULL({start_time},-99)=-99 OR `statistics_financial`.`accept_time` >= {start_time})  
        AND (IFNULL({end_time},-99)=-99 OR `statistics_financial`.`accept_time` < {end_time})) 
    OR ({choose} = 'accept'  
        AND (IFNULL({start_time},-99)=-99 OR `statistics_financial`.`accept_time` >= {start_time})  
        AND (IFNULL({end_time},-99)=-99 OR `statistics_financial`.`accept_time` < {end_time}))     
    OR ({choose} = 'audit'  
        AND (IFNULL({start_time},-99)=-99 OR `statistics_financial`.`accept_time` >= {start_time})  
        AND (IFNULL({end_time},-99)=-99 OR `statistics_financial`.`accept_time` < {end_time}))) 
    AND `statistics_financial`.`category` = 'deposit' 
    AND `statistics_financial`.`type` = 'deposit' 
    AND `statistics_financial`.`member_login` = `view_member_agents`.`member` 
    AND (IFNULL({member_agent},-99)=-99 OR `view_member_agents`.`agent` = {member_agent}) 
    AND (IFNULL({member_general_agent},-99)=-99 OR `view_member_agents`.`general_agent` = {member_general_agent}) 
    AND (IFNULL({member_share_login},-99)=-99 OR FIND_IN_SET(`view_member_agents`.`share_login`, {member_share_login})) 
GROUP BY `statistics_financial`.`code` `statistics_financial`.`name` 
 
