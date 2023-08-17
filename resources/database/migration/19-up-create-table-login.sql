CREATE TABLE IF NOT EXISTS `ipbanlist` (
  `list` varchar(15) NOT NULL default '',
  `btime` datetime NOT NULL,
  `rtime` datetime NOT NULL,
  `reason` varchar(255) NOT NULL default '',
  PRIMARY KEY (`list`, `btime`)
) ENGINE=MyISAM;

CREATE TABLE IF NOT EXISTS `login` (
  `account_id` int(11) unsigned NOT NULL auto_increment,
  `userid` varchar(23) NOT NULL default '',
  `user_pass` varchar(32) NOT NULL default '',
  `sex` enum('M','F','S') NOT NULL default 'M',
  `email` varchar(39) NOT NULL default '',
  `group_id` tinyint(3) NOT NULL default '0',
  `state` int(11) unsigned NOT NULL default '0',
  `unban_time` int(11) unsigned NOT NULL default '0',
  `expiration_time` int(11) unsigned NOT NULL default '0',
  `logincount` mediumint(9) unsigned NOT NULL default '0',
  `lastlogin` datetime,
  `last_ip` varchar(100) NOT NULL default '',
  `birthdate` DATE,
  `character_slots` tinyint(3) unsigned NOT NULL default '0',
  `pincode` varchar(4) NOT NULL DEFAULT '',
  `pincode_change` int(11) unsigned NOT NULL DEFAULT '0',
  `vip_time` int(11) unsigned NOT NULL default '0',
  `old_group` tinyint(3) NOT NULL default '0',
  `web_auth_token` varchar(17) null,
  `web_auth_token_enabled` tinyint(2) NOT NULL default '0',
  PRIMARY KEY  (`account_id`),
  KEY `name` (`userid`),
  UNIQUE KEY `web_auth_token_key` (`web_auth_token`)
) ENGINE=MyISAM AUTO_INCREMENT=2000000; 

-- added standard accounts for servers, VERY INSECURE!!!
-- inserted into the table called login which is above
-- TODO: make a endpoint to add this dinamically
INSERT INTO `login` (`account_id`, `userid`, `user_pass`, `sex`, `email`) VALUES ('1', 's1', 'p1', 'S','athena@athena.com');
