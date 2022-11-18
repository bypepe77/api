package bootstrap

import (
	"github.com/bypepe77/api/internal/common/database"
	"github.com/bypepe77/api/internal/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Run() error {
	log := logrus.New()

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	db := database.DatabaseConection()

	config := server.NewConfig("Test Api", "localhost", "8080")

	server := server.NewServer(config, db, log)

	return server.Run()
}
