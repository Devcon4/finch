SELECT
   pg_terminate_backend (pg_stat_activity.pid)
FROM
   pg_stat_activity
WHERE
   pg_stat_activity.datname = '<dbname>';

drop DATABASE if exists <dbname>