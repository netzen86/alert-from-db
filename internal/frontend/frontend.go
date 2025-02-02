package frontend

import (
	"fmt"

	pb "alert-from-db/proto/backend"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	AgentgRPCEndpoint string = "localhost:3200"
	HTTPport          string = "8000"
)

// GetgRPCCli функция для создания клиента gRPC сервера
func GetgRPCCli() (pb.BackendClient, error) {
	// устанавливаем соединение с сервером
	conn, err := grpc.NewClient(AgentgRPCEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error when connect to server %w", err)
	}
	// получаем переменную интерфейсного типа MetricClient,
	// через которую будем отправлять сообщения
	cli := pb.NewBackendClient(conn)
	return cli, nil
}
