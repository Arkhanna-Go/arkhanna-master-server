CREATE TABLE IF NOT EXISTS `market` (
  `name` varchar(50) NOT NULL DEFAULT '',
  `nameid` int(10) UNSIGNED NOT NULL,
  `price` INT(11) UNSIGNED NOT NULL,
  `amount` INT(11) NOT NULL,
  `flag` TINYINT(2) UNSIGNED NOT NULL DEFAULT '0',
  PRIMARY KEY  (`name`,`nameid`)
) ENGINE = MyISAM;