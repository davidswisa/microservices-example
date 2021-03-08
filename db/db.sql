CREATE TABLE IF NOT EXISTS reservations (
    id int PRIMARY KEY,
	Name VARCHAR ( 50 ) NOT NULL,
	Date VARCHAR ( 50 ) NOT NULL,
    Party int NOT NULL,
	Hour int NOT NULL
);


-- select * from reservations

-- drop table reservations
