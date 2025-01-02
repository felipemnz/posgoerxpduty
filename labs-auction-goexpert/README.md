# fc-pos-go-leilao
Terceiro lab pós go expert (leilão)

## NOTA:
  * Eu não consertei nenhum (mentira, só um) dos problemas já existentes no programa inicial disponibilizado, apenas adicionei a feature nova pedida.
  * Por favor, não falhar o lab por problemas provenientes do projeto inicial (como a falta de uma rota pra criar user ou o filtro de status de auctions não funcionando na rota de get auction)

## Requerimentos
  * golang versão 1.22.3 ou superior
  * ou docker

## Rodando o projeto
  * Pode rodar com `docker-compose up`. Subirá o servidor e o banco mongodb necessários.

## Rodar testes automáticos
  * Pode-se testar a função nova rodando os unit testes adicionados com `go test ./...`

## Funcionamento
  * Após a criação de uma nova auction, ele irá atualizar após o tempo definido no arquivo `.env` no caminho `cmd/auction/.env` pela variável `AUCTION_INTERVAL=20s`. Caso deseje mudar esse tempo, basta alterar essa variável (e subir o projeto novamente com `docker-compose up --build`)
  * Pode-se conferir que, após o tempo estipulado anteriormente, o valor do status da auction é alterado para 1 (Completed)

  Ex:
  * Para criar uma nova auction:
     ```json
     POST http://localhost:8080/auction

    {
    "product_name": "notebook",
    "category": "eletronics",
    "description": "a normal notebook",
    "condition": 0
    }
     ```

     

  * Para buscar as auctions:

     ```json
     GET http://localhost:8080/auction?status=0
     ```
  * Se tiver o id de uma auction (pela request acima, por exemplo) pode vê-la em específico pela requisição:

     ```json
     GET http://localhost:8080/auction/winner/:auctionId
     ```
  Substituindo a variável :auctionId pelo id da auction

OBS: Lembrando, apenas a funcionalidade nova foi adicionada. Os pontos que não estiverem funcionando além da funcionalidade de compleção automática da auction é porque já não estavam funcionando no projeto original... (tem muita coisa sem funcionar)
