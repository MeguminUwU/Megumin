Grupo 14
Joaquin Contreras 201973527-6
Martin Rodriguez 201973620-5

Instrucciones de ejecucion:

1ro Tierra
  -Tener docker desktop abierto
  -Abrir terminal dentro de la carpeta L3
  -Construir la imagen usando el comando: "docker build -t tierra ."
  -Prender el contenedor con el comando: "docker run -d -p 50051:50051 tierra"

2do Equipos
  -Abrir terminal dentro de la carpeta Cliente
  -Ejecutar archivo go con el comando: "go run Equipo.go"