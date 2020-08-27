create table chartype (
    id serial primary key, 
    name char(255),
    job varchar(15),
    company text
);


insert into chartype (id, name, job, company) values (1,'vicky', 'SDE1', 'Deftouch');