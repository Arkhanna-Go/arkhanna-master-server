CREATE TABLE IF NOT EXISTS `achievement` (
  `char_id` int(11) unsigned NOT NULL default '0',
  `id` bigint(11) unsigned NOT NULL,
  `count1` int unsigned NOT NULL default '0',
  `count2` int unsigned NOT NULL default '0',
  `count3` int unsigned NOT NULL default '0',
  `count4` int unsigned NOT NULL default '0',
  `count5` int unsigned NOT NULL default '0',
  `count6` int unsigned NOT NULL default '0',
  `count7` int unsigned NOT NULL default '0',
  `count8` int unsigned NOT NULL default '0',
  `count9` int unsigned NOT NULL default '0',
  `count10` int unsigned NOT NULL default '0',
  `completed` datetime,
  `rewarded` datetime,
  PRIMARY KEY (`char_id`,`id`),
  KEY `char_id` (`char_id`)
) ENGINE=MyISAM;