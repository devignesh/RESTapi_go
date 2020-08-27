create type season as enum ('summer', 'winter', 'spring', 'autumn');

create table enumtype(                                                          
    id serial primary key,
    festivel varchar(255),
    seasons season
);


insert into enumtype(id, festivel, seasons) values (1, 'village fes', 'summer');