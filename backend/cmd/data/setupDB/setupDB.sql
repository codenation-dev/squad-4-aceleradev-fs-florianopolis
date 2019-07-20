CREATE DATABASE uati;

\c uati;

DROP TABLE IF EXISTS public_funcs, customers, users, warnings, usuario, importacao, cliente, funcionario_salario, funcionario;

-- TODO: trocar nome das tabelas para o singular
CREATE TABLE IF NOT EXISTS customers	(
id SERIAL,
-- TODO: gerar hash das 30 primeiras letras do nome
name varchar(30),
wage NUMERIC(10,2),
is_public bit,
sent_warning TEXT -- Tem como usar isso para guardar os id dos users que receberam os warnings?
);

CREATE TABLE IF NOT EXISTS public_funcs (
id SERIAL,
name TEXT,
wage NUMERIC(10,2),
place TEXT 
);

CREATE TABLE IF NOT EXISTS users(
id SERIAL,
-- login TEXT,
email TEXT UNIQUE,
pass TEXT
);

CREATE TABLE IF NOT EXISTS warnings(
id SERIAL,
dt TEXT,
msg TEXT,
sent_to text,
from_customer TEXT
);

create table if not exists funcionario(
    id_funcionario SERIAL primary key,
    mes_referencia date not null,
    nome varchar(255) not null,
    nome_pesquisa varchar(32) not null,
    cargo varchar(100),
    orgao varchar(100),
    estado varchar(50),
    salario_mensal numeric(15,2),
    salario_ferias numeric(15,2),
    pagto_eventual numeric(15,2),
    licenca_premio numeric(15,2),
    abono_salario numeric(15,2),
    redutor_salarial numeric(15,2),
    total_liquido numeric(15,2)
);

CREATE TABLE IF NOT EXISTS usuario(
    id_usuario serial primary key,
    email varchar(255) unique not null,
    password varchar(60) not null
);

CREATE TABLE IF NOT EXISTS cliente(
    id_cliente SERIAL not null,
    nome varchar(30) not null,
    nome_pesquisa varchar(30) NOT NULL UNIQUE,
    salario NUMERIC(15,2)
);

alter table cliente add constraint pk_cliente primary key (id_cliente);

-- TODO: criar tabela com evolução de salários