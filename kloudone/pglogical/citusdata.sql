-- Citus_Data Implementation

-- Download sample Data

1   curl https://examples.citusdata.com/tutorial/companies.csv > companies.csv 

2   curl https://examples.citusdata.com/tutorial/campaigns.csv > campaigns.csv

3   curl https://examples.citusdata.com/tutorial/ads.csv > ads.csv


-- start the cluster

/usr/lib/postgresql/12/bin/pg_ctl -D /home/devignesh/postgres/vicky_clu1 start
 
/usr/lib/postgresql/12/bin/pg_ctl -D /home/devignesh/postgres/vicky_clu3 start

/usr/lib/postgresql/12/bin/pg_ctl -D /home/devignesh/postgres/vicky_clu5 start

-- Citus data

psql -p {cluster port number} -d {dbname} -U {username}

-- Create Extension

create EXTENSION citus;

-- adding node

SELECT * from master_add_node('{connection ip or localhost}', {portnumber});
SELECT * from master_add_node('{connection ip or localhost}', {portnumber});
select * from master_get_active_worker_nodes();


