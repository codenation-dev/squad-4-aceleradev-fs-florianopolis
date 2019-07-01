# Preparando ambiente para desenvolvimento


## Clonar este repositório

Este repositório deve ser clonado DIRETAMENTE no gopath, dentro da pasta /src/github.com/codenation-dev, assim fica com os nomes dos imports corretos, iguais ao da url do github.


## Caso já esteja com o docker do postgres instalado

Pode usar diretamente os comandos abaixo:

Descobrir o número do container:
`docker ps -a`

Iniciar e terminar de usar o container já criado:
```
docker start <seu_containder_id>
docker stop <seu_containder_id>

docker exec -ti containder_id psql -U postgres
```


## Instalação do postgres

Caso ainda não o tenha em sua máquina.


### Subindo com Docker

```
docker pull postgres
docker volume create pgdata
docker run --name postgres -e POSTGRES_PASSWORD=12345 -v
pgdata:/var/lib/postgresql/data -d postgres
```

### POSTGRESQL CLIENT

Para gerenciar nosso postgres podemos usar o psql:

Vamos descobrir o IP de nosso server postgres

```
docker inspect postgres | grep IPAddress
// Output:
// "SecondaryIPAddresses": null,
// "IPAddress": "172.17.0.2",

docker run -it --rm postgres psql -h 172.17.0.2 -U postgres
// Output:
// postgres=#
```


### Setup do Banco de dados

Dentro da linha do psql `postgres=#` colar o conteúdo do arquivo 'squad-4-aceleradev-fs-florianopolis/backend/cmd/data/setupDB/setupDB.sql'.


### Iniciar a aplicação

Dentro da pasta 'squad-4-aceleradev-fs-florianopolis', digitar:

```
go build main.go
./main
```


## Documentação

https://documenter.getpostman.com/view/7983176/S1a7UQKp?version=latest












# Gestão de clientes Banco Uati

## Objetivo

O objetivo deste produto é monitorar e gerar alertas da captura de uma determinada fonte com base em uma determinada base do cliente e regra pré estabelecida.


## Contextualização

O Banco Uati gostaria de monitorar de forma contínua e automatizada caso um de seus clientes vire um funcionário público do estado de SP (http://www.transparencia.sp.gov.br/busca-agentes.html) ou seja um bom cliente com um salário maior que 20 mil reais.

A lista de clientes do banco Uati encontra-se no arquivo ``clientes.csv`` contido neste projeto.


## Requisitos técnicos obrigatórios

- Tela de login;
- Uma tela para cadastrar os usuários que devem receber os alertas;
- Uma tela para importação dos clientes do banco (Upload de CSV);
- Uma tela para controle do monitoramento/dashboard, incluindo gráficos utilizando técnicas de estatística descritiva, sob carteira de clientes, número de alertas e outras funcionalidades que o grupo julgar interessantes;
- Uma tela para listar e detalhar os alertas,  listar os envios de emails e para quem foi enviado, data, hora e outras funcionalidades que o grupo julgar interessantes;
- Enviar um alerta através de e-mail quando um cliente se tornar um funcionário do banco;
- Todas essas funcionalidades devem ser expostas para clientes que queiram integrar através de uma API.