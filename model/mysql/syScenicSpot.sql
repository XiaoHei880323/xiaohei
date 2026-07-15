CREATE TABLE `sy_scenic_spot` (
    `id`          int(11)        NOT NULL AUTO_INCREMENT COMMENT '编号',
    `spot_name`   varchar(100)   NOT NULL DEFAULT ''     COMMENT '景点名称',
    `longitude`   decimal(10,6)  NOT NULL DEFAULT 0.000000 COMMENT '经度',
    `latitude`    decimal(10,6)  NOT NULL DEFAULT 0.000000 COMMENT '纬度',
    `ticket_price` decimal(10,2) NOT NULL DEFAULT 0.00   COMMENT '票价',
    `description` longtext                               COMMENT '景点描述（富文本）',
    `add_time`    timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`     int(11)        NOT NULL DEFAULT 0       COMMENT '添加人ID',
    `update_time` timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`  int(11)        NOT NULL DEFAULT 0       COMMENT '修改人ID',
    `is_delete`   tinyint(3)     NOT NULL DEFAULT 0       COMMENT '是否删除 0:未删除 1:已删除',
    `status`      tinyint(3)     NOT NULL DEFAULT 1       COMMENT '状态 0:下架 1:上架',
    PRIMARY KEY (`id`),
    KEY `spot_name` (`spot_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='景点表';
