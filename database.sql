CREATE TABLE `stocks`(
  `code` VARCHAR(16)  NOT NULL COMMENT'代码',
  `name` VARCHAR(16) NOT NULL COMMENT'名称',
  `industry` VARCHAR(16) NOT NULL COMMENT'行业',
  `area` VARCHAR(16) NOT NULL COMMENT'地区',
  `pe` FLOAT NOT NULL COMMENT'市盈率',
  `outstanding` DOUBLE NOT NULL COMMENT'流通股本',
  `totals` DOUBLE NOT NULL COMMENT'总股本',
  `total_assets` DOUBLE NOT NULL COMMENT'总资产',
  `liquid_assets` DOUBLE NOT NULL COMMENT'流动资产',
  `fixed_assets` DOUBLE NOT NULL COMMENT'固定资产',
  `reserved` DOUBLE NOT NULL COMMENT'公积金',
  `reserved_per_share` FLOAT NOT NULL COMMENT'每股公积金',
  `eps` FLOAT NOT NULL COMMENT'每股收益',
  `bvps` FLOAT NOT NULL COMMENT'每股净资',
  `pb` FLOAT NOT NULL COMMENT'市净率',
  `time_to_market` DATE NOT NULL COMMENT'上市日期'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;