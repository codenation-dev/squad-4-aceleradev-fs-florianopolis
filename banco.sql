CREATE TABLESPACE uati
  OWNER postgres
  LOCATION '/var/lib/postgresql/data';

ALTER TABLESPACE uati
  OWNER TO postgres;

CREATE DATABASE uati
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = uati
    CONNECTION LIMIT = -1;


create table usuario (
	id_usuario int not null,	
	nome varchar(255) not null,
	email varchar(11) not null,
	senha varchar(50) not null,
primary key (id_usuario))


CREATE TABLE cliente(
    id_cliente integer NOT NULL,
    nome varchar(255) NOT NULL,
	PRIMARY KEY (id_cliente)
)

CREATE TABLE public.funcionario
(
    id_funcionario integer NOT NULL,
    nome varchar(255) NOT NULL,
    salario numeric(15,2),
    PRIMARY KEY (id_funcionario)    
)
