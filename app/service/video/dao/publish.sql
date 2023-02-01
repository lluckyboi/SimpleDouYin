create table publish
(
    title        varchar(255)    not null comment '标题',
    publish_time bigint unsigned not null comment '时间戳 unix',
    user_id      bigint unsigned not null comment '用户id',
    video_id     bigint unsigned not null comment '视频id',
    constraint user_id
        foreign key (user_id) references user (user_id),
    constraint video
        foreign key (video_id) references video (video_id)
);

