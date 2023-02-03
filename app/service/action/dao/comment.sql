create table comment
(
    user_id     bigint unsigned           not null,
    comment_id  bigint unsigned           not null
        primary key,
    content     varchar(255) charset utf8 not null,
    create_date datetime                  not null,
    constraint cmt_user_id
        foreign key (user_id) references user (user_id)
);

