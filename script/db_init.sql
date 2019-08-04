CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签管理';

CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  `image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `deleted_at` dateTime,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

CREATE TABLE `auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `nickname` varchar(100) DEFAULT '' COMMENT '昵称',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `articleId` int(10) unsigned NOT NULL COMMENT '文章id',
  `content` text DEFAULT '' COMMENT '内容',
  `created_on` int(11) DEFAULT NOT NULL COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '评论者昵称',
  `email` varchar(50) DEFAULT '' COMMENT '评论者邮箱',
   `reply_num` int(5) unsigned DEFAULT '0' COMMENT '评论被回复的次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `reply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `commentId` int(10) unsigned NOT NULL COMMENT '评论id',
  `articleId` int(10) unsigned NOT NULL COMMENT '文章id',
  `content` text DEFAULT '' COMMENT '内容',
  `created_on` int(11) DEFAULT NOT NULL COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '回复者昵称',
  `email` varchar(50) DEFAULT '' COMMENT '回复者邮箱',
   `reply_by` varchar(100) DEFAULT '' COMMENT '被回复者昵称',
   `email` varchar(50) DEFAULT '' COMMENT '被回复者邮箱',
   `deleted_at` dateTime,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');