# video表
create table video
(
    video_id       bigint unsigned             not null comment 'video唯一标识'
        primary key,
    play_url       varchar(255)                not null comment '播放地址',
    cover_url      varchar(255)                not null comment '封面地址',
    favorite_count bigint unsigned default '0' not null comment '点赞数',
    comment_count  bigint unsigned default '0' not null comment '评论数'
);

