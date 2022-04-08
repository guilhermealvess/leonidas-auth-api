package main

import (
	"api-auth/src/adapter/grpc/pb"
	"api-auth/src/adapter/grpc/service"
	"api-auth/src/adapter/jwt"
	"api-auth/src/adapter/repository"
	"api-auth/src/adapter/rest"
	"api-auth/src/db"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
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

	go startHttpServer(mongoDB, redisCache)

	startGRPCServer(mongoDB, redisCache)
}

func startGRPCServer(db repository.DocumentDB, cache repository.Cache) {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("START GRPC SERVE ON PORT " + port)

	grpcServer := grpc.NewServer()
	services := service.NewAccountServiceGRPC(db, cache, jwt.NewJWTMaker())
	pb.RegisterApiV1ServicesServer(grpcServer, services)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func startHttpServer(db repository.DocumentDB, cache repository.Cache) {
	port := os.Getenv("HTTP_SERVER_PORT")
	log.Println("START HTTP SERVE ON PORT " + port)
	router := mux.NewRouter()

	accountController := rest.NewAccountController(db, cache)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") })
	router.HandleFunc("/account/activation-link", accountController.ActivationAccount).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
