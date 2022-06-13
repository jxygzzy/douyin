/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : heart

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 13/06/2022 10:13:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `video_id` int(0) NOT NULL,
  `user_id` int(0) NOT NULL,
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_date` datetime(0) NOT NULL,
  `delete_at` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `t_comment_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `t_comment_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_comment
-- ----------------------------
INSERT INTO `t_comment` VALUES (1, 18, 1, '评论一下，视频很不错', '2022-06-04 14:23:47', 0);
INSERT INTO `t_comment` VALUES (3, 18, 1, '再来评论一下，视频真的很不错', '2022-06-04 15:27:27', 0);
INSERT INTO `t_comment` VALUES (4, 18, 2, '好可爱的猫猫', '2022-06-06 12:53:27', 0);
INSERT INTO `t_comment` VALUES (5, 19, 1, '猫猫', '2022-06-06 13:50:57', 0);
INSERT INTO `t_comment` VALUES (6, 29, 2, 'lalala', '2022-06-06 14:22:50', 0);
INSERT INTO `t_comment` VALUES (7, 30, 1, 'hello', '2022-06-09 14:47:54', 1);
INSERT INTO `t_comment` VALUES (8, 30, 1, 'hello', '2022-06-09 14:49:35', 0);

-- ----------------------------
-- Table structure for t_favorite
-- ----------------------------
DROP TABLE IF EXISTS `t_favorite`;
CREATE TABLE `t_favorite`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `video_id` int(0) NOT NULL,
  `user_id` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `t_favorite_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `t_video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `t_favorite_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '点赞表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_favorite
-- ----------------------------
INSERT INTO `t_favorite` VALUES (3, 18, 2);
INSERT INTO `t_favorite` VALUES (4, 23, 1);
INSERT INTO `t_favorite` VALUES (5, 20, 1);
INSERT INTO `t_favorite` VALUES (6, 19, 1);
INSERT INTO `t_favorite` VALUES (7, 18, 1);
INSERT INTO `t_favorite` VALUES (8, 24, 1);
INSERT INTO `t_favorite` VALUES (9, 27, 2);
INSERT INTO `t_favorite` VALUES (10, 26, 2);
INSERT INTO `t_favorite` VALUES (12, 25, 2);
INSERT INTO `t_favorite` VALUES (13, 30, 4);
INSERT INTO `t_favorite` VALUES (14, 30, 1);
INSERT INTO `t_favorite` VALUES (15, 29, 1);
INSERT INTO `t_favorite` VALUES (16, 22, 1);
INSERT INTO `t_favorite` VALUES (17, 21, 1);

-- ----------------------------
-- Table structure for t_relation
-- ----------------------------
DROP TABLE IF EXISTS `t_relation`;
CREATE TABLE `t_relation`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `user_id` int(0) NOT NULL,
  `to_user_id` int(0) NOT NULL COMMENT '被关注者id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `to_user_id`(`to_user_id`) USING BTREE,
  CONSTRAINT `t_relation_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `t_relation_ibfk_2` FOREIGN KEY (`to_user_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '关系表，关注与被关注' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_relation
-- ----------------------------
INSERT INTO `t_relation` VALUES (6, 4, 2);
INSERT INTO `t_relation` VALUES (7, 4, 1);
INSERT INTO `t_relation` VALUES (8, 1, 4);
INSERT INTO `t_relation` VALUES (13, 1, 2);

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `follower_count` int(0) NOT NULL DEFAULT 0,
  `follow_count` int(0) NOT NULL DEFAULT 0,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://pic2.zhimg.com/80/v2-b951b9e3e65678a93c7f2d6184b81337_720w.jpg',
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://pic4.zhimg.com/80/v2-2501a93fe6acebfc81f8c1c825df2367_720w.jpg',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '什么都无法舍弃的人，什么都无法改变。',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (1, 'jxygzzy', '1a69e2dc480768b4e7e80f94ae332651', '抖声用户dfxs', 1, 2, 'https://pic2.zhimg.com/80/v2-b951b9e3e65678a93c7f2d6184b81337_720w.jpg', 'https://pic4.zhimg.com/80/v2-2501a93fe6acebfc81f8c1c825df2367_720w.jpg', '什么都无法舍弃的人，什么都无法改变。');
INSERT INTO `t_user` VALUES (2, 'heart', '8fd52bb1654248a17f1890f500bd11f0', '抖声用户wydh', 2, 0, 'https://pic2.zhimg.com/80/v2-b951b9e3e65678a93c7f2d6184b81337_720w.jpg', 'https://pic4.zhimg.com/80/v2-2501a93fe6acebfc81f8c1c825df2367_720w.jpg', '什么都无法舍弃的人，什么都无法改变。');
INSERT INTO `t_user` VALUES (3, 'jxygzzy1', '9766e44c9485ae3cf4ae6b2418864440', '抖声用户MKSS', 0, 0, 'https://pic2.zhimg.com/80/v2-b951b9e3e65678a93c7f2d6184b81337_720w.jpg', 'https://pic4.zhimg.com/80/v2-2501a93fe6acebfc81f8c1c825df2367_720w.jpg', '什么都无法舍弃的人，什么都无法改变。');
INSERT INTO `t_user` VALUES (4, 'heart1', 'f1936ffeeec86f5b1424d01b5c2606a1', '抖声用户mNoO', 1, 2, 'https://pic2.zhimg.com/80/v2-b951b9e3e65678a93c7f2d6184b81337_720w.jpg', 'https://pic4.zhimg.com/80/v2-2501a93fe6acebfc81f8c1c825df2367_720w.jpg', '什么都无法舍弃的人，什么都无法改变。');

