package main

import (
	"log"

	"github.com/treehole-backend/envconfig"

	"github.com/treehole-backend/platform/authenticator"
	restful "github.com/treehole-backend/web/app"
)

func main() {
	var err error
	var env envconfig.Env
	if err = envconfig.Init(&env); err != nil {
		log.Fatal("- Failed to load config from environment variables ", err)
	}
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("- Failed to initialize the authenticator: %v", err)
	}
	r, err := restful.Register(env, auth)
	if err != nil {
		log.Fatal("- Failed to register routes ", err)
	}

	log.Print("+ Starting server on port ", env.Port)
	if err := r.Run(":" + env.Port); err != nil {
		log.Fatal("- Failed to start server ", err)
	}

	// TODO - graceful shutdown 可以研究
}
