SELECT t.team_id,
       t.team_name,
       ifnull(k.num_points, 0) AS num_points
FROM teams AS `t`
LEFT JOIN
  (SELECT m.team_id AS team_id,
          sum(m.num_points) AS num_points
   FROM
     (SELECT host_team AS `team_id`,
             '3' AS `num_points`
      FROM matches
      WHERE host_goals > guest_goals
      UNION SELECT guest_team AS `team_id`,
                   '3' AS `num_points`
      FROM matches
      WHERE guest_goals > host_goals
      UNION SELECT host_team AS `team_id`,
                   '1' AS `num_points`
      FROM matches
      WHERE guest_goals = host_goals
      UNION SELECT guest_team AS `team_id`,
                   '1' AS `num_points`
      FROM matches
      WHERE guest_goals = host_goals ) AS `m`
   GROUP BY m.team_id) AS `k` ON t.team_id = k.team_id
ORDER BY num_points DESC,
         team_id ASC;
