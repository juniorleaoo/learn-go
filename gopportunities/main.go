package main

import (
	"github.com/juniorleaoo/learn-go/gopportunities/config"
	"github.com/juniorleaoo/learn-go/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	router := config.Init()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	router.Initialize()
}
