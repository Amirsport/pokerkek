package main

import (
	"flag"
	"log"
	"pokerok/internal/config"
	"pokerok/internal/provider"
	"pokerok/internal/server"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "./configs/config.yaml", "путь к конфигу")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	DBprov := provider.CreateDBProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.DBname, cfg.DB.User, cfg.DB.Password)

	srv := server.CreateNewServer(DBprov, cfg.Port, cfg.IP)
	srv.Run()
}
