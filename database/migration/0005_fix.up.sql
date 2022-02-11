DROP TABLE dishes;
DROP TABLE restaurant;
DROP TABLE session;
DROP TABLE users;


CREATE TABLE users
(
    id        TEXT NOT NULL PRIMARY KEY,
    email     VARCHAR(50) NOT NULL,
    firstname VARCHAR(20) NOT NULL,
    lastname  VARCHAR(20) NOT NULL,
    userid    VARCHAR(25) NOT NULL ,
    password  TEXT        NOT NULL,
    mobileno  VARCHAR(10) NOT NULL,
    createdat TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    role role,
    createdby text,
    UNIQUE (email, id, mobileno)
);

CREATE TABLE IF NOT EXISTS session(
                                      id TEXT NOT NULL ,
                                      userid Text NOT NULL,
                                      createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                                      CONSTRAINT fk_id
                                          FOREIGN KEY(id)
                                              REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS restaurant(
                                         name TEXT NOT NULL,
                                         lat FLOAT,
                                         lng FLOAT,
                                         restaurantid int primary key ,
                                         createdby text,
                                         CONSTRAINT fk_id
                                             FOREIGN KEY(createdby)
                                                 REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS dishes(
                                     restaurantid int,
                                     id SERIAL,
                                     dishName TEXT,
                                     price INT,
                                     CONSTRAINT fk_id
                                         FOREIGN KEY(restaurantid)
                                             REFERENCES restaurant(restaurantid)
);