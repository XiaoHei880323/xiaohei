CREATE TABLE `course_evaluation` (
    `id`          int(11)       NOT NULL AUTO_INCREMENT COMMENT '评价ID',
    `course_id`   int(11)       NOT NULL                COMMENT '课程ID，关联course_main.id',
    `eval_type`   tinyint(3)    NOT NULL DEFAULT 0      COMMENT '类型 0:学生对老师 1:家长对老师 3:老师对学生',
    `content`     text          NOT NULL                COMMENT '评价内容',
    `rating`      tinyint(3)    NOT NULL DEFAULT 5      COMMENT '评分 1-5',
    `create_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `add_uid`     int(11)       NOT NULL DEFAULT 0      COMMENT '创建人ID',
    `update_uid`  int(11)       NOT NULL DEFAULT 0      COMMENT '更新人ID',
    `is_delete`   tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `idx_eval_course_id` (`course_id`),
    KEY `idx_eval_type` (`eval_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课程评价表';
