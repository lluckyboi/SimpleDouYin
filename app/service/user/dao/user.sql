# create table user
# (
#     uid  bigint unsigned auto_increment comment 'uid'primary key,
#     user_true_name varchar(36) default ''     not null comment '用户真名',
#     user_nick_name varchar(36) default 'user' not null comment '昵称',
#     sex            enum ('男', '女')            null comment '性别',
#     user_school    varchar(36) default ''     not null comment '学校名',
#     user_stno      varchar(36) default ''     null comment '学工号',
#     user_role      enum ('教师', '学生')          not null comment '身份: 教师或学生',
#     user_pnum      varchar(16) default ''     not null comment '手机号',
#     user_password  varchar(36) default ''     not null comment '用户密码',
#     creatTime      datetime                   null on update CURRENT_TIMESTAMP comment '创建时间',
#     updateTime     datetime                   null on update CURRENT_TIMESTAMP comment '修改时间',
#     primary key (uid)
# )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
create table gzdemo.user
(
    user_true_name varchar(36) default ''     not null comment '用户真名',
    user_nick_name varchar(36) default 'user' not null comment '昵称',
    sex            enum ('男', '女')            null comment '性别',
    user_school    varchar(36) default ''     not null comment '学校名',
    user_stno      varchar(36) default ''     null comment '学工号',
    user_role      enum ('教师', '学生')          not null comment '身份: 教师或学生',
    user_pnum      varchar(16) default ''     not null comment '手机号',
    user_password  varchar(36) default ''     not null comment '用户密码',
    creatTime      datetime                   null on update CURRENT_TIMESTAMP comment '创建时间',
    updateTime     datetime                   null on update CURRENT_TIMESTAMP comment '修改时间',
    constraint user_pnum
        unique (user_pnum)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

alter table gzdemo.user
    add primary key (user_pnum);

