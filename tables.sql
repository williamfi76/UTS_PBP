CREATE TABLE accounts (
	id int PRIMARY KEY AUTO_INCREMENT,
    username varchar(255)
);

CREATE TABLE games (
	id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255),
    max_player int
);

CREATE TABLE rooms (
	id int PRIMARY KEY AUTO_INCREMENT,
    room_name varchar(255),
    id_game int,
    FOREIGN KEY (id_game) REFERENCES games(id)
);

CREATE TABLE participants (
	id int PRIMARY KEY AUTO_INCREMENT,
    id_room int,
    id_account int,
    FOREIGN KEY (id_room) REFERENCES rooms(id),
    FOREIGN KEY (id_account) REFERENCES accounts(id)
);

INSERT INTO Accounts (ID, Username) VALUES
(1, 'Andi'),
(2, 'Budi'),
(3, 'Citra'),
(4, 'Dewi'),
(5, 'Eko'),
(6, 'Fitri'),
(7, 'Gunawan'),
(8, 'Hani'),
(9, 'Indra'),
(10, 'Joko');

INSERT INTO Games (ID, Name, Max_player) VALUES
(1, 'League of Legends', 10),
(2, 'Counter-Strike: Global Offensive', 10),
(3, 'Dota 2', 10),
(4, 'Fortnite', 100),
(5, 'Apex Legends', 60),
(6, 'Overwatch', 12),
(7, 'Rainbow Six Siege', 10),
(8, 'Call of Duty: Warzone', 150),
(9, 'World of Warcraft', 40),
(10, 'Rocket League', 8);

INSERT INTO Rooms (ID, Room_name, Id_game) VALUES
(1, 'Room 101', 1),
(2, 'Room 102', 1),
(3, 'Room 201', 2),
(4, 'Room 202', 2),
(5, 'Room 301', 3),
(6, 'Room 302', 3),
(7, 'Room 401', 4),
(8, 'Room 402', 4),
(9, 'Room 501', 5),
(10, 'Room 502', 5);

INSERT INTO Participants (ID, Id_room, Id_account) VALUES
(1, 3, 7),
(2, 1, 9),
(3, 2, 5),
(4, 5, 4),
(5, 4, 8);