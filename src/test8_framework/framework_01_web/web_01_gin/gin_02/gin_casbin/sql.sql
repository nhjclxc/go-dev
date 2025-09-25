

CREATE TABLE `user`
(
    `id`         BIGINT AUTO_INCREMENT PRIMARY KEY,
    `username`   VARCHAR(50)  NOT NULL UNIQUE,
    `password`   VARCHAR(255) NOT NULL,
    `email`      VARCHAR(100),
    `status`     TINYINT   DEFAULT 1 COMMENT '1=active,0=disabled',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO user (id, username, password) VALUES
                                              (1, 'superadminUser', '$2a$10$K5b1SMcqqObPWyT5dCVKauepxwsAHLJCruRam3oNfg5SXArFeCojq'),
                                              (2, 'adminUser', '$2a$10$HMv7fwShhxKUCPyvDUb.ou8xNjIIQSeOqhHpscu5It.n4R46jjHQK'),
                                              (3, 'commonUser', '$2a$10$H5iz6nT6BHBDuBYfBMrR0.lWRbhTpyYBqFqyS9B4W/TQttH9Ghmca');

CREATE TABLE `role`
(
    `id`          BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name`        VARCHAR(50) NOT NULL UNIQUE,
    `description` VARCHAR(255),
    `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO role (id, name) VALUES (1, 'superadmin'), (2, 'admin'), (3, 'common');


CREATE TABLE `user_role`
(
    `id`      BIGINT AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL,
    `role_id` BIGINT NOT NULL,
    UNIQUE KEY `uk_user_role` (`user_id`, `role_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
INSERT INTO user_role (id, user_id, role_id) VALUES (1, 1, 1), (2, 2, 2), (3, 3, 3);


CREATE TABLE `menu`
(
    `id`         BIGINT AUTO_INCREMENT PRIMARY KEY,
    `parent_id`  BIGINT    DEFAULT 0,
    `name`       VARCHAR(50)  NOT NULL,
    `path`       VARCHAR(255) NOT NULL COMMENT 'URL或权限标识',
    `type`       TINYINT   DEFAULT 1 COMMENT '1=菜单,2=按钮/API',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO menu (id, parent_id, name, path, type)
VALUES (1, 0, '首页', '/home', 1),(2, 0, '个人中心', '/profile', 1),
       (3, 0, '用户管理', '/user', 1),(4, 3, '新增用户', '/api/user/insert', 2),(5, 3, '修改用户', '/api/user/update', 2),
       (6, 0, '角色管理', '/role', 1),(7, 6, '查询角色列表', '/api/role/get', 2);


CREATE TABLE `role_menu`
(
    `id`      BIGINT AUTO_INCREMENT PRIMARY KEY,
    `role_id` BIGINT NOT NULL,
    `menu_id` BIGINT NOT NULL,
    UNIQUE KEY `uk_role_menu` (`role_id`, `menu_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

INSERT INTO role_menu (role_id, menu_id) VALUES
-- superadmin 拥有所有菜单
(1, 1),(1, 2),(1, 3),(1, 4),(1, 5),(1, 6),(1, 7),
-- admin 拥有部分
(2, 1), (2, 2),(2, 3),(2, 4),(2, 5),
-- common 仅首页和个人中心
(3, 1),(3, 2);

CREATE TABLE `casbin_rule`
(
    `id`    BIGINT AUTO_INCREMENT PRIMARY KEY,
    `ptype` VARCHAR(100) NOT NULL COMMENT '策略类型，如 p（policy）、g（角色继承）',
    `v0`    VARCHAR(100) COMMENT '第一个值（通常是 subject 用户/角色）',
    `v1`    VARCHAR(100) COMMENT '第二个值（对象/资源，例如 URL 或资源名',
    `v2`    VARCHAR(100) COMMENT '第三个值（动作，例如 GET/POST）',
    `v3`    VARCHAR(100) COMMENT '第四个值（可选）',
    `v4`    VARCHAR(100) COMMENT '第五个值（可选）',
    `v5`    VARCHAR(100) COMMENT '第六个值（可选）'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- 正确写法：用 g 来绑定用户到角色，用 p 来定义角色到资源的权限;;

-- 用户绑定到角色（用 g）
-- 用户 superadminUser 是 superadmin 角色;
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('g', 'superadminUser', 'superadmin', ''),('g', 'adminUser', 'admin', ''),('g', 'commonUser', 'common', '');

-- 角色到资源的策略（用 p）
-- common 角色可以访问 /home和/profile
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('p', 'common', '/api/home', '*'),('p', 'common', '/api/profile', '*');

-- admin 角色可以 user 相关的
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('p', 'admin', '/api/user', '*'),('p', 'admin', '/api/user/insert', 'POST'),('p', 'admin', '/api/user/update', 'PUT');

-- superadmin 角色可以 role 相关的
INSERT INTO casbin_rule (ptype, v0, v1, v2)
VALUES ('p', 'superadmin', '/api/role', '*'),('p', 'superadmin', '/api/role/get', 'GET');

-- superadmin 继承 admin， 所以 superadmin 既能做 admin 的事情，还能有自己额外的权限。
INSERT INTO casbin_rule (ptype, v0, v1)
VALUES ('g', 'admin', 'common'), ('g', 'superadmin', 'admin');
