CREATE TABLE IF NOT EXISTS User (
	gitlab_id INT PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	slack_id VARCHAR(9) NOT NULL,
	name VARCHAR(255) NOT NULL,
	default_channel VARCHAR(32) NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS Project (
	id INT PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	default_channel VARCHAR(32) NULL DEFAULT ''
);
