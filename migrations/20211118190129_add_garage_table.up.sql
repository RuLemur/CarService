alter table users
    rename column car_id to garage_id;

create table garage
(
    id       serial PRIMARY KEY,
    model_id int not null
);