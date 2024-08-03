CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE DATABASE ecomm;
\c ecomm;

CREATE TABLE "categories" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50)
);

CREATE TABLE "products" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    categories_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_categories_products FOREIGN KEY (categories_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE "basket" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content JSONB
);

CREATE TABLE "users_infos" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    adresse VARCHAR(255),
    country VARCHAR(150)
);

CREATE TABLE "users" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password VARCHAR(255) NOT NULL,
    users_infos UUID,
    basket_id UUID,

    CONSTRAINT fk_basket_users FOREIGN KEY (basket_id) REFERENCES basket(id),
    CONSTRAINT fk_users_infos_users FOREIGN KEY (users_infos) REFERENCES users_infos(id) ON DELETE CASCADE
);