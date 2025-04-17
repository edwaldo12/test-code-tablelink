package main

import (
	"context"
	"log"
	"net"

	pb "test_tablelink/internal/delivery/grpc"
	server "test_tablelink/internal/delivery/server"
	"test_tablelink/internal/repository"
	"test_tablelink/internal/usecase"
	"test_tablelink/pkg/config"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	app := fiber.New()

	dbPool, err := config.NewPostgresPool()
	if err != nil {
		log.Fatalf("Error establishing DB connection: %v", err)
	}
	defer dbPool.Close()

	ingredientRepo := repository.NewIngredientRepo(dbPool)
	itemRepo := repository.NewItemRepo(dbPool)

	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo)

	go func() {
		lis, err := net.Listen("tcp", ":9001")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
				log.Printf("Received request: %s", info.FullMethod)
				return handler(ctx, req)
			}),
		)
		pb.RegisterIngredientServiceServer(grpcServer, server.NewIngredientServer(ingredientUsecase))
		pb.RegisterItemServiceServer(grpcServer, server.NewItemServer(itemUsecase))

		reflection.Register(grpcServer)

		log.Println("gRPC server listening on :9001")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Start HTTP server
	log.Println("HTTP server listening on :3001")
	log.Fatal(app.Listen(":3001"))
}
