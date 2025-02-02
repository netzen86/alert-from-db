package main

// Package main - пакет сервера
// Приложение для получения и храненния метрик.
// Приложение позволяет хранить метрики в текстовом файле, ОЗУ и в базе данных.

import (
	"alert-from-db/internal/backend"
	"alert-from-db/internal/repository/db"
	"context"
	"log"
	"net"
	_ "net/http/pprof"
)

func main() {
	var err error
	ctx := context.Background()
	storage, err := db.NewDBStorage(ctx, db.DBconstring)
	if err != nil {
		log.Fatalf("error when creating storage %v ", err)
	}

	err = backend.MakeDBMigrations(ctx, storage)
	if err != nil {
		log.Fatalf("error when making migration %v ", err)
	}

	// определяем порт для gRPC сервера
	listen, err := net.Listen(backend.ProtoTCP, backend.EndpointRPC)
	if err != nil {
		log.Fatalf("error when setup net listen %v", err)
	}
	gSRV, err := backend.GetgRPCSrv(ctx, storage)
	if err != nil {
		log.Fatalf("error when gettin gRPC server %v", err)
	}

	log.Printf("!!! SERVER START at %v %v!!!\n", backend.EndpointRPC, backend.ProtoTCP)

	// запуск gRPC сервера
	if err = gSRV.Serve(listen); err != nil {
		log.Fatalf("error when run gRPC server %v", err)
	}
}
