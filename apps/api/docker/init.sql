CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE DATABASE pictio;
\c pictio;

CREATE TABLE "rooms" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    time TIMESTAMP NOT NULL
);

CREATE TABLE "users" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
);

CREATE TABLE "rooms_users" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    rooms_id UUID,
    users_id UUID,
    time TIMESTAMP NOT NULL,
    
    fk_rooms_rooms_users FOREIGN KEY (rooms_id) REFERENCES rooms(id) ON DELETE CASCADE,
    fk_users_rooms_rooms FOREIGN KEY (users_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE "chat" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content JSONB NOT NULL,
    users_id UUID,
    time TIMESTAMP NOT NULL,

    fk_users_chat FOREIGN KEY (users_id) REFERENCES users(id) ON DELETE CASCADE
)