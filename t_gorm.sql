/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 192.168.101.90:3306
 Source Schema         : t_gorm

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 29/06/2023 10:18:09
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
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_admin
-- ----------------------------
INSERT INTO `t_admin` VALUES (3, 'admin', '1c94ae27aa606eb4717ec1e44686a521', '7702151f-f8e6-47d7-82c7-7a8ec21c16b1', '2022-10-16 17:03:46', 'Gorm');

SET FOREIGN_KEY_CHECKS = 1;
