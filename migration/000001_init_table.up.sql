SET NAMES utf8mb4;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '账号',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT 'email',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '标题',
  `desc` varchar(50) NOT NULL DEFAULT '' COMMENT '描述',
  `content` varchar(50) NOT NULL DEFAULT '' COMMENT '内容',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_by` int(10) NOT NULL DEFAULT '0' COMMENT '创建人',
  `updated_by` int(10) NOT NULL DEFAULT '0' COMMENT '修改人',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';