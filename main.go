package main

import (
	"log"

	"github.com/tigerowo/bytebroad-canvas/config"
	"github.com/tigerowo/bytebroad-canvas/handler"
	"github.com/tigerowo/bytebroad-canvas/router"
	"github.com/tigerowo/bytebroad-canvas/service"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	if err := service.EnsureDefaultAdmin(); err != nil {
		log.Fatal(err)
	}
	if err := service.EnsureDefaultAgentSkills(); err != nil {
		log.Fatal(err)
	}
	service.StartPromptSyncScheduler()
	service.StartCanvasProjectCleanupScheduler()
	handler.StartVideoTaskPoller()
	log.Fatal(router.New().Run(":" + config.Cfg.Port))
}
