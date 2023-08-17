CREATE TABLE IF NOT EXISTS `party` (
  `party_id` int(11) unsigned NOT NULL auto_increment,
  `name` varchar(24) NOT NULL default '',
  `exp` tinyint(11) unsigned NOT NULL default '0',
  `item` tinyint(11) unsigned NOT NULL default '0',
  `leader_id` int(11) unsigned NOT NULL default '0',
  `leader_char` int(11) unsigned NOT NULL default '0',
  PRIMARY KEY  (`party_id`)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `party_bookings` (
  `world_name` varchar(32) NOT NULL,
  `account_id` int(11) unsigned NOT NULL,
  `char_id` int(11) unsigned NOT NULL,
  `char_name` varchar(23) NOT NULL,
  `purpose` smallint(5) unsigned NOT NULL DEFAULT '0',
  `assist` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `damagedealer` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `healer` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `tanker` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `minimum_level` smallint(5) unsigned NOT NULL,
  `maximum_level` smallint(5) unsigned NOT NULL,
  `comment` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`world_name`, `account_id`, `char_id`)
) ENGINE=MyISAM;