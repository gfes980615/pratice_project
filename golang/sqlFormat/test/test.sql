SELECT COUNT(1) as total, COALESCE(SUM(tmp.transfer_amount),0) as amount
SELECT COUNT(1) as total, COALESCE(SUM(tmp.transfer_amount), 0) as amount