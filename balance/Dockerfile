# Utiliza a imagem oficial do Golang como imagem base
FROM golang:1.22.2

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o script de espera para o container
COPY wait-for-it.sh /wait-for-it.sh

# Torna o script de espera executável
RUN chmod +x /wait-for-it.sh

# Copia os arquivos do módulo Go e o arquivo go.sum para o diretório atual
COPY go.mod ./
COPY go.sum ./

# Baixa as dependências de Go
RUN go mod download

# Expõe a porta 3003
EXPOSE 3003

CMD /wait-for-it.sh mysql:3306 --timeout=90 --strict && \
    /wait-for-it.sh rabbitmq:5672 --timeout=90 --strict && \
    /wait-for-it.sh walletcore:8080 --timeout=500 --strict && \
    go run ./cmd/balance/
