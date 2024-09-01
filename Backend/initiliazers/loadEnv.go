package initiliazers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("can't load env variable file ")
	}

}
