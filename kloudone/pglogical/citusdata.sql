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


-- Coordinate with master node -- 1111

CREATE TABLE companies (
    
    id bigint NOT NULL,
    
    name text NOT NULL,
    
    image_url text,
    
    created_at timestamp without time zone NOT NULL,
    
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE campaigns (
    
    id bigint NOT NULL,
    
    company_id bigint NOT NULL,
    
    name text NOT NULL,
    
    cost_model text NOT NULL,
    
    state text NOT NULL,
    
    monthly_budget bigint,
    
    blacklisted_site_urls text[],
    
    created_at timestamp without time zone NOT NULL,
    
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE ads (
    
    id bigint NOT NULL,
    
    company_id bigint NOT NULL,
    
    campaign_id bigint NOT NULL,
    
    name text NOT NULL,
    
    image_url text,
    
    target_url text,
    
    impressions_count bigint DEFAULT 0,
    
    clicks_count bigint DEFAULT 0,
    
    created_at timestamp without time zone NOT NULL,
    
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE companies ADD PRIMARY KEY (id);

ALTER TABLE campaigns ADD PRIMARY KEY (id, company_id);

ALTER TABLE ads ADD PRIMARY KEY (id, company_id);


-- citus replication model :

SET citus.replication_model = 'streaming';


-- Distribute the tables:

SELECT create_distributed_table('companies', 'id');

SELECT create_distributed_table('campaigns', 'company_id');

SELECT create_distributed_table('ads', 'company_id');


-- Load the data from CSV file:

\copy companies from '../devignesh/Downloads/companies.csv' with csv

\copy campaigns from '../devignesh/Downloads/campaigns.csv' with csv

\copy ads from '../devignesh/Downloads/ads.csv' with csv
