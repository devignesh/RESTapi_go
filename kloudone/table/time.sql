create table timedata(
    id serial primary key,
    name varchar(255),
    officetime time,
    offendtime time
);


select localtime;
select localtime(0);