CREATE TABLE `sy_activity_config_item` (
    `id`            int(11)       NOT NULL AUTO_INCREMENT COMMENT '编号',
    `config_id`     int(11)       NOT NULL DEFAULT 0      COMMENT '活动配置ID，sy_activity_config.id',
    `activity_id`   int(11)       NOT NULL DEFAULT 0      COMMENT '关联活动/商品/景点ID',
    `sort`          int(11)       NOT NULL DEFAULT 0      COMMENT '排序',
    `add_time`      timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`       int(11)       NOT NULL DEFAULT 0      COMMENT '添加人ID，sy_admin.id',
    `update_time`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`    int(11)       NOT NULL DEFAULT 0      COMMENT '修改人ID，sy_admin.id',
    `is_delete`     tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `config_id` (`config_id`),
    KEY `activity_id` (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页活动配置项表';
