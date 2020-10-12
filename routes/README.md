# Routes #

Routes é uma API para consulta e cadastro de rotas de um ponto ao outro. Também é possível consultar a rota mais barata, indiferente do número de escalas que possa existir.


## Modo de utilização ##

Para inicializar a aplicação é necessário informar o arquivo onde as rotas serão armazenadas. O arquivo informado poderá já conter rotas desde que obedeça o seguinte fomato:

    ORIGEM,DESTINO,VALOR

### Exemplo do conteúdo: ###
```csv
GRU,BRC,10
BRC,SCL,5
GRU,CDG,75
GRU,SCL,20
GRU,ORL,56
ORL,CDG,5
SCL,ORL,20
```

**Obs: Nesta API são aceitos apenas valores interos.**


# Execução ##

## Inicialização por terminal: ##
```shell
    $ routes input-routes.csv
    Listening :8080
    please enter the route:
``` 
### Consultando pelo terminal ###

Através do terminal é possível consultar apenas consultar o custo de uma rota, inteferente da quantidades de escalas que será nessárias. Para informar uma rota, foi definido o formato DE-PARA

```shell
    $ routes input-routes.csv
    Listening :8080
    please enter the route: GRU-CDG
    best route: GRU - BRC - SCL - ORL - CDG > $40
``` 

## Interface REST ##

### Cadastro de rota ###

Para cadastra um nova rota deve ser realização uma requisição POST para ```http://{host}:8080/route```.

```http 
    POST  http://{host}:8080/route
```

O body da requisição deve conter um JSON com no seguinte formato:
```json 
    {
        "origin":"ITP",
        "destination":"BCA",
        "cost":1000
    }
    
```

### Consulta de rota ###

Para consultar uma rota cadastrada basta realizar uma requisião GET para ```http://{host}:8080/route/{rota-consultada}```, informando após o /route/ a roda que deseja consultar no formato DE-PARA.

```http 
    GET  http://{host}:8080/route/GRU-CDG
```

Caso a rota pesquisada não esteja cadastrada será retornada StatusCode 404-Not Found. Se encontrada retornará um JSON no seguinte formato:

```json 
    {
        "origin":"GRU",
        "destination":"CDG",
        "cost":75
    }
    
```

### Consulta de rota mais barata ###

Para realizar a consulta de rotas mais bairata, desconsiderando o número de escalas deve ser utilizado o Query Param ```?cheapest=true```

```http 
    GET  http://{host}:8080/route/GRU-CDG?cheapest=true
```

Em caso de sucesso será retornad um JSON no seguinte formato:
```json 
    {
        "bestRoute":"GRU - BRC - SCL - ORL - CDG",
        "cost":40
    }
    
```

----
----

## Estrutura de pacotes ##

- handler: responsável pelos recusos de comunicação HTTP;
- scale: responsável por gestão do cache e também da persistência dos rotas
- util: reponsável pelos procedimentos auxiliares de operação

