CREATE TABLE address(
    id TEXT NOT NULL ,
    lat float4,
    lng float4,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        CONSTRAINT fk_id
        FOREIGN KEY(id)
        REFERENCES users(id)
            ON DELETE CASCADE
);