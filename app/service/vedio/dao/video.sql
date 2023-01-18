# video表
create table video
(
    video_id       bigint        not null comment 'vedio唯一标识'
        primary key,
    play_url       varchar(255)  not null comment '播放地址',
    cover_url      varchar(255)  not null comment '封面地址',
    favorite_count int default 0 not null comment '点赞数',
    comment_count  int           not null comment '评论数'
);

