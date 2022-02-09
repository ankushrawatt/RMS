DROP TABLE users;
CREATE TYPE role as enum ('user','admin', 'sub-admin');
CREATE TABLE users
(
    ID        UUID,
    Email     VARCHAR(50) NOT NULL,
    Role    role,
    FirstName VARCHAR(20) NOT NULL,
    LastName  VARCHAR(20) NOT NULL,
    UserID    VARCHAR(25) NOT NULL,
    Password  TEXT        NOT NULL,
    MobileNo  VARCHAR(10) NOT NULL,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE (Email, ID, UserId, MobileNo)
);