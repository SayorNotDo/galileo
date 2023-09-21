SELECT "q_basic_event"."$uid" "virtual_user_id"
FROM (SELECT "sq1".*
      FROM (SELECT *
            FROM (SELECT "#event_name", "#event_time", "$uid"
                  FROM ib.dt_test_tesla.tv_ods_all_event) "sq0") "sq1") "q_basic_event"

ORDER BY