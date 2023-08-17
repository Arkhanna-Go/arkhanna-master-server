CREATE TABLE IF NOT EXISTS `db_roulette` (
  `index` int(11) NOT NULL default '0',
  `level` smallint(5) unsigned NOT NULL,
  `item_id` int(10) unsigned NOT NULL,
  `amount` smallint(5) unsigned NOT NULL DEFAULT '1',
  `flag` smallint(5) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`index`)
) ENGINE=MyISAM;