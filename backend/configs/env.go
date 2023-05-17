package configs

import(
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Введите url базы")
	}

	return os.Getenv("MONGOURI")
}

func EnvPort() string{
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Введите порт сервера")
	}

	return os.Getenv("PORT")
}