create table deptt(
    dept_no int constraint deptt_pk primary key,
    dept_name text NOT NULL,  
    Location varchar(15)
);


create table student(                                                      
    Stu_id int, 
    Stu_Name text, 
    Stu_Age int, 
    Stu_address char(30)
);