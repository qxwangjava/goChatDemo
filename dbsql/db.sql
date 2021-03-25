CREATE TABLE `group` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                         `name` varchar(50) NOT NULL COMMENT '群组名称',
                         `group_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '群组头像',
                         `introduction` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '群组简介',
                         `user_num` int(11) NOT NULL DEFAULT '0' COMMENT '群组人数',
                         `type` tinyint(4) NOT NULL COMMENT '群组类型，1：小群；2：大群',
                         `extra` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '附加属性',
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         PRIMARY KEY (`id`)
)  COMMENT='群组';

CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                        `user_name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
                        `nick_name` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',
                        `head_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '头像地址',
                        `birthday` datetime DEFAULT NULL COMMENT '生日',
                        `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
                        `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
                        `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `sex` int(1) DEFAULT NULL COMMENT '性别',
                        `ex` varchar(1024) DEFAULT NULL COMMENT '扩展字段',
                        PRIMARY KEY (`id`)
) COMMENT='用户表';

CREATE TABLE `friend` (
                          `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                          `user_id` bigint(20) NOT NULL COMMENT '用户Id',
                          `friend_id` bigint(20) DEFAULT NULL COMMENT '好友Id',
                          `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                          `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                          `remark` varchar(255) DEFAULT NULL COMMENT '备注',
                          `status` tinyint(255) DEFAULT NULL COMMENT '状态 1-申请 2-同意 3-删除 4-拉黑',
                          PRIMARY KEY (`id`)
) COMMENT='好友表';

CREATE TABLE `group_user` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                              `group_id` bigint(20) unsigned NOT NULL COMMENT '组id',
                              `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
                              `member_type` tinyint(4) NOT NULL COMMENT '成员类型，1：群主；2：管理员；3：普通成员',
                              `remarks` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
                              `extra` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '附加属性',
                              `status` tinyint(255) NOT NULL COMMENT '禁言状态 0 否 1是',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`)
)  COMMENT='群组成员';