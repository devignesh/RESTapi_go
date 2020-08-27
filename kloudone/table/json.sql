create table json(                                                                                 
    id serial primary key,
    name varchar,
    data json
);


insert into json(data) values ('{ "customer": "vicky", "items": {"product": "jim beam","qty": 22}}');