package config

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"github.com/joho/godotenv"
	"os"
)

var logging = log.GetLogger()

type Env struct {
	JwtSecret string `json:"jwt___secret,omitempty"`
}
func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		logging.Fatal(err)
		panic(err)
	}
	return &Env{
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}