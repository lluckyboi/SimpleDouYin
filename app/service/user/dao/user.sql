create table user
(
    user_id        bigint auto_increment comment 'id',
    username       varchar(255)     not null,
    password       varchar(255)     not null,
    name           varchar(50)      null,
    follow_count   bigint default 0 not null comment '关注总数',
    follower_count bigint default 0 not null comment '粉丝总数',
    constraint user_user_id_uindex
        unique (user_id)
);

