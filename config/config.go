package config

import (
	"github.com/vivaldy22/mekar-regis-user-service/tools/viper"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func InitDB() (*sql.DB, error) {
	dbUser := viper.ViperGetEnv("DB_USER", "root")
	dbPass := viper.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := viper.ViperGetEnv("DB_HOST", "localhost")
	dbPort := viper.ViperGetEnv("DB_PORT", "3306")
	schemaName := viper.ViperGetEnv("DB_SCHEMA", "schema")
	driverName := viper.ViperGetEnv("DB_DRIVER", "mysql")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, err := sql.Open(driverName, dbPath)

	if err != nil {
		return nil, err
	}

	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	return dbConn, nil
}

func RunServer(db *sql.DB) {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	// tokenService := token.NewService()
	// authservice.RegisterJWTTokenServer(srv, tokenService)

	// levelService := level.NewService(db)
	// authservice.RegisterLevelCRUDServer(srv, levelService)

	// userService := user.NewService(db)
	// authservice.RegisterUserCRUDServer(srv, userService)

	reflection.Register(srv)

	log.Printf("Starting GRPC Eat N' Fit Auth Server at %v port: %v\n", host, port)
	if err = srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}