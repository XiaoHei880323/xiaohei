CREATE TABLE `sy_notice` (
    `id`              int(11)       NOT NULL AUTO_INCREMENT COMMENT '编号',
    `notice_name`     varchar(200)  NOT NULL DEFAULT ''     COMMENT '公告名称/标题',
    `notice_time`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '公告发布时间',
    `notice_content`  longtext                              COMMENT '公告内容（富文本）',
    `add_uid`         int(11)       NOT NULL DEFAULT 0      COMMENT '添加人ID，sy_admin.id',
    `publish_uid`     int(11)       NOT NULL DEFAULT 0      COMMENT '发布人ID，sy_admin.id',
    `add_time`        timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `update_time`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`      int(11)       NOT NULL DEFAULT 0      COMMENT '修改人ID，sy_admin.id',
    `is_delete`       tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `notice_time` (`notice_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告表';
