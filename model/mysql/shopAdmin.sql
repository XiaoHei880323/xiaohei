CREATE TABLE `shop_admin` (
                              `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
                              `userName` varchar(50) COLLATE utf8_german2_ci NOT NULL DEFAULT '' COMMENT '用户名称',
                              `password` varchar(100) COLLATE utf8_german2_ci NOT NULL DEFAULT '' COMMENT '密码',
                              `nickname` varchar(50) COLLATE utf8_german2_ci DEFAULT NULL COMMENT '别名',
                              `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否有效  0：无效 1：有效，',
                              PRIMARY KEY (`id`),
                              KEY `userName` (`userName`),
                              KEY `password` (`password`),
                              KEY `status` (`status`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_german2_ci COMMENT='管理员账号';
