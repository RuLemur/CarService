create table users
(
    id              serial PRIMARY KEY,
    username        varchar not null,
    registration_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

create table user_car
(
    id              SERIAL PRIMARY KEY,
    user_id         int not null,
    model_id        int not null,
    production_year int not null,
    mileage         int,
    added_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

create table car_models
(
    id          serial PRIMARY KEY,
    brand       varchar not null,
    model       varchar not null,
    equipment   varchar,
    engine_type varchar
);

create table service
(
    id            serial PRIMARY KEY,
    car_id        int     not null,
    service_type  varchar not null,
    price         float,
    details_price float,
    details       varchar,
    note          text,
    service_date  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
