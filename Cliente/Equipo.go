package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	pb "L3/Proto"

	"google.golang.org/grpc"
)

// Dirección del servidor gRPC
const serverAddress = "localhost:50051"

// Función para enviar solicitudes al servidor gRPC
func enviarSolicitud(client pb.MyServiceClient, id string, at int32, mp int32, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	defer wg.Done()

	for {
		// Crea una solicitud con el ID, AT y MP proporcionados
		solicitud := &pb.RequestMessage{
			ID: id,
			AT: at,
			MP: mp,
		}

		// Realiza la llamada al servidor gRPC
		respuesta, err := client.MyMethod(context.Background(), solicitud)
		if err != nil {
			log.Printf("Error al llamar al servidor gRPC: %v", err)
			continue
		}
		// Verifica la respuesta del servidor
		if respuesta.Respuesta {
			log.Printf("Equipo %s solicitando %d AT y %d MP ; Resolucion: --APROBADA-- ; Conquista Exitosa!, cerrando comunicación", id, at, mp)
			break
		} else {
			log.Printf("Equipo %s solicitando %d AT y %d MP ; Resolucion: --DENEGADA-- ; Reintentando en 3 segs...", id, at, mp)
			time.Sleep(3 * time.Second) // Espera 3 segundos antes de volver a intentar
		}
	}
}

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor gRPC: %v", err)
	}
	defer conn.Close()

	// Crea un cliente gRPC
	cliente := pb.NewMyServiceClient(conn)

	var wg sync.WaitGroup

	// Inicializa cuatro subprocesos para enviar solicitudes al servidor
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		rand.Seed(time.Now().UnixNano())
		at := int32(rand.Intn(11) + 20)
		mp := int32(rand.Intn(6) + 10)
		go enviarSolicitud(cliente, strconv.Itoa(i), at, mp, &wg)
	}

	// Espera a que todos los procesos terminen
	wg.Wait()
	log.Printf("Todos los equipos estan listos.")
}
