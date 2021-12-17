--创建普通用户表
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_name` VARCHAR(255) NOT NULL DEFAULT '' UNIQUE KEY COMMENT '唯一用户名',
    `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '电话',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户表';