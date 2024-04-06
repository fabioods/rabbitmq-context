<h1 align="center">Welcome to Desafio Event Driven Arch 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
</p>

> Projeto que contem as apps wallet core desenvolvida no curso full cycle e envia dados para um topico de transaction e balances. Além da aplicação balance que consulta o topico de balances e atualiza os dados das contas conforme os valores recebidos.

> Os dados das bases de dados são preenchidos no primeiro start do container, caso deseje limpar os dados e preencher novamente, basta executar o comando `docker-compose down -v` e subir novamente.

> A criação dos tópicos também são feitas na execução do container, utilizando o projeto kafka-setup.

> Para testar as aplicações amabas possuem a pasta api com os CURL's correspondentes, também existe um endpoint geral de teste /ping para ambos os serviços.

> CURL -X POST http://localhost:8080/ping

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
