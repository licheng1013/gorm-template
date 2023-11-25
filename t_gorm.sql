/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80029 (8.0.29)
 Source Host           : 192.168.101.90:3306
 Source Schema         : t_gorm

 Target Server Type    : MySQL
 Target Server Version : 80029 (8.0.29)
 File Encoding         : 65001

 Date: 25/11/2023 13:53:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_admin
-- ----------------------------
DROP TABLE IF EXISTS `t_admin`;
CREATE TABLE `t_admin`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '管理员id',
  `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `salt` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '盐',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `nick_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12317 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_admin
-- ----------------------------
INSERT INTO `t_admin` VALUES (3, 'vben', '1c94ae27aa606eb4717ec1e44686a521', '7702151f-f8e6-47d7-82c7-7a8ec21c16b1', '2022-10-16 17:03:46', 'Gorm1112323');
INSERT INTO `t_admin` VALUES (12313, 'adminaa', '1c94ae27aa606eb4717ec1e44686a521', '7702151f-f8e6-47d7-82c7-7a8ec21c16b1', '2023-08-01 15:18:08', 'Gorm');
INSERT INTO `t_admin` VALUES (12314, 'aadsfasd', 'fasdfsadf', 'asdfasdfasdf', '2023-11-10 11:15:13', 'asdfasdfafd');
INSERT INTO `t_admin` VALUES (12316, '111', '12312', '123123', '2023-11-24 17:18:18', '1121323123');

-- ----------------------------
-- Table structure for t_user_info
-- ----------------------------
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user_info
-- ----------------------------
INSERT INTO `t_user_info` VALUES (6, 'asdf', 'asbaba', '2023-11-10 17:41:42');
INSERT INTO `t_user_info` VALUES (7, 'bbb', 'aaa', '2023-11-10 17:47:49');
INSERT INTO `t_user_info` VALUES (8, 'bbb', 'aaa', '2023-11-10 17:47:52');
INSERT INTO `t_user_info` VALUES (9, 'aaa', 'adfasdf', '2023-11-10 17:47:56');
INSERT INTO `t_user_info` VALUES (10, 'sadadfv', 'afsdfadsf', '2023-11-10 17:48:00');
INSERT INTO `t_user_info` VALUES (11, 'asdfasdf', 'asdfsdfasd', '2023-11-10 17:48:03');
INSERT INTO `t_user_info` VALUES (12, 'asdfasdf', 'asdfsdf', '2023-11-10 17:48:06');
INSERT INTO `t_user_info` VALUES (13, 'adfasdf', 'asdfsadfa', '2023-11-10 17:48:10');
INSERT INTO `t_user_info` VALUES (18, '112', '123123', '2023-11-25 09:45:15');

SET FOREIGN_KEY_CHECKS = 1;
