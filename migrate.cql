create keyspace new_new_relic with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
use new_new_relic;

create table request_info (service_name text, url text, method text, status smallint, response_time int, created_at timestamp, primary key (service_name, created_at, method, url));
