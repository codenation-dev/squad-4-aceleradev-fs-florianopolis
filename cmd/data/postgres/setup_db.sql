CREATE DATABASE uati;

\c uati;

-- DROP TABLE IF EXISTS public_funcs, customers, users, warnings;
DROP TABLE IF EXISTS public_func; --, customer, user, warning;

CREATE TABLE IF NOT EXISTS users (
id SERIAL,
email VARCHAR(50) UNIQUE,
password VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS public_func (
id SERIAL,
complete_name VARCHAR(100),
short_name VARCHAR(30),
wage NUMERIC(10,2),
departament VARCHAR(50),
function VARCHAR(50)
);



-- -- TODO: trocar nome das tabelas para o singular
-- CREATE TABLE IF NOT EXISTS customers	(
-- id SERIAL,
-- -- TODO: gerar hash das 30 primeiras letras do nome
-- name VARCHAR(30),
-- wage NUMERIC(10,2),
-- is_public bit,
-- sent_warning TEXT -- Tem como usar isso para guardar os id dos users que receberam os warnings?
-- );



-- CREATE TABLE IF NOT EXISTS warnings(
-- id SERIAL,
-- dt TEXT,
-- msg TEXT,
-- sent_to text,
-- from_customer TEXT
-- );

-- -- TODO: criar tabela com evolução de salários