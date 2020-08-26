create database employ;

CREATE TABLE employee
(
    emp_id INTEGER NOT NULL,
    birth_date DATE,
    email CHAR(20) ,
    emp_name CHAR(30),
    join_date DATE,
    phone CHAR(10),
    CONSTRAINT employee_pkey PRIMARY KEY (emp_id)
);

//Insert Data

INSERT into employee(emp_id, birth_date, email, emp_name, join_date, phone) VALUES (2, '09-01-1997', 'test@gmail.com', 'vicky', '03-01-2019', '9876543210')

select * from employee;



SELECT COUNT(emp_id) FROM employee;SELECT emp_name,phone FROM employee;

UPDATE employee SET phone = 9047660920 WHERE emp_id = 2;


//alter

alter table employee alter column email set default false;


alter table employee alter column email drop default;



//stored procedure


create or replace procedure emp("emp_id" integer, "birth_date" date, "email" character(20), "emp_name" character(20), "join_date" date, "phone" character(10), "created_at" timestamp without time zone)
LANGUAGE SQL
AS $$
INSERT INTO public."employee" VALUES ("emp_id", "birth_date", "email", "emp_name", "join_date", "phone", "created_at");
$$;


CALL emp(3,'1-1-2020', 'vv@g.com', 'test', '2-2-2222', '3333333');



//table 


create table sample
(
    sno integer NOT NULL,
    siid integer,sname varchar(20),sd date,ed date,sid integer,status boolean DEFAULT false,CONSTRAINT pk_snoa PRIMARY KEY (sno));

);


INSERT INTO SAMPLE (sno, siid, sd, ed, sid, status) VALUES (1,1,'2013-04-04','2013-04-04',2,'f' );



CREATE PROCEDURE sample_insert(_sno integer, _siid integer, _sd date, _ed date, _sid integer, _status boolean)
LANGUAGE SQL
AS $BODY$
INSERT INTO SAMPLE(sno, siid, sd, ed, sid, status)
VALUES(_sno, _siid, _sd, _ed, _sid, _status);
$BODY$;


CALL sample_insert (3, 103, '1993-12-24', '1995-12-02', 4, 't');



Autonomous Transactions:
----------------------


CREATE OR REPLACE PROCEDURE test()
LANGUAGE plpgsql
AS $$
DECLARE
BEGIN
CREATE TABLE tnew1 (id int);
INSERT INTO tnew1 VALUES (1);
COMMIT;
CREATE TABLE tnew2 (id int);
INSERT INTO tnew2 VALUES (1);
ROLLBACK;
END $$;


call test();



Displaying a message on the screen 
------------------------------------


CREATE OR REPLACE PROCEDURE display_message (INOUT msg TEXT)
AS $$
BEGIN
RAISE NOTICE 'Procedure Parameter: %', msg ;
END;
$$LANGUAGE plpgsql;


call display_message('Hi vicky');





//Ms SQL :


Dtabaase:

CREATE DATABASE customer;
GO

CREATE DATABASE custom;
GO


USE custom
GO


CREATE TABLE Inventory (id INT, name NVARCHAR(50), quantity INT)
GO

INSERT INTO Inventory VALUES (1, 'banana', 150); INSERT INTO Inventory VALUES (2, 'orange', 154);
GO

SELECT * FROM Inventory WHERE quantity > 152;
GO


//SP

CREATE PROCEDURE cus
AS
SELECT * FROM custom
GO

EXEC cus;
GO



//Triggers and Procedure


CREATE OR REPLACE FUNCTION log_last_name_changes()
RETURNS TRIGGER 
LANGUAGE PLPGSQL
AS
$$
BEGIN
IF NEW.last_name <> OLD.last_name THEN
INSERT INTO employee_audits(employee_id,last_name,changed_on)
VALUES(OLD.id,OLD.last_name,now());
END IF;
RETURN NEW;
END;
$$;



CREATE TRIGGER last_name_changes
BEFORE UPDATE
ON employees
FOR EACH ROW
EXECUTE PROCEDURE log_last_name_changes();


INSERT INTO employees (first_name, last_name) VALUES ('1', 'vicky', 'kumar');
INSERT INTO employees (id, first_name, last_name) VALUES ('2', 'test kloud', 'manage');


UPDATE employees SET last_name = 'micky' WHERE ID = 2;


SELECT * FROM employees;
SELECT * FROM employee_audits;


// function and procedure


CREATE FUNCTION function1(INOUT p1 TEXT) 
AS $$
BEGIN
RAISE NOTICE 'Function Parameter: %', p1;
END;
$$
LANGUAGE plpgsql;

SELECT function1('vicky kloud');




//mssql


USE vickytest
GO

CREATE PROCEDURE GetEmployees
WITH ENCRYPTION
AS
BEGIN
SET NOCOUNT ON
SELECT Empid, empname from emp
END
GO


sp_helptext GetEmployees
GO




Renaming the stored procedure:
------------------------------


sp_remname 'procedure name','new procedure name'


//Store Procedure

CREATE OR REPLACE PROCEDURE genre_traverse() LANGUAGE plpgsql SECURITY DEFINER
AS $$

DECLARE
genre_rec record;
BEGIN
for genre_rec in (select "GenreId","Name" from public."Genre" order by "GenreId")
loop
RAISE NOTICE 'Genre Id is : % , Name is : %', genre_rec."GenreId",genre_rec."Name";
end loop;
END;
$$ ;



create or replace procedure accout(
sender int,
receiver int,
amount dec
)
language plpgsql
as $$
begin
update accounts
set balance = balance + amount
where id = receiver;
update accounts
set balance = balance - amount
where id = sender;
commit;
end;$$
;



CREATE OR REPLACE PROCEDURE datestyle_change()
LANGUAGE plpgsql
SET datestyle TO postgres, dmy
AS $$
begin
raise notice 'Current Date is : % ', now();
end;
$$;


call datestyle_change();



