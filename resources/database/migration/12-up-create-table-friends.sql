CREATE TABLE IF NOT EXISTS `friends` (
  `char_id` int(11) NOT NULL default '0',
  `friend_id` int(11) NOT NULL default '0',
  PRIMARY KEY (`char_id`, `friend_id`)
) ENGINE=MyISAM;