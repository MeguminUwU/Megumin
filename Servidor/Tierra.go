package main

import (
	pb "L3/Proto"
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	mutex     sync.Mutex
	AT_bodega int32
	MP_bodega int32
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) MyMethod(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	mutex.Lock() // Bloquea el mutex antes de acceder a los recursos de bodega
	defer mutex.Unlock()
	response := &pb.ResponseMessage{
		Respuesta: solicitarM(req.ID, req.AT, req.MP),
	}
	return response, nil
}

func solicitarM(ID string, AT int32, MP int32) bool {
	if AT_bodega >= AT && MP_bodega >= MP {
		AT_bodega -= AT
		MP_bodega -= MP
		log.Printf("Recepción de solicitud desde equipo %s, %d AT y %d MP ; Resolucion: --APROBADA-- ; AT EN SISTEMA: %d ; MP EN SISTEMA: %d", ID, AT, MP, AT_bodega, MP_bodega)
		return true
	} else {
		log.Printf("Recepción de solicitud desde equipo %s, %d AT y %d MP ; Resolucion: --DENEGADA-- ; AT EN SISTEMA: %d ; MP EN SISTEMA: %d", ID, AT, MP, AT_bodega, MP_bodega)
		return false
	}
}

func main() {
	AT_bodega = 0
	MP_bodega = 0
	intervalo := 5 * time.Second

	go func() {
		for {
			if AT_bodega+10 > 50 {
				AT_bodega = 50
			} else {
				AT_bodega += 10
			}
			if MP_bodega+5 > 20 {
				MP_bodega = 20
			} else {
				MP_bodega += 5
			}
			time.Sleep(intervalo)
		}
	}()

	// Inicia el servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	log.Println("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
