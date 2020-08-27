create table date(
    id serial primary key,
    emp_name varchar,
    joining_date date
);


create table emp_datetypes(
    emp_id serial primary key,
    f_name varchar(255),
    l_name varchar(255),
    dob date,
    joined_date date default current_date
);

select current_date;
select_now()::date;
select to_char(now() :: date, 'dd/mm/yyyy');

select to_char(now() :: date, 'Mon dd, yyyy');

