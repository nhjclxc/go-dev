create table tab_user
(
    user_id       bigint unsigned auto_increment comment 'Standard field for the primary key',
    name          varchar(255)                       not null comment 'A regular string field',
    email         varchar(255)                       null comment 'A pointer to a string, allowing for null values',
    age           tinyint unsigned                   not null comment 'An unsigned 8-bit integer',
    birthday      datetime                           null comment 'A pointer to time.Time, can be null',
    member_number varchar(255)                       null comment 'Uses sql.NullString to handle nullable strings',
    remark        varchar(128)                       null comment '备注',
    activated_at  datetime                           null comment 'Uses sql.NullTime for nullable time fields',
    created_at    datetime default CURRENT_TIMESTAMP not null comment 'Automatically managed by GORM for creation time',
    updated_at    datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment 'Automatically managed by GORM for update time',
    primary key (user_id)
)
    comment 'tab_user' charset = utf8mb4;

