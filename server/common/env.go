package common

import (
	"log"
	"os"
	"regexp"
	"syscall"

	"github.com/joho/godotenv"
)

func EnvString(key, fallback string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	return fallback
}

type Config struct {
	GatewayServiceHost  string
	ProblemsServiceHost string
	HTTPPort            string
	GRPCPort            string
}

var projectDirName = "server"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file at %v", string(rootPath))
	}
}

func GetConfig() *Config {
	loadEnv()
	httpPort := EnvString(("HTTP_PORT"), "8080")
	grpcPort := EnvString(("GRPC_PORT"), "3000")
	problemsHost := EnvString(("PROBLEMS_HOST"), "problems")
	gatewayHost := EnvString(("GATEWAY_HOST"), "gateway")

	return &Config{
		HTTPPort:            httpPort,
		GRPCPort:            grpcPort,
		ProblemsServiceHost: problemsHost,
		GatewayServiceHost:  gatewayHost,
	}
}
