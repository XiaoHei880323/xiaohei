CREATE TABLE `course_error_collection` (
    `id`              int(11)       NOT NULL AUTO_INCREMENT COMMENT '错题ID',
    `course_id`       int(11)       NOT NULL                COMMENT '课程ID，关联course_main.id',
    `question`        text          NOT NULL                COMMENT '错题内容',
    `correct_answer`  text          NOT NULL                COMMENT '正确答案',
    `student_answer`  text          NOT NULL                COMMENT '学生答案',
    `analysis`        text          NOT NULL                COMMENT '解析',
    `knowledge_point` varchar(500)  NOT NULL DEFAULT ''     COMMENT '知识点',
    `create_time`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `add_uid`         int(11)       NOT NULL DEFAULT 0      COMMENT '创建人ID',
    `update_uid`      int(11)       NOT NULL DEFAULT 0      COMMENT '更新人ID',
    `is_delete`       tinyint(3)    NOT NULL DEFAULT 0      COMMENT '是否删除 0:未删除 1:已删除',
    PRIMARY KEY (`id`),
    KEY `idx_err_course_id` (`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='课程错题集表';
