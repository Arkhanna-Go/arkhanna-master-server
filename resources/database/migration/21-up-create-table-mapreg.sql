CREATE TABLE IF NOT EXISTS `mapreg` (
  `varname` varchar(32) binary NOT NULL,
  `index` int(11) unsigned NOT NULL default '0',
  `value` varchar(255) NOT NULL,
  PRIMARY KEY (`varname`,`index`)
) ENGINE=MyISAM;