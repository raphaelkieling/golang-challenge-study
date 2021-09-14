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

## Referencias

Este é meu primeiro projeto go tirando os estudos de sintaxe. Então vou tentar colocar aqui referencias que estou utilizando para basear minhas desições, algumas são chutes ou experiencias proprias em outras linguagens (kotlin, java, javascript, php, ...).

- https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go/25162021
- https://www.youtube.com/watch?v=iYrMAVSOkkA
- https://www.youtube.com/watch?v=LvgVSSpwND8&t=709s
- https://www.youtube.com/watch?v=YEKjSzIwAdA&t=751s
- https://www.youtube.com/watch?v=DqHb5KBe7qI
- https://github.com/golang-standards/project-layout