# Utiliza a imagem oficial do Golang como imagem base
FROM golang:1.21

# Atualiza os pacotes e instala a biblioteca librdkafka-dev
RUN apt-get update && apt-get install -y librdkafka-dev

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

# Copia o diretório do projeto para o diretório de trabalho
COPY . .

# Compila o aplicativo Go para um binário
RUN go build -o wallet ./cmd/walletcore/

# Expõe a porta 8080
EXPOSE 8080

# Utiliza o script de espera para garantir que o MySQL esteja pronto antes de iniciar a aplicação
CMD /wait-for-it.sh mysql:3306 --timeout=90 --strict -- /wait-for-it.sh kafka:9092 --timeout=90 --strict -- ./wallet
