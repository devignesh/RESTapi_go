-- // vicky_clu1 running on port:1111

create DATABASE students;

CREATE EXTENSION pglogical;

CREATE TABLE students(id INT PRIMARY KEY,name TEXT,address TEXT,city TEXT);

SELECT pglogical.create_node(node_name := 'new_prov',dsn := 'host=localhost port=1111 dbname=students');

SELECT pglogical.replication_set_add_table(set_name := 'default', relation :='students', synchronize_data := true, columns :='{id,name,city}', row_filter := 'id<5');

INSERT INTO students VALUES(1,'Karan','Besant Nagar','Chennai');

INSERT INTO students VALUES(2,'Kundan','Mount Road','Chennai');

INSERT INTO students VALUES(3,'Sumeet','Anna Nagar','Chennai');

INSERT INTO students VALUES(4,'Bharath','Anna Nagar','Chennai');

INSERT INTO students VALUES(5,'Selvan','Arumbakkam','Chennai');

INSERT INTO students VALUES(6,'Jai','East Coast Road','Chennai');

INSERT INTO students VALUES(7,'Atheeq','OMR','Chennai');

INSERT INTO students VALUES(8,'Abhijith','OMR','Chennai');

SELECT * FROM students;

--  Subscriber running on port:1113

CREATE EXTENSION Pglogical;

CREATE TABLE students(id INT PRIMARY KEY,name TEXT,address TEXT,city TEXT);

SELECT pglogical.create_node(node_name := 'new_sub',dsn := 'host=localhost port=1113 dbname=students');

SELECT pglogical.create_subscription( subscription_name := 'new_subp',provider_dsn := 'host=localhost port=1111 dbname=students');

SELECT * FROM students;