<h1 align="center">RabbitMQ with Wallet & balance 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
</p>

> A proposta deste projeto é criar um ambiente com RabbitMQ, Wallet e Balance,
> onde o Wallet é responsável por receber requisições e enviar para o RabbitMQ,
> o Balance é responsável por consumir as mensagens do RabbitMQ e atualizar os salvos das contas.

> Ambientes de teste
> 
> CURL -X POST http://localhost:8080/ping
> 
> CURL http://localhost:3003/ping


## Install

```sh
docker-compose up -d
```

## Author

👤 **Fabio dos santos**


## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
