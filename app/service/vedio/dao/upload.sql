# upload表
create table upload
(
    upload_time datetime     not null on update CURRENT_TIMESTAMP comment '上次时间戳',
    title       varchar(255) not null comment '标题',
    user_id     bigint       not null comment '用户id',
    video_id    bigint       not null comment '视频id',
    constraint user_id
        foreign key (user_id) references user (user_id),
    constraint video
        foreign key (video_id) references video (video_id)
);

