CREATE TABLE IF NOT EXISTS `interlog` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `time` datetime NOT NULL,
  `log` varchar(255) NOT NULL default '',
  PRIMARY KEY (`id`),
  INDEX `time` (`time`)
) ENGINE=MyISAM;