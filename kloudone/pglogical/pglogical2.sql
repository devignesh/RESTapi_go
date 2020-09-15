-- Provider running on port:1111

CREATE DATABASE {pglogicaldb};

CREATE EXTENSION pglogical;

CREATE TABLE test(id INT PRIMARY KEY,name TEXT,job TEXT);

SELECT pglogical.create_node(
    node_name := 'vicky_clu1',
    dsn := 'host=localhost port=1111 dbname=pglogicaldb'
);

select pglogical.replication_set_add_table(set_name := 'default', relation := 'pglogicaldb' , synchronize_data := 'true');

INSERT INTO test VALUES (1,'john','Automation testing');

INSERT INTO test VALUES (2,'Ilyas','Data admin');
 
INSERT INTO test VALUES (3,'Riyaz','Java developer');
 
INSERT INTO test VALUES (4,'Junaid','Automation testing');
 
INSERT INTO test VALUES (5,'Deepak','Phyton developer');
 
INSERT INTO test VALUES (6,'Wasim','Devops');

// Subscriber running on port:1112

CREATE DATABASE pglogicaldb;

CREATE EXTENSION pglogical;

CREATE TABLE test(id INT PRIMARY KEY,name TEXT,job TEXT);

SELECT pglogical.create_node(
    node_name := 'vicky_clu2',
    dsn := 'host=localhost port=1112 dbname=pglogicaldb'
);

SELECT pglogical.create_subscription(
    subscription_name := 'vicky_clu1',
    provider_dsn := 'host=locahost port=1111 dbname=pglogicaldb'
);

SELECT * FROM test;