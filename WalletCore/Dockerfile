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

# O diretório do projeto será montado como um volume, então não é necessário copiá-lo

# Expõe a porta 8080
EXPOSE 8080

# Compila e executa o aplicativo em tempo real quando o container é iniciado
CMD /wait-for-it.sh mysql:3306 --timeout=90 --strict -- /wait-for-it.sh rabbitmq:5672 --timeout=90 --strict -- go run ./cmd/walletcore/
