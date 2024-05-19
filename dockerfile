# Usa una imagen base de Golang
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /L3

# Copia el archivo Tierra.go desde la carpeta Servidor
COPY . .

COPY ./Servidor/Tierra.go .
# Compila tu aplicación
RUN go mod download
RUN go build -o Tierra .

# Expone el puerto utilizado por tu servidor gRPC
EXPOSE 50051

# Comando para ejecutar tu servidor gRPC cuando el contenedor se inicie
CMD ["./Tierra"]