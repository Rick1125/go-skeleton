package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/urfave/cli.v2"
	"pkg.cocoad.mobi/x/http"
	"pkg.cocoad.mobi/x/log"
)

func main() {
	app := &cli.App{
		Name:     "app",
		Version:  "v1.0.0",
		Compiled: time.Now().UTC(),
		Usage:    usage(),
		Flags:    flags(),
		Commands: commands(),
	}
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	if c.Bool("dev") {
		log.SetMode(log.Develop)
		log.SetLevel(log.DebugLevel)
	}
	fmt.Print(fmt.Sprintf(Banner, Version))
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Info("starting to run " + c.App.Name)
	cfg, err := NewConfig(c.App.Name, c.String("config"))
	if err != nil {
		log.Fatal(err)
	}
	app, err := NewApplication(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//here we start a http server to handle healthy state check, maybe reload or start/stop collector.
	if cfg.Health != "" {
		go startMetrics(cfg.Health)
	}
	app.Run()

	return nil
}

func startMetrics(listen string) {
	server, _ := xhttp.NewHTTPServer(listen)
	server.Handle("/metrics", promhttp.Handler())
	if err := server.ServeHTTP(); err != nil {
		log.Warn(err)
	}
}

func usage() string {
	return "usage"
}

func flags() []cli.Flags {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
		&cli.BoolFlag{
			Name:  "dev",
			Usage: "develop mode or production",
			Value: false,
		},
	},
}

func commands() []*cli.Command {
	return []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "run app",
				Action:  run,
			},
	}
}
