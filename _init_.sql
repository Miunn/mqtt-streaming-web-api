CREATE DATABASE CCSVideos;
CREATE USER 'ccs-videos-connector'@'%' IDENTIFIED BY 'connectorPassword123*';
GRANT ALL PRIVILEGES ON CCSVideos.* TO 'ccs-videos-connector'@'%';