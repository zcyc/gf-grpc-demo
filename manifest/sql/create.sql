create table "user"
(
    id serial primary key,
    passport  varchar(45) not null unique,
    password  varchar(45) not null,
    nickname  varchar(45) not null,
    create_at timestamp,
    update_at timestamp,
    delete_at timestamp
);
