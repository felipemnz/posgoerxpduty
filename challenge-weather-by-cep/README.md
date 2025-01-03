# Desafio Sistema de temperatura por CEP

***Objetivo:*** Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

***Requisitos:***
- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
   - Em caso de sucesso:
      - Código HTTP: 200
      - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
   - Em caso de falha, caso o CEP não seja válido (com formato correto):
      - Código HTTP: 422
      - Mensagem: invalid zipcode
   - ​​​Em caso de falha, caso o CEP não seja encontrado:
      - Código HTTP: 404
      - Mensagem: can not find zipcode
- Deverá ser realizado o deploy no Google Cloud Run.

***Dicas:***
- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
   - Sendo F = Fahrenheit
   - Sendo C = Celsius
   - Sendo K = Kelvin

***Entrega:***
- O código-fonte completo da implementação.
- Documentação explicando como rodar o projeto em ambiente dev e production.
- Testes automatizados demonstrando o funcionamento.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

### Execução Local

1. Clone o repositório e acesse a subpasta do desafio:

   ```bash
   git clone https://github.com/felipemnz/posgoerxpduty.git
   cd posgoerxpduty/challenge-weather-by-cep
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Execute o servidor:

   ```bash
   go run main.go
   ```

4. Execute a curl em outro terminal ou abra a url no navegador:

   ```bash
   curl --request GET --url 'http://localhost:8080/<CEP>'
   ou
   http://localhost:8080/<CEP>
   ```

### Execução Local via Docker-Compose
1. Suba o Docker Compose
   ```bash
   docker compose up
   ```

2. Execute a curl em outro terminal ou abra a url no navegador
   ```bash
   curl --request GET --url 'http://localhost:8080/<CEP>'
   ou
   http://localhost:8080/<CEP>
   ```

### Execução no Cloud-Run
1. Abra a URL abaixo:

   ```bash
   # invalid zip code
   https://test-weather-by-cep-490185259293.us-central1.run.app/111

   # zip code not found
   https://test-weather-by-cep-490185259293.us-central1.run.app/01153001

   # success
   https://test-weather-by-cep-490185259293.us-central1.run.app/13450402
   ```
