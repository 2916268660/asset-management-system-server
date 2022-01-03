-- 创建普通用户表
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` VARCHAR(255) NOT NULL DEFAULT '' UNIQUE KEY COMMENT '唯一用户名',
    `user_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '姓名',
    `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '电话',
    `department` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '部门名称',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `is_admin` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否为管理员',
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

-- 资产信息表
CREATE TABLE `details` (
    `id` BIGINT(3) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `serial_id` VARCHAR(255) NOT NULL DEFAULT '' UNIQUE KEY COMMENT '资产序列号',
    `serial_img` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '序列号生成的二维码路径',
    `category` INT(3) NOT NULL DEFAULT '0' COMMENT '资产品类',
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '名称',
    `status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '资产状态',
    `price` INT(25) NOT NULL DEFAULT '0' COMMENT '资产价格',
    `provide` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '出厂商',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '购进时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT '资产信息表';

-- 资产记录表
CREATE TABLE `record` (
    `id` BIGINT(3) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `task_id` BIGINT(3) NOT NULL DEFAULT '0' COMMENT '任务ID',
    `serial_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '资产序列号',
    `status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '资产的状态。',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `expire_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '到期时间',
    PRIMARY KEY(`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT '资产记录表';

-- 任务单表
CREATE TABLE `task` (
    `id` BIGINT(3) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人账号',
    `user_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人姓名',
    `user_phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人联系方式',
    `department` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人所属部门',
    `category` INT(3) NOT NULL DEFAULT '0' COMMENT '资产品类',
    `nums` INT(10) NOT NULL DEFAULT '0' COMMENT '申请资产数量',
    `days` INT(10) NOT NULL DEFAULT '0' COMMENT '申请使用天数',
    `assets` TEXT(1000) NOT NULL COMMENT '申请资产的序列号json字符串',
    `admin_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '同意领用的管理员账号',
    `admin_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '同意领用的管理员姓名',
    `admin_phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '同意领用的管理员联系方式',
    `provider_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '发放资产人员账号',
    `provider_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '发放资产人员姓名',
    `provider_phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '发放资产人员联系方式',
    `sign_path` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '电子签名生成图片的存储地址',
    `remake` TEXT(1000) NOT NULL COMMENT '备注信息',
    `status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '任务单状态。',
    `property` INT(3) NOT NULL DEFAULT '0' COMMENT '任务单属性',
    `expire_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '到期时间',
    `provide_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '发放时间',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
    `agree_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '审批时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `rollback_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '撤回时间',
    PRIMARY KEY(`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT '任务单表';

-- 维修单表
CREATE TABLE `repairs` (
    `id` BIGINT(3) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人账号',
    `user_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人姓名',
    `user_phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '申请人联系方式',
    `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '维修资产所在地址',
    `assets` TEXT(1000) NOT NULL COMMENT '申请维修资产的序列号json字符串',
    `remake` TEXT(1000) NOT NULL  COMMENT '备注信息',
    `repairer_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '维修人员的联系方式',
    `repairer_phone` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '维修人员的联系方式',
    `status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '维修单的状态',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `receive_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '接单时间',
    `repaired_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '维修好时间',
    `rollback_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '撤回时间',
    PRIMARY KEY(`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT '维修单表';