create table boolean (
id int primary key, 
active boolean not null 
);


insert into boolean (id, active) values (1, '0');
insert into boolean (id, active) values (2, 'yes');
insert into boolean (id, active) values (3,'1');
insert into boolean (id, active) values (4,'f');
insert into boolean (id, active) values (5,'t');
insert into boolean (id, active) values (6,'y');

