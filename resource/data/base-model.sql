CREATE TABLE `base_region`
(
    `id`                int(10) unsigned NOT NULL,
    `parent_id`         int(10) unsigned                 DEFAULT 0 COMMENT '父ID',
    `region_name`       varchar(45) COLLATE utf8mb4_bin  DEFAULT '' COMMENT '名称',
    `merger_name`       varchar(200) COLLATE utf8mb4_bin DEFAULT '' COMMENT '全称',
    `short_name`        varchar(45) COLLATE utf8mb4_bin  DEFAULT '' COMMENT '简称',
    `merger_short_name` varchar(200) COLLATE utf8mb4_bin DEFAULT '' COMMENT '简称合并',
    `level`             int(11)                          DEFAULT 0 COMMENT '层级，1是省份，2是城市，3是区县',
    `city_code`         varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '城市代码',
    `zip_code`          varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '邮编号码',
    `full_pinyin`       varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '全拼',
    `simplified_pinyin` varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '简拼',
    `first_char`        varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '第一个字',
    `longitude`         varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '纬度',
    `latitude`          varchar(45) COLLATE utf8mb4_bin  DEFAULT NULL COMMENT '经度',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = DYNAMIC COMMENT ='行政区域省市区县';

CREATE TABLE base_user
(
    user_id               BIGINT(11)   NOT NULL COMMENT '用户ID',
    uid                   VARCHAR(255) NOT NULL COMMENT '用户UUID',
    ma_open_id            VARCHAR(255) COMMENT '小程序openid',
    union_id              VARCHAR(255) COMMENT '微信开放平台id',
    nick_name             VARCHAR(255) COMMENT '用户昵称',
    avatar                VARCHAR(255) COMMENT '用户头像',
    gender                INT COMMENT '性别 0-未知，1-男，2-女',
    phone                 VARCHAR(20) COMMENT '手机号',
    email                 VARCHAR(255) COMMENT '邮箱',
    rel_name              VARCHAR(255) COMMENT '真实姓名',
    credential_type       INT COMMENT '证件类型,1-身份证，2-国内外护照，3-港澳通行证，4-香港身份证，5-澳门身份证，6-台湾身份证',
    credential_code       VARCHAR(255) COMMENT '证件号码',
    user_img_url          VARCHAR(255) COMMENT '用户照片',
    id_card_front_img_url VARCHAR(255) COMMENT '证件正面照片',
    id_card_back_img_url  VARCHAR(255) COMMENT '证件背面照片',
    verify_time           DATETIME COMMENT '实名认证时间',
    verify_state          INT COMMENT '实名认证状态，0-未认证，1-认证中，2-认证通过，3-认证不通过',
    audit_time            DATETIME COMMENT '审核时间',
    audit_user            VARCHAR(255) COMMENT '审核人',
    audit_remark          VARCHAR(255) COMMENT '审核意见',
    last_login_time       DATETIME COMMENT '上次登录时间',
    last_login_ip         VARCHAR(255) COMMENT '上次登录IP',
    last_login_info       VARCHAR(255) COMMENT '上次登录设备信息',
    subscribe_num         INT COMMENT '关注次数',
    subscribe_scene       VARCHAR(255) COMMENT '返回用户关注的渠道来源',
    subscribe_time        DATETIME COMMENT '关注时间',
    cancel_subscribe_time DATETIME COMMENT '取消关注时间',
    qr_scene_str          VARCHAR(255) COMMENT '二维码扫码场景',
    language              VARCHAR(50) COMMENT '语言',
    subscribe             INT COMMENT '是否关注公众号 0-否；1-是；',
    long_and_lati         VARCHAR(50) COMMENT '经纬度  经度,纬度',
    created_at            DATETIME COMMENT '创建时间',
    updated_at            DATETIME COMMENT '修改时间',
    updated_by            VARCHAR(255) COMMENT '修改人',
    deleted_at            DATETIME COMMENT '删除时间',
    PRIMARY KEY (user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

