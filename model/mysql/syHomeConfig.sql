CREATE TABLE `sy_home_config` (
    `id`            int(11)       NOT NULL AUTO_INCREMENT COMMENT '编号',
    `config_name`   varchar(100)  NOT NULL DEFAULT ''     COMMENT '配置名称',
    `config_image`  varchar(500)  NOT NULL DEFAULT ''     COMMENT '配置图片',
    `activity_id`   int(11)       NOT NULL DEFAULT 0      COMMENT '活动ID，sy_activity.id',
    `activity_type` tinyint(3)    NOT NULL DEFAULT 1      COMMENT '活动类型 1:签到活动 2:秒杀活动',
    `sort`          int(11)       NOT NULL DEFAULT 0      COMMENT '排序',
    `add_time`      timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`       int(11)       NOT NULL DEFAULT 0      COMMENT '添加人ID，sy_admin.id',
    `update_time`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`    int(11)       NOT NULL DEFAULT 0      COMMENT '修改人ID，sy_admin.id',
    `is_delete`     tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `activity_id` (`activity_id`),
    KEY `activity_type` (`activity_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页配置表';
