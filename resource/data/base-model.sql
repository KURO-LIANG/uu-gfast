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
