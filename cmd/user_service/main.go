package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	userservice "github.com/manc88/demo/internal/user_service"
	grpctransport "github.com/manc88/demo/internal/user_service/grpc_transport"
	messagebroker "github.com/manc88/demo/internal/user_service/message_broker"
	"github.com/manc88/demo/internal/user_service/repo"
	"github.com/manc88/demo/internal/user_service/repo/cache"
	"github.com/manc88/demo/internal/user_service/repo/storage"
	"github.com/manc88/demo/pkg/kafka"
	"github.com/manc88/demo/pkg/pgx"
	"github.com/manc88/demo/pkg/redis"
)

const (
	_TRANSPORT_CONFIG    = "./configs/grpc.yaml"
	_PG_CONFIG           = "./configs/pg.yaml"
	_REDIS_CONFIG        = "./configs/redis.yaml"
	_KAFKA_CONFIG        = "./configs/kafka.yaml"
	_USER_SERVICE_CONFIG = "./configs/user_service.yaml"
)

const (
	_DEFAULT_RESPONSE_TIMEOUT = 3
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	logger := log.New(os.Stdout, "[USER_SERVICE] ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("version 001")

	//SERVICE_CONFIG
	usConfig, err := userservice.NewConfig(_USER_SERVICE_CONFIG)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	//PG STORAGE
	pgConfig, err := pgx.NewConfigFromFile(_PG_CONFIG)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	pool, err := pgx.NewPGX(ctx, pgConfig)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	store := storage.NewPgStorage(ctx, pool)
	store.SetLogger(logger)

	//CACHE
	cacheConfig, err := redis.NewConfig(_REDIS_CONFIG)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	red := redis.NewRedis(cacheConfig)
	userCache := cache.NewUserCache(red, usConfig.CacheUsersListKey)

	//REPOSITORY
	r := repo.NewUserRepository(time.Second * _DEFAULT_RESPONSE_TIMEOUT)
	r.SetStorage(store)
	r.SetCache(userCache)
	r.SetLogger(logger)

	//BROKER
	kConfig, err := kafka.NewConfig(_KAFKA_CONFIG)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	b := messagebroker.New(kafka.New(kConfig))

	//SERVICE
	us := userservice.New(ctx, usConfig)
	us.SetRepository(r)
	us.SetMessageBroker(b)
	us.SetLogger(logger)

	//TRANSPORT
	grpcConfig, err := grpctransport.NewConfig(_TRANSPORT_CONFIG)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	t := grpctransport.New(ctx, grpcConfig)
	t.SetService(us)
	log.Fatal(t.Serve())
}
