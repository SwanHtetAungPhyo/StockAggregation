package config

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"os"
)

var logging = log.GetLogger()

type Env struct {
	JwtSecret string `json:"jwt___secret,omitempty"`
}

func GetEnv() *Env {

	return &Env{
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
