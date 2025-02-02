package backend

import (
	"context"
	"fmt"

	"alert-from-db/internal/data/films"
	"alert-from-db/internal/repository/db"
	pb "alert-from-db/proto/backend"

	"google.golang.org/grpc"
)

const (
	EndpointRPC string = "localhost:3200"
	ProtoTCP    string = "tcp"
)

type BackendServer struct {
	// нужно встраивать тип pb.Unimplemented
	// для совместимости с будущими версиями
	pb.UnimplementedBackendServer

	// используем для доступа к методам работы с сервером
	storage *db.DBStorage
}

// AddFilm реализует интерфейс добавления фильма в хранилище.
func (srv *BackendServer) AddFilm(ctx context.Context, in *pb.AddFilmRequest) (*pb.AddFilmResponse, error) {
	var response pb.AddFilmResponse
	var err error

	f := films.Film{Title: in.Film.Title, Director: in.Film.Directorr}

	err = srv.storage.AddFilmRepo(ctx, f)
	if err != nil {
		response.Error = fmt.Sprintf("addfilm: error when add film  %v", err)
		return &response, err
	}
	return &response, nil
}

// DelFilm реализует интерфейс добавления фильма в хранилище.
func (srv *BackendServer) DelFilm(ctx context.Context, id *pb.DelFilmRequest) (*pb.DelFilmResponse, error) {
	var response pb.DelFilmResponse
	err := srv.storage.DelFilmRepo(ctx, id.Id)
	if err != nil {
		response.Error = fmt.Sprintf("delfilm: error when del film  %v", err)
		return &response, err
	}

	return &response, nil
}

// GetFilms реализует интерфейс для получеия фильмов
func (srv *BackendServer) GetFilms(ctx context.Context, in *pb.GetFilmsRequest) (*pb.GetFilmsResponse, error) {
	var response pb.GetFilmsResponse
	var err error
	var movies []films.Film

	movies, err = srv.storage.GetFilmsRepo(ctx)
	if err != nil {
		response.Error = fmt.Sprintf("getfilm: error when getting all mtrics name %v", err)
		return &response, err
	}

	for _, f := range movies {
		response.Films = append(response.Films, &pb.Film{Id: f.ID, Title: f.Title, Directorr: f.Director})
	}
	return &response, err
}

func GetgRPCSrv(ctx context.Context, storage *db.DBStorage) (*grpc.Server, error) {
	var server BackendServer
	var err error

	server.storage = storage
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterBackendServer(s, &server)
	return s, err
}
