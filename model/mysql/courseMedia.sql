CREATE TABLE `course_media` (
    `id`          int(11)       NOT NULL AUTO_INCREMENT COMMENT '资源ID',
    `course_id`   int(11)       NOT NULL                COMMENT '课程ID，关联course_main.id',
    `media_type`  tinyint(3)    NOT NULL DEFAULT 0      COMMENT '类型 0:学生上传 1:老师上传 3:上课录音',
    `media_url`   varchar(1024) NOT NULL DEFAULT ''     COMMENT '资源地址',
    `create_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `add_uid`     int(11)       NOT NULL DEFAULT 0      COMMENT '创建人ID，sy_admin.id',
    `update_uid`  int(11)       NOT NULL DEFAULT 0      COMMENT '更新人ID，sy_admin.id',
    `is_delete`   tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `idx_media_course_id` (`course_id`),
    KEY `idx_media_type` (`media_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课程图片|录音资源表';
