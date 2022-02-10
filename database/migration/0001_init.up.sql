CREATE TYPE role as enum ('user','admin', 'sub-admin');
CREATE TABLE users
(
    ID        TEXT,
    Email     VARCHAR(50) NOT NULL,
    FirstName VARCHAR(20) NOT NULL,
    LastName  VARCHAR(20) NOT NULL,
    UserID    VARCHAR(25) NOT NULL,
    Password  TEXT        NOT NULL,
    MobileNo  VARCHAR(10) NOT NULL,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE (Email, ID, UserId, MobileNo)
);

CREATE TABLE IF NOT EXISTS session(
                                      id TEXT NOT NULL ,
                                      userid Text NOT NULL,
                                      createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS restaurant(
                                         name TEXT NOT NULL,
                                         lat FLOAT,
                                         lng FLOAT,
                                         restaurantid int primary key ,
                                         createdby text
);

CREATE TABLE IF NOT EXISTS dishes(
                                     restaurantid int,
                                     dishName TEXT,
                                     price INT,
                                     CONSTRAINT fk_id
                                         FOREIGN KEY(restaurantid)
                                             REFERENCES restaurant(restaurantid)
);