-- ----------------------------
-- Table structure for t_video
-- ----------------------------
DROP TABLE IF EXISTS `t_video`;
CREATE TABLE `t_video`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `play_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应七牛云的key，用于获取下载链接',
  `cover_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` int(0) NOT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `comment_count` int(0) NOT NULL DEFAULT 0,
  `favorite_count` int(0) NOT NULL DEFAULT 0,
  `create_date` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `t_video_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '视频表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_video
-- ----------------------------
INSERT INTO `t_video` VALUES (9, 'b7f8e4afbf630ff93bdc4a875baf5dcd.mp4', '3847ac3da4e285a23445e9dc661d903c.jpeg', 1, '测试上传视频标题1', 0, 0, '2022-06-01 01:36:15');
INSERT INTO `t_video` VALUES (10, '0abf507bfd71d94a2c4759c8897cb31d.mp4', 'a23fcb7bd96ce31e820dbdb33f2a6f07.jpeg', 1, '测试上传视频标题2', 0, 0, '2022-06-01 01:49:24');
INSERT INTO `t_video` VALUES (11, '672bbb90a7a0fb660bdce453724e9d56.mp4', '3ff1ab2b76cf9c480b17706d5f399f8d.jpeg', 1, '#百变酒精 #猫咪的迷惑行为 #喵星人', 0, 0, '2022-06-03 08:55:42');
INSERT INTO `t_video` VALUES (12, '5e427469d036f318ac900eda52e8b59e.mp4', '749505425183b56c3c3ef4a65558b5f5.jpeg', 1, '#可爱的小猫咪 #微风吹', 0, 0, '2022-06-03 08:55:57');
INSERT INTO `t_video` VALUES (13, '262a27c5eb29f138881b1b8731572da5.mp4', '5e05e8a5e0c44971bffb3b44e5a0a45e.jpeg', 1, '#小猫咪能有什么坏心眼 明（māo）星（mī）幕后花絮大曝光', 0, 0, '2022-06-03 08:56:15');
INSERT INTO `t_video` VALUES (14, '65bcec4df55077a0b8b6b5d596764581.mp4', 'e0afadf9401b9fc770baa46665af6358.jpeg', 1, 'Mood', 0, 0, '2022-06-03 08:56:30');
INSERT INTO `t_video` VALUES (15, 'f9774ff83f3635ac62d5b6520b3e2d27.mp4', 'ac5d1517d2e550b5b439874336e2d89a.jpeg', 1, 'Supermarket Flowers', 0, 0, '2022-06-03 08:56:36');
INSERT INTO `t_video` VALUES (16, '8a8a62bbc57fd4a75bf5d12b132a68fa.mp4', '8f8711f8b98da630a5abdd0bb8c534d0.jpeg', 1, '拜托拜托，请大数据把我推给喜欢我的姨姨们#脸红研究所', 0, 0, '2022-06-03 08:56:49');
INSERT INTO `t_video` VALUES (17, '65c7032d89ba799a963d257368cf1765.mp4', '263f51dda0df30c057a312113b55099d.jpeg', 1, '给孩子送去学表演吧，戏可太多了#曼基康矮脚 #猫咪的迷惑行为', 0, 0, '2022-06-03 08:57:07');
INSERT INTO `t_video` VALUES (18, '69f4a60b2ad9060d12e0d178587c14e7.mp4', 'aa8fb3c63ef0f3e0eccfaaec316f14db.jpeg', 1, '工作一天辛苦啦，来看小猫咪喝neinei解压吧～#萌宠 #猫咪', 3, 2, '2022-06-03 08:57:19');
INSERT INTO `t_video` VALUES (19, '3c1128a789079f6fa9bd8110e3e3fde7.mp4', 'c9f521c8f1a155b7762d546d1b17bbe0.jpeg', 1, '今天的天气很好 但是小猫咪也不知道和谁出去玩', 1, 1, '2022-06-03 08:57:33');
INSERT INTO `t_video` VALUES (20, '9e6c7c7915dac1b3f2151e95dbaad799.mp4', '496e2848c76696faa94774f96f2f43f1.jpeg', 1, '看它睡觉都觉得幸福……', 0, 1, '2022-06-03 08:57:44');
INSERT INTO `t_video` VALUES (21, '32de46e2576c5bf050f5f04ec20b70bb.mp4', '0a379d7bfe6c7a0ada7821e8fa982e6b.jpeg', 1, '猫咪好治愈啊……#歌曲孤孤单单', 0, 1, '2022-06-03 08:57:55');
INSERT INTO `t_video` VALUES (22, '45b77343ce19e86af42accbf16752096.mp4', '6da2036332c92e166b3ea3ceb76e81dc.jpeg', 1, '没抓到，好尴尬啊…随便啦，我只是只小猫#猫咪的迷惑行为 #家有傻猫', 0, 1, '2022-06-03 08:58:06');
INSERT INTO `t_video` VALUES (23, '5de9b21ebcffec54895110fa4468d83a.mp4', '3726a0b4b9a055ebd2e2edfb5e478ce2.jpeg', 1, '每天一遍，快乐无限 #猫咪搞笑 #萌宠宝宝的日常', 0, 1, '2022-06-03 08:58:18');
INSERT INTO `t_video` VALUES (24, 'a66b878de49a26f5093db94e07694cb4.mp4', '4d7868cfc8adf22e993e770b69ed5e06.jpeg', 1, '它带着小背包过来了#猫 #可爱 #治愈 #小可爱哈比', 0, 1, '2022-06-03 08:58:29');
INSERT INTO `t_video` VALUES (25, 'da7520aa81284eca69614a5216219200.mp4', '9c7e309c531defa0d3a111bdc22078cb.jpeg', 1, '小猫咪帮你们抓到了，麻袋自备哦#别再冬眠去热烈的夏天', 0, 1, '2022-06-03 08:58:38');
INSERT INTO `t_video` VALUES (26, '0656be253ab57cba3a89b4161456ec34.mp4', 'b60878b76a1477d4b489eb335d6f9730.jpeg', 1, '小猫咪有多喜欢被摸头～#萌宠 #抖音动物图鉴 #猫咪的日常', 0, 1, '2022-06-03 08:59:06');
INSERT INTO `t_video` VALUES (27, '94c100280116c834538693fa475349ce.mp4', '27712dfcc7646a21678cf242de6dd3c5.jpeg', 1, '新手养猪： 请问一下你们的猪背后也有神奇按钮吗', 0, 1, '2022-06-03 08:59:14');
INSERT INTO `t_video` VALUES (28, 'cdcd339617cab47f8c81cfc68b8fdffc.mp4', 'dc758ab680c3c0c77a91bb1f323b5df1.jpeg', 1, '原来是你偷吃了我的比巴卜#猫咪的迷惑行为', 0, 0, '2022-06-03 08:59:29');
INSERT INTO `t_video` VALUES (29, 'd4e2922092678ec86c8a02943c9d1293.mp4', 'b0e9a8c36d67fbbe9ae7cc427aa699d9.jpeg', 1, '最近画画又进步了，我跟我的小猫咪都很满意今天的作品', 1, 1, '2022-06-03 08:59:41');
INSERT INTO `t_video` VALUES (30, '96f2709c4d0d90204c78f3ecc16fe9aa.mp4', 'd8af3139238884e1f6bf151e1cc56998.jpeg', 2, '快喊你身边的憨憨来看欢乐恰恰恰🌈 #二哈 #狗狗 #狗狗搞笑 #狗 - 抖音', 1, 2, '2022-06-06 14:36:42');
INSERT INTO `t_video` VALUES (31, '551e0084f7ce27f20762ea85c1fc5539.mp4', '36115a70721a283dc4d9a0deb70bd859.jpeg', 1, '狗子：咱俩到底谁是狗？ #萌宠 #狗狗 #精神小狗 #萌宠搞笑 #动物', 0, 0, '2022-06-12 10:42:19');

SET FOREIGN_KEY_CHECKS = 1;
