create table sys_dict_details
(
    id int unsigned auto_increment
        primary key,
    dict_id int unsigned not null,
    label varchar(45) not null,
    value int default 0 not null,
    sort int unsigned default 0 not null,
    created datetime default CURRENT_TIMESTAMP null,
    updated datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table sys_dicts
(
    id int unsigned auto_increment
        primary key,
    name varchar(45) not null,
    dict_type varchar(45) not null,
    status tinyint(1) not null,
    more varchar(100) default '' not null comment '字典描述',
    created datetime default CURRENT_TIMESTAMP null,
    updated datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table sys_menus
(
    id int unsigned auto_increment
        primary key,
    parent_id int unsigned default 0 not null,
    path varchar(45) not null,
    name varchar(45) not null,
    is_hidden tinyint(1) default 0 not null,
    component varchar(45) not null,
    sort smallint(4) default 0 null,
    title varchar(45) not null,
    icon varchar(45) not null,
    created datetime default CURRENT_TIMESTAMP null,
    updated datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table sys_roles
(
    id int auto_increment
        primary key,
    p_id int unsigned default 0 not null,
    role_name varchar(45) null,
    created datetime default CURRENT_TIMESTAMP null,
    updated datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

create table sys_users
(
    id int unsigned auto_increment
        primary key,
    username varchar(45) not null,
    password varchar(45) not null,
    role_id int unsigned default 0 not null,
    created timestamp default CURRENT_TIMESTAMP not null,
    updated timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint username_UNIQUE
        unique (username)
);

create index idx_db_users_created
	on sys_users (created);

create index idx_db_users_role_id
	on sys_users (role_id);

create index idx_db_users_updated
	on sys_users (updated);

