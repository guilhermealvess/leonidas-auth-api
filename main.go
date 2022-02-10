package main

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/adapter/grpc/service"
	"api-auth/src/adapter/repository"
	"api-auth/src/db"
	"context"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:   0,
	})

	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	database := os.Getenv("MONGO_DATABASE")
	uri := "mongodb://" + host + ":" + port + "/" + database
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	redisCache := db.NewCacheRedisInstance(*rdb)
	mongoDB := db.NewMongoDBInstance(*mongoClient, database)

	startGRPCServer(mongoDB, redisCache)
}

func startGRPCServer(db repository.DocumentDB, cache repository.Cache) {
	lis, err := net.Listen("tcp", "localhost:"+os.Getenv("GRPC_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("START GRPC SERVE ON PORT " + os.Getenv("GRPC_PORT"))

	projectService := service.NewProjectServiceGRPC(db, cache)
	accountService := service.NewAccountServiceGRPC(db, cache)

	grpcServer := grpc.NewServer()
	pb.RegisterProjectsServer(grpcServer, projectService)
	pb.RegisterAccountServicesServer(grpcServer, accountService)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
