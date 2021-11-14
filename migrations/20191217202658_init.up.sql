create table car_services
(
    id            SERIAL PRIMARY KEY,
    user_id       bigint         not null,
    user_car_id   bigint         not null,
    service_name  varchar        NOT NULL,
    service_price numeric(12, 4) NOT NULL,
    service_date  TIMESTAMP      NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP

);
create table cars
(
    id    serial primary key,
    mark  varchar not null,
    model varchar not null
);
create table garage
(
    id              serial PRIMARY KEY,
    model_id        bigint  NOT NULL,
    user_id         bigint  NOT NULL,
    car_name        varchar NOT NULL,
    mileage         bigint,
    production_year smallint

);
create table users
(
    id       serial PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);