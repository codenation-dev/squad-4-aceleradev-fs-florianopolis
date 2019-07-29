CREATE DATABASE uati;

\c uati;

DROP TABLE IF EXISTS public_funcs, customers, users, warnings;
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
function VARCHAR(50),
relevancia smallint
);

CREATE TABLE IF NOT EXISTS customer	(
id SERIAL,
name VARCHAR(30)
);


-- alter table public_func add column relevancia smallint;

-- update public_func set relevancia = floor(random() * 10) + 1;
-- commit;

-- CREATE TABLE IF NOT EXISTS warnings(
-- id SERIAL,
-- dt TEXT,
-- msg TEXT,
-- sent_to text,
-- from_customer TEXT
-- );

-- -- TODO: criar tabela com evolução de salários