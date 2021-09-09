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

1 - Executar:
```
docker-compose up -f docker-compose.yml -f docker-compose.dev.yml
```

2 - Popular:
> Garanta que voce tenha o make instalado na sua máquina. Normalmente ambientes linux já tem. Caso contrario pode executar manualmente usando `go run ./tools/populate/main.go`
```
make populate
```