package main

import (
	"context"
	"log"

	ggmanager "github.com/mugioka/slack-bots/pkg/gg-manager"
)

func main() {
	ggm, err := ggmanager.Init()
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()
	err = ggm.Run(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
