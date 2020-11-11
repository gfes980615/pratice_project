SELECT COUNT(1) as total, COALESCE(SUM(tmp.transfer_amount),0) as amount
FROM(
    SELECT level.name as level,
        deposit.id as order_id,
        deposit.member_login as member_account,
        view_member_agents.agent as agent,
        view_member_agents.general_agent as general_agent,
        view_member_agents.share_login as share_login,
        deposit.card_bank_name as member_bank_name,
        deposit.card_name as member_account_name,
        deposit.transfer_method as member_transfer_method,
        IFNULL(deposit.transfer_amount, 0) as transfer_amount,
        IFNULL(deposit.discount_amount, 0) as other_discount_amount,
        IFNULL(deposit.deposit_amount, 0) as total,
        deposit.account_bank_name as company_bank_name,
        deposit.account_name as company_account_name,
        deposit.account_number as company_account_number,
        deposit.accept_login as operator,
        deposit.audit_login as audit,
        IFNULL(deposit.status, -99) as status,
        deposit.added_time as added_time,
        deposit.transfer_time as transfer_time,
        deposit.accept_time as accept_time,
        deposit.audit_time as audit_time
    FROM deposit,
        member,
        level,
        view_member_agents
    WHERE (({choose} = 'added'
            AND (IFNULL({start_time},-99)=-99 OR deposit.added_time >= {start_time})
            AND (IFNULL({end_time},-99)=-99 OR deposit.added_time < {end_time}))
        OR ({choose} = 'accept'
            AND (IFNULL({start_time},-99)=-99 OR deposit.accept_time >= {start_time})
            AND (IFNULL({end_time},-99)=-99 OR deposit.accept_time < {end_time}))
        OR ({choose} = 'audit'
            AND (IFNULL({start_time},-99)=-99 OR deposit.audit_time >= {start_time})
            AND (IFNULL({end_time},-99)=-99 OR deposit.audit_time < {end_time})))
        AND deposit.status = 1
        AND deposit.member_login = member.login
        AND member.level_code = level.code
        AND deposit.member_login = view_member_agents.member
        AND deposit.account_number = {account}
        AND (IFNULL({account_name},-99)=-99 OR deposit.account_name = {account_name})
        AND (IFNULL({member_agent},-99)=-99 OR view_member_agents.agent = {member_agent})
        AND (IFNULL({member_general_agent},-99)=-99 OR view_member_agents.general_agent = {member_general_agent})
        AND (IFNULL({member_share_login},-99)=-99 OR find_in_set(view_member_agents.share_login, {member_share_login}))
    UNION ALL
    SELECT level.name as level,
        tmp.id as order_id,
        tmp.member_login as member_account,
        view_member_agents.agent as agent,
        view_member_agents.general_agent as general_agent,
        view_member_agents.share_login as share_login,
        member.card_bank_name as member_bank_name,
        tmp.member_account_name as member_account_name,
        tmp.member_transfer_method as member_transfer_method,
        IFNULL(tmp.transfer_amount, 0) as transfer_amount,
        IFNULL(tmp.other_discount_amount, 0) as other_discount_amount,
        IFNULL(tmp.deposit_amount, 0) as total,
        tmp.company_bank_name as company_bank_name,
        tmp.account_name as company_account_name,
        tmp.account_number as company_account_number,
        tmp.operator as operator,
        tmp.audit as audit,
        tmp.status as status,
        tmp.added_time as added_time,
        tmp.transfer_time as transfer_time,
        tmp.accept_time as accept_time,
        tmp.audit_time as audit_time
    FROM(
        SELECT mandeposit.id as id,
            IFNULL(mandeposit.action_code, '0') as action_code,
            mandeposit.member_login as member_login,
            mandeposit.audit_login as member_account_name,
            '人工存入' as member_transfer_method,
            mandeposit.transfer_amount as transfer_amount,
            mandeposit.discount_amount as other_discount_amount,
            mandeposit.deposit_amount as deposit_amount,
            account.name as company_bank_name,
            mandeposit.account_name as account_name,
            mandeposit.account_number as account_number,
            mandeposit.accept_login as operator,
            mandeposit.audit_login as audit,
            1 as status,
            mandeposit.added_time as added_time,
            mandeposit.added_time as transfer_time,
            mandeposit.accept_time as accept_time,
            mandeposit.audit_time as audit_time
        FROM mandeposit, account
        WHERE mandeposit.account_number = account.number
    ) as tmp,
        member,
        level,
        view_member_agents
    WHERE (({choose} = 'added'
            AND (IFNULL({start_time},-99)=-99 OR tmp.added_time >= {start_time})
            AND (IFNULL({end_time},-99)=-99 OR tmp.added_time < {end_time}))
        OR ({choose} = 'accept'
            AND (IFNULL({start_time},-99)=-99 OR tmp.accept_time >= {start_time})
            AND (IFNULL({end_time},-99)=-99 OR tmp.accept_time < {end_time}))
        OR ({choose} = 'audit'
            AND (IFNULL({start_time},-99)=-99 OR tmp.audit_time >= {start_time})
            AND (IFNULL({end_time},-99)=-99 OR tmp.audit_time < {end_time})))
        AND tmp.member_login = member.login
        AND member.level_code = level.code
        AND tmp.member_login = view_member_agents.member
        AND tmp.account_number = {account}
        AND (IFNULL({account_name},-99)=-99 OR tmp.account_name = {account_name})
        AND (IFNULL({member_agent},-99)=-99 OR view_member_agents.agent = {member_agent})
        AND (IFNULL({member_general_agent},-99)=-99 OR view_member_agents.general_agent = {member_general_agent})
        AND (IFNULL({member_share_login},-99)=-99 OR find_in_set(view_member_agents.share_login, {member_share_login}))
        AND IFNULL(tmp.account_name, '') != ''
        AND IFNULL(tmp.account_number, '') != ''
) as tmp