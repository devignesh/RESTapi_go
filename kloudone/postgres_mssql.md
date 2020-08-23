Database creation:

postgres=# create database employ;

Table:

employ=# CREATE TABLE employee
employ-# (
employ(#    emp_id INTEGER NOT NULL,
employ(#    birth_date DATE,
employ(#    email CHAR(20) ,    
employ(#    emp_name CHAR(30),
employ(#    join_date DATE,
employ(#    phone CHAR(10),
employ(#    CONSTRAINT employee_pkey PRIMARY KEY (emp_id)
employ(# );

viewing table :

employ=# \d
              List of relations
 Schema |     Name     |   Type   |  Owner   
--------+--------------+----------+----------
 public | employee     | table    | postgres
 public | users        | table    | postgres
 public | users_id_seq | sequence | postgres
(3 rows)


employ=# \d employee
        Table "public.employee"
   Column   |     Type      | Modifiers 
------------+---------------+-----------
 emp_id     | integer       | not null
 birth_date | date          | 
 email      | character(20) | 
 emp_name   | character(30) | 
 join_date  | date          | 
 phone      | character(10) | 
Indexes:
    "employee_pkey" PRIMARY KEY, btree (emp_id)


employ=# select * from employee;
 emp_id | birth_date |        email         |            emp_name            | join_date  |   phone    
--------+------------+----------------------+--------------------------------+------------+------------
      1 | 1987-10-10 | abc@gmail.com        | Sameer Rawat                   | 2015-01-03 | 9876543210
      2 | 1997-09-01 | test@gmail.com       | vicky                          | 2019-01-03 | 9876543210
(2 rows)

employ=#

employ=# SELECT COUNT(emp_id) FROM employee;SELECT emp_name,phone FROM employee;
 count 
-------
     2
(1 row)

            emp_name            |   phone    
--------------------------------+------------
 Sameer Rawat                   | 9876543210
 vicky                          | 9876543210
(2 rows)

employ=# 


Update values:

employ=# UPDATE employee SET phone = 9047660920 WHERE emp_id = 2;
UPDATE 1

employ=# select * from employee 
employ-# ;
 emp_id | birth_date |        email         |            emp_name            | join_date  |   phone    
--------+------------+----------------------+--------------------------------+------------+------------
      1 | 1987-10-10 | abc@gmail.com        | Sameer Rawat                   | 2015-01-03 | 9876543210
      2 | 1997-09-01 | test@gmail.com       | vicky                          | 2019-01-03 | 9047660920
(2 rows)

employ=# 

// Default value to column:

alter table employee alter column email set default false;

employ=# alter table employee alter column email set default false;
ALTER TABLE
employ=# \d employee
                 Table "public.employee"
   Column   |            Type             |   Modifiers   
------------+-----------------------------+---------------
 emp_id     | integer                     | not null
 birth_date | date                        | 
 email      | character(20)               | default false
 emp_name   | character(30)               | 
 join_date  | date                        | 
 phone      | character(10)               | 
 created_at | timestamp without time zone | 
Indexes:
    "employee_pkey" PRIMARY KEY, btree (emp_id)

employ=#

//Remove default value:

alter table employee alter column email drop default;


//stored procedure

employ=# create or replace procedure emp("emp_id" integer, "birth_date" date, "email" character(20), "emp_name" character(20), "join_date" date, "phone" character(10), "created_at" timestamp without time zone)
LANGUAGE SQL
AS $$
INSERT INTO public."employee" VALUES ("emp_id", "birth_date", "email", "emp_name", "join_date", "phone", "created_at");
$$;
CREATE PROCEDURE

CALL emp(3,'1-1-2020', 'vv@g.com', 'test', '2-2-2222', '3333333');

CALL

employ=# create or replace procedure emp("emp_id" integer, "birth_date" date, "email" character(20), "emp_name" character(20), "join_date" date, "phone" character(10))
LANGUAGE SQL
AS $$
INSERT INTO public."employee" VALUES ("emp_id", "birth_date", "email", "emp_name", "join_date", "phone");
$$;
CREATE PROCEDURE
employ=# CALL emp(3,'1-1-2020', 'vv@g.com', 'test', '2-2-2222', '3333333');
ERROR:  duplicate key value violates unique constraint "employee_pkey"
DETAIL:  Key (emp_id)=(3) already exists.
CONTEXT:  SQL function "emp" statement 1
employ=# CALL emp(4,'1-1-2020', 'vv@g.com', 'test', '2-2-2222', '3333333');
CALL
employ=# 
employ=# select * from employee 
employ-# ;
 emp_id | birth_date |        email         |            emp_name            | join_date  |   phone    
--------+------------+----------------------+--------------------------------+------------+------------
      1 | 1987-10-10 | abc@gmail.com        | Sameer Rawat                   | 2015-01-03 | 9876543210
      2 | 1987-10-10 | abc@gmail.com        | Sameer Rawat                   | 2015-01-03 | 9876543210
      3 | 1987-10-10 | abc@gmail.com        | Sameer Rawat                   | 2015-01-03 | 9876543210
    102 | 1988-02-03 | av@gmail.com         | Arvind Verma                   | 2014-02-03 | 1234567890
      4 | 2020-01-01 | vv@g.com             | test                           | 2222-02-02 | 3333333   
(5 rows)

employ=# 



employ=# create table sample
employ-# (
employ(# sno integer NOT NULL,
employ(# siid integer,sname varchar(20),sd date,ed date,sid integer,status boolean DEFAULT false,CONSTRAINT pk_snoa PRIMARY KEY (sno));
CREATE TABLE
employ=# INSERT INTO SAMPLE (sno, siid, sd, ed, sid, status)
employ-# VALUES (1,1,'2013-04-04','2013-04-04',2,'f' );
INSERT 0 1
employ=# INSERT INTO SAMPLE (sno, siid, sd, ed, sid, status) VALUES (2, 2, '1993-12-24', '1995-12-02', 4, 't');
INSERT 0 1
employ=# CREATE PROCEDURE sample_insert(_sno integer, _siid integer, _sd date, _ed date, _sid integer, _status boolean)
employ-# LANGUAGE SQL
employ-# 
employ-# AS $BODY$
employ$# INSERT INTO SAMPLE(sno, siid, sd, ed, sid, status)
employ$# VALUES(_sno, _siid, _sd, _ed, _sid, _status);
employ$# $BODY$;
CREATE PROCEDURE
employ=# CALL sample_insert (3, 103, '1993-12-24', '1995-12-02', 4, 't');
CALL
employ=# select * from sample;
 sno | siid | sname |     sd     |     ed     | sid | status 
-----+------+-------+------------+------------+-----+--------
   1 |    1 |       | 2013-04-04 | 2013-04-04 |   2 | f
   2 |    2 |       | 1993-12-24 | 1995-12-02 |   4 | t
   3 |  103 |       | 1993-12-24 | 1995-12-02 |   4 | t
(3 rows)

employ=#


Autonomous Transactions:
----------------------

Procedures allow you to define autonomous transactions like COMMIT or ROLLBACK within the procedure

employ=# CREATE OR REPLACE PROCEDURE test() 
employ-# LANGUAGE plpgsql
employ-#  AS $$
employ$#  DECLARE
employ$# BEGIN
employ$# CREATE TABLE tnew1 (id int);
employ$# INSERT INTO tnew1 VALUES (1);
employ$#  COMMIT;
employ$# CREATE TABLE tnew2 (id int);
employ$# INSERT INTO tnew2 VALUES (1);
employ$#  ROLLBACK;
employ$# END $$;
CREATE PROCEDURE
employ=# call test();
CALL
employ=# select * from tnew1 ;
 id 
----
  1
(1 row)

employ=# 



Displaying a message on the screen 
------------------------------------


employ=# CREATE OR REPLACE PROCEDURE display_message (INOUT msg TEXT)
employ-# AS $$
employ$# BEGIN
employ$# RAISE NOTICE 'Procedure Parameter: %', msg ;
employ$# END ;
employ$# $$
employ-# LANGUAGE plpgsql ;
CREATE PROCEDURE
employ=# call display_message('Hi vicky');
NOTICE:  Procedure Parameter: Hi vicky
   msg    
----------
 Hi vicky
(1 row)








//MS SQL 

----------------------------------

1> CREATE DATABASE customer;
2> GO
Msg 1801, Level 16, State 3, Server devignesh, Line 1
Database 'customer' already exists. Choose a different database name.
1> CREATE DATABASE custom;
2> GO
1> USE custom
2> GO
Changed database context to 'custom'.
1> CREATE TABLE Inventory (id INT, name NVARCHAR(50), quantity INT)
2> GO
1> INSERT INTO Inventory VALUES (1, 'banana', 150); INSERT INTO Inventory VALUES (2, 'orange', 154);
2> GO

(1 rows affected)

(1 rows affected)
1> SELECT * FROM Inventory WHERE quantity > 152;
2> GO


//Stored procedure


1> CREATE PROCEDURE cus
2> AS
3> SELECT * FROM custom
4> GO
1> EXEC cus;
2> GO




//Triggers and Procedure



employ=# CREATE OR REPLACE FUNCTION log_last_name_changes()
employ-#   RETURNS TRIGGER 
employ-#   LANGUAGE PLPGSQL  
employ-#   AS
employ-# $$
employ$# BEGIN
employ$# IF NEW.last_name <> OLD.last_name THEN
employ$# 
ABORT           COMMENT         DISCARD         GRANT           NOTIFY          REVOKE          START           WITH
ALTER           COMMIT          DO              IMPORT          PREPARE         ROLLBACK        TABLE           
ANALYZE         COPY            DROP            INSERT          REASSIGN        SAVEPOINT       TRUNCATE        
BEGIN           CREATE          END             LISTEN          REFRESH         SECURITY LABEL  UNLISTEN        
CHECKPOINT      DEALLOCATE      EXECUTE         LOAD            REINDEX         SELECT          UPDATE          
CLOSE           DECLARE         EXPLAIN         LOCK            RELEASE         SET             VACUUM          
CLUSTER         DELETE FROM     FETCH           MOVE            RESET           SHOW            VALUES          
employ$#  INSERT INTO employee_audits(employee_id,last_name,changed_on)
employ$# 
ABORT           COMMENT         DISCARD         GRANT           NOTIFY          REVOKE          START           WITH
ALTER           COMMIT          DO              IMPORT          PREPARE         ROLLBACK        TABLE           
ANALYZE         COPY            DROP            INSERT          REASSIGN        SAVEPOINT       TRUNCATE        
BEGIN           CREATE          END             LISTEN          REFRESH         SECURITY LABEL  UNLISTEN        
CHECKPOINT      DEALLOCATE      EXECUTE         LOAD            REINDEX         SELECT          UPDATE          
CLOSE           DECLARE         EXPLAIN         LOCK            RELEASE         SET             VACUUM          
CLUSTER         DELETE FROM     FETCH           MOVE            RESET           SHOW            VALUES          
employ$#  VALUES(OLD.id,OLD.last_name,now());
employ$# END IF;
employ$# 
employ$# RETURN NEW;
employ$# END;
employ$# $$
employ-# ;
CREATE FUNCTION
employ=# CREATE TRIGGER last_name_changes
employ-#   BEFORE UPDATE
employ-#   ON employees
employ-#   FOR EACH ROW
employ-#   EXECUTE PROCEDURE log_last_name_changes();
CREATE TRIGGER
employ=# INSERT INTO employees (first_name, last_name)
employ-# VALUES ('vicky', 'kumar');
ERROR:  null value in column "id" violates not-null constraint
DETAIL:  Failing row contains (null, vicky, kumar).
employ=# INSERT INTO employees (id, first_name, last_name)
VALUES ('1', 'vicky', 'kumar');
INSERT 0 1
employ=# INSERT INTO employees (id, first_name, last_name)
VALUES ('1', 'test kloud', 'manage');
ERROR:  duplicate key value violates unique constraint "employees_pkey"
DETAIL:  Key (id)=(1) already exists.
employ=# INSERT INTO employees (id, first_name, last_name)
VALUES ('2', 'test kloud', 'manage');
INSERT 0 1
employ=# 
employ=# 
employ=# 
employ=# 
employ=# \d employee
employee         employee_audits  employee_pkey    employees        employees_pkey   
employ=# \d employee_audits 
              Table "public.employee_audits"
   Column    |              Type              | Modifiers 
-------------+--------------------------------+-----------
 id          | integer                        | 
 employee_id | integer                        | not null
 last_name   | character varying(40)          | not null
 changed_on  | timestamp(6) without time zone | not null

employ=# select * from employee_audits 
employ-# ;
 id | employee_id | last_name | changed_on 
----+-------------+-----------+------------
(0 rows)

employ=# SELECT * FROM employees;
 id | first_name | last_name 
----+------------+-----------
  1 | vicky      | kumar
  2 | test kloud | manage
(2 rows)

employ=# UPDATE employees
employ-# SET last_name = 'Brown'
employ-# WHERE ID = 2;
UPDATE 1
employ=# SELECT * FROM employees;
 id | first_name | last_name 
----+------------+-----------
  1 | vicky      | kumar
  2 | test kloud | Brown
(2 rows)

employ=# SELECT * FROM employee_audits;
 id | employee_id | last_name |         changed_on         
----+-------------+-----------+----------------------------
    |           2 | manage    | 2020-08-23 00:15:54.739615






