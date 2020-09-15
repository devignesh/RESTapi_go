-- PROVIDER (Port-1114)

CREATE EXTENSION pglogical;

CREATE TABLE tbl (id int primary key,name varchar);

SELECT pglogical.create_node (
    node_name := 'provider',
    dns := 'host=localhost port=1114 dbname=sample'
);

SELECT pglogical.replication_set_add_table (
    set_name := 'default', relation := 'tbl', synchronize_data := true
);



-- SUBSCRIBER1 (Port-1115)

CREATE EXTENSION pglogical;

CREATE TABLE tbl (id int primary key,name varchar);

SELECT pglogical.create_node (
    node_name := 'subscriber1',
    dsn := 'host=localhost port=1115 dbname=subscribe1'
);

SELECT pglogical.create_subscription (
    subscription_name := 'subscription1',
    provider_dsn := 'host=localhost port=1114 dbname=sample'
);



-- SUBSCRIBER2 (Port-1115)

CREATE EXTENSION pglogical;

CREATE TABLE tbl (id int primaty key, name varchar);

SELECT pglogical.create_node (
    node_name := 'subscriber2',
    dsn := 'host=localhost port=1115 dbname=sub2'
);

SELECT pglogical.create_subscription (
    subscription_name := 'subscribtion2',
    provider_dsn := 'host=localhost port=1114 dbname=sample'
);


