-- 创建一个cms_account库，用于存储用户信息
CREATE DATABASE IF NOT EXISTS cms_account DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE cms_account;

CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
    user_id VARCHAR(50) DEFAULT '' COMMENT '用户ID，自增主键',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名，唯一且不能为空',
    email VARCHAR(100) NOT NULL UNIQUE COMMENT '邮箱，唯一且不能为空',
    password_hash VARCHAR(255) NOT NULL COMMENT '密码的哈希值，不能为空',
    full_name VARCHAR(100) COMMENT '真实姓名',
    date_of_birth DATE COMMENT '出生日期',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，默认为当前时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，每次更新均为当前时间'
)   COMMENT = '用户信息表，用于存储用户账户的基本信息';
