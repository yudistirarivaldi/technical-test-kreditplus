package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/yudistirarivaldi/technical-test-kreditplus/config"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/handler"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/middleware"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/repository"
	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/service"
)

type databaseConnections struct {
	mysql *sql.DB
}

type appServices struct {
	consumerService    *service.ConsumerService
	authService        *service.AuthService
	transactionService *service.TransactionService
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbs, err := initDatabases(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize databases: %v", err)
	}
	defer closeDatabases(dbs)

	services := initServices(dbs, cfg)
	startHTTPServer(cfg, dbs, services)
}

func initDatabases(cfg *config.Config) (*databaseConnections, error) {
	dbs := &databaseConnections{}

	dbMysql, err := config.NewMySQLConnection(cfg.DatabaseMysql)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	dbs.mysql = dbMysql
	log.Println("Connected to MySQL")

	return dbs, nil
}

func closeDatabases(dbs *databaseConnections) {
	if dbs.mysql != nil {
		dbs.mysql.Close()
	}
}

func initServices(dbs *databaseConnections, cfg *config.Config) *appServices {
	consumerRepo := repository.NewConsumerRepository(dbs.mysql)
	authRepo := repository.NewAuthRepository(dbs.mysql)
	transactionRepo := repository.NewTransactionRepository(dbs.mysql)
	consumerLimitRepo := repository.NewConsumerLimitRepository(dbs.mysql)

	consumerService := service.NewConsumerService(consumerRepo)
	authService := service.NewAuthService(authRepo, consumerLimitRepo, cfg.JWT.Secret)
	transactionService := service.NewTransactionService(dbs.mysql, transactionRepo, consumerLimitRepo)

	return &appServices{
		consumerService:    consumerService,
		authService:        authService,
		transactionService: transactionService,
	}
}

func startHTTPServer(cfg *config.Config, dbs *databaseConnections, services *appServices) {
	consumerHandler := handler.NewConsumerHandler(services.consumerService)
	authHandler := handler.NewAuthHandler(services.authService)
	transactionHandler := handler.NewTransactionHandler(services.transactionService)

	http.HandleFunc("/api/auth/register", authHandler.HandleRegister)
	http.HandleFunc("/api/auth/login", authHandler.HandleLogin)

	http.HandleFunc("/api/consumers", middleware.JWTMiddleware(cfg.JWT.Secret, consumerHandler.HandleGetProfile))
	http.HandleFunc("/api/consumers/update", middleware.JWTMiddleware(cfg.JWT.Secret, consumerHandler.HandleUpdateConsumer))

	http.HandleFunc("/api/transactions", middleware.JWTMiddleware(cfg.JWT.Secret, transactionHandler.HandleInsertTransaction))
	http.HandleFunc("/api/transactions/history", middleware.JWTMiddleware(cfg.JWT.Secret, transactionHandler.HandleGetTransactionsByConsumer))

	log.Printf("Server starting on port %s...", cfg.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
