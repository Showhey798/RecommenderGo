CREATE TABLE movies (
	id SERIAL NOT NULL,
	title VARCHAR ( 50 ) NOT NULL,
	created_on TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE users (
	id VARCHAR(50),
	username VARCHAR ( 50 ) NOT NULL,
	pwd VARCHAR(50),
	created_on TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE clicks (
	id SERIAL,
	movie_id INT NOT NULL,
	user_id VARCHAR(50) NOT NULL,
	clicked_on TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE ratings (
	id SERIAL,
	item_id INT NOT NULL,
	user_id VARCHAR(50) NOT NULL,
	rating FLOAT NOT NULL,
	clicked_on TIMESTAMP NOT NULL,
	PRIMARY KEY (id)
);