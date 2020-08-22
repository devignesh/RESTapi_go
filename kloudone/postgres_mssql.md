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



