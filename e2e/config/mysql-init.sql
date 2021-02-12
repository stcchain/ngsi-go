DROP DATABASE IF EXISTS idm;
DROP USER IF EXISTS 'keyrock'@'%';
CREATE USER 'keyrock'@'%' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON idm.* TO 'keyrock'@'%';
FLUSH PRIVILEGES;
