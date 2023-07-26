package env

import (
	"os"

	"cyberpull.com/gotk/v3/log"

	"github.com/joho/godotenv"
)

func init() {
	_, err := os.Stat(".env")

	if err != nil {
		return
	}

	err = godotenv.Load()

	if err != nil {
		log.Println(err)
	}
}
