CREATE TABLE `sy_signin_activity_scenic` (
    `id`           int(11)       NOT NULL AUTO_INCREMENT COMMENT '编号',
    `activity_id`  int(11)       NOT NULL DEFAULT 0      COMMENT '签到活动ID，sy_activity.id',
    `scenic_id`    int(11)       NOT NULL DEFAULT 0      COMMENT '景点ID，sy_scenic_spot.id',
    `sign_points`  int(11)       NOT NULL DEFAULT 0      COMMENT '在此景点签到可获积分',
    `qr_code_url`  varchar(500)  NOT NULL DEFAULT ''     COMMENT '打卡二维码地址',
    `status`       tinyint(3)    NOT NULL DEFAULT 1      COMMENT '状态 0:禁用 1:启用',
    `add_time`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `add_uid`      int(11)       NOT NULL DEFAULT 0      COMMENT '添加人ID',
    `update_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_uid`   int(11)       NOT NULL DEFAULT 0      COMMENT '修改人ID',
    `is_delete`    tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `activity_id` (`activity_id`),
    KEY `scenic_id`   (`scenic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='签到活动-景点关联表';

-- 修改商品表：新增富文本描述字段
ALTER TABLE `sy_good`
    ADD COLUMN `good_desc` longtext COMMENT '商品描述（富文本）' AFTER `good_price`;
