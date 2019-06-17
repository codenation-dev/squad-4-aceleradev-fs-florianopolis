CREATE DATABASE uati;

\c uati;

DROP TABLE IF EXISTS public_funcs, customers, users, warnings;

CREATE TABLE IF NOT EXISTS customers	(
id SERIAL,
name TEXT,
wage NUMERIC(10,2),
is_public bit,
sent_warning TEXT -- Tem como usar isso para guardar os id dos users que receberam os warnings?
);

-- CREATE TABLE IF NOT EXISTS public_funcs (
-- fid SERIAL,
-- name TEXT,
-- wage NUMERIC(10,2),
-- place TEXT 
-- );

-- CREATE TABLE IF NOT EXISTS users(
-- uid SERIAL,
-- login TEXT UNIQUE,
-- email TEXT UNIQUE,
-- pass TEXT
-- );

-- CREATE TABLE IF NOT EXISTS warnings(
-- wid SERIAL,
-- dt TEXT,
-- msg TEXT,
-- sent_to text,
-- from_customer TEXT
-- );
