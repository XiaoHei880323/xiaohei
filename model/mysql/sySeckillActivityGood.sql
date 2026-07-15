CREATE TABLE `sy_seckill_activity_good` (
    `id`            int(11)         NOT NULL AUTO_INCREMENT COMMENT '编号',
    `activity_id`   int(11)         NOT NULL DEFAULT 0       COMMENT '秒杀活动ID，sy_activity.id',
    `good_id`       int(11)         NOT NULL DEFAULT 0       COMMENT '商品ID，sy_good.id',
    `seckill_price` decimal(10,2)   NOT NULL DEFAULT 0.00    COMMENT '该活动中的秒杀价格',
    `add_time`      timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`       int(11)         NOT NULL DEFAULT 0       COMMENT '添加人ID，sy_admin.id',
    `update_time`   timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`    int(11)         NOT NULL DEFAULT 0       COMMENT '修改人ID，sy_admin.id',
    `is_delete`     tinyint(3)      NOT NULL DEFAULT 0       COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `activity_id` (`activity_id`),
    KEY `good_id`     (`good_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='秒杀活动商品表';
