package main

import (
	"net/http"
	"os"

	"github.com/funmi4194/instashop/primer"
	"github.com/opensaucerer/barf"
)

func main() {
	// set env path
	if os.Getenv("ENV_PATH") == "" {
		os.Setenv("ENV_PATH", ".env")
	}

	// load environment variables
	if err := barf.Env(primer.ENV, os.Getenv("ENV_PATH")); err != nil {
		barf.Logger().Fatalf(`[main.main] [barf.Env(primer.ENV, os.Getenv("ENV_PATH"))] %s`, err.Error())
	}

	// configure barf
	if err := barf.Stark(barf.Augment{
		Port:         primer.ENV.Port,
		Logging:      barf.Allow(), // enable request logging
		Recovery:     barf.Allow(), // enable panic recovery
		WriteTimeout: 30,
		ReadTimeout:  30,
		CORS: &barf.CORS{
			AllowedOrigins: []string{
				"*",
			},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodOptions,
			},
		},
	}); err != nil {
		barf.Logger().Fatalf(`[main.main] [barf.Stark(barf.Augment)] %s`, err.Error())
	}

}
