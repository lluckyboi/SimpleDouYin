create table user
(
    user_id        bigint           not null comment 'id',
    username       varchar(255)     not null,
    password       varchar(255)     not null,
    name           varchar(50)      null,
    follow_count   bigint default 0 not null comment '关注总数',
    follower_count bigint default 0 not null comment '粉丝总数',
    constraint user_user_id_uindex
        unique (user_id)
);

create table follow
(
    uid        BIGINT not null comment '关注者id',
    target_uid BIGINT not null comment '被关注者id'
)
    comment '关注表';
