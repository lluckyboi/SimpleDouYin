create table user
(
    user_id  bigint auto_increment comment 'id',
    username varchar(255) not null,
    password varchar(255) not null,
    name     varchar(50)  null,
    constraint user_user_id_uindex
        unique (user_id)
);

