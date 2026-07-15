CREATE TABLE `sy_good` (
    `id`          int(11)        NOT NULL AUTO_INCREMENT COMMENT '编号',
    `good_name`   varchar(100)   NOT NULL DEFAULT ''     COMMENT '商品名',
    `good_img`    varchar(255)   NOT NULL DEFAULT ''     COMMENT '商品图片',
    `good_price`  decimal(10,2)  NOT NULL DEFAULT 0.00   COMMENT '商品价格',
    `good_desc`   longtext                               COMMENT '商品描述（富文本）',
    `status`      tinyint(3)     NOT NULL DEFAULT 1      COMMENT '状态 0:下架 1:上架',
    `add_time`    timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`     int(11)        NOT NULL DEFAULT 0      COMMENT '添加人ID',
    `update_time` timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`  int(11)        NOT NULL DEFAULT 0      COMMENT '修改人ID',
    `is_delete`   tinyint(3)     NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `good_name` (`good_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- 若表已存在（由旧 sy_points_good 改名），执行以下 ALTER 补字段：
-- ALTER TABLE `sy_good`
--     ADD COLUMN `status`    tinyint(3) NOT NULL DEFAULT 1 COMMENT '状态 0:下架 1:上架' AFTER `good_desc`,
--     ADD COLUMN `add_time`  timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间' AFTER `status`,
--     ADD COLUMN `add_uid`   int(11)    NOT NULL DEFAULT 0 COMMENT '添加人ID' AFTER `add_time`,
--     ADD COLUMN `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间' AFTER `add_uid`,
--     ADD COLUMN `update_uid` int(11)   NOT NULL DEFAULT 0 COMMENT '修改人ID' AFTER `update_time`,
--     ADD COLUMN `is_delete` tinyint(3) NOT NULL DEFAULT 0 COMMENT '是否删除' AFTER `update_uid`;
