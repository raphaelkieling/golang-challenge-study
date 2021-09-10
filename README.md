## Desafio pessoal

URL: https://github.com/delivery-much/backend-challenge

Pelo desafio proposto pareceu realmente ser apenas uma api muito pequena que será composta por várias outras apis que obviamente irão consumi-la. Os pontos que deve-se tomar cuidado:

- Idempotência
- Tratamento de erros
- Documentação para outros desenvolvedores
    - Para os que irão **desenvolver** nela (README.md)
    - Para os que irão **consumir** ela (Swagger)

## Estrutura
```
# Configuração e documentação
api -> Swagger
conf -> Leitura das variaveis de ambiente
resource -> Arquivos como seed
tools -> Golang files uteis para algumas operações como `populate`

# Serviço
server.go -> Inicia e configura todo o servidor e rotas
```


## Como rodar localmente?

0 - Criar um .env
> Como isso é apenas um teste vou disponibilizar um:
```
APP_PORT=3000

DB_HOST=database                
DB_USER=postgres
DB_PASSWORD=deliverymuch
DB_NAME=deliverymuch
DB_PORT=5432
```

1 - Executar:
```
docker-compose up -f docker-compose.yml
```

2 - Popular:
```
#Docker
docker exec -it <<container-id>> /populate
```

3 - Verificar postgres para garantir que os produtos foram populados