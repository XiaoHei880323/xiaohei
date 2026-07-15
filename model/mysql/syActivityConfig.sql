CREATE TABLE `sy_activity_config` (
    `id`            int(11)       NOT NULL AUTO_INCREMENT COMMENT '编号',
    `config_name`   varchar(100)  NOT NULL DEFAULT ''     COMMENT '配置名称',
    `config_image`  varchar(500)  NOT NULL DEFAULT ''     COMMENT '配置图片',
    `start_time`    datetime      NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '开始时间',
    `end_time`      datetime      NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '结束时间',
    `is_default`    tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否默认 0:否 1:是',
    `activity_type` tinyint(3)    NOT NULL DEFAULT 1      COMMENT '活动类型 1:签到活动 2:秒杀活动 3:商品活动 4:景点活动',
    `status`        tinyint(3)    NOT NULL DEFAULT 0      COMMENT '状态 0:下线 1:上线',
    `add_time`      timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`       int(11)       NOT NULL DEFAULT 0      COMMENT '添加人ID，sy_admin.id',
    `update_time`   timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`    int(11)       NOT NULL DEFAULT 0      COMMENT '修改人ID，sy_admin.id',
    `is_delete`     tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `is_default` (`is_default`),
    KEY `activity_type` (`activity_type`),
    KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页活动配置表';
