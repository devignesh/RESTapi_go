create table arrays(                                                        
    id serial primary key,
    mobile text[]
);


insert into arrays(mobile) values (array['44444','33333','345345','vicky']);
