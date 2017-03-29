package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"

	"github.com/danielqsj/kubernetes-slackbot/pkg/kubernetes"
	"github.com/danielqsj/kubernetes-slackbot/pkg/slack"
	"github.com/danielqsj/kubernetes-slackbot/server/options"
)

type SlackBotServer struct {
	healthzPort    int
	configMap      string
	botToken       string
	kubeConfigFile string
	slackBot       slack.SlackBot
}

func NewSlackBotServerDefault(config *options.SlackBotServerConfig) *SlackBotServer {
	s := SlackBotServer{
		healthzPort:    config.HealthzPort,
		configMap:      config.ConfigMap,
		botToken:       config.BotToken,
		kubeConfigFile: config.KubeConfigFile,
	}
	output, err := kubernetes.TestConnection(s.kubeConfigFile)
	if err != nil {
		log.Fatal("Connect to kubernetes master failed\n")
	} else {
		log.Printf("Connect to kubernetes master successful:\n%s\n", output)
	}
	s.slackBot = slack.NewSlackBot(s.botToken)
	return &s
}

func (server *SlackBotServer) Run() {
	pflag.VisitAll(func(flag *pflag.Flag) {
		log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
	//setupSignalHandlers()

	go server.handleSigterm()
	server.setupHealthzHandlers()
	log.Printf("Setting up Healthz Handler(/readiness, /cache) on port :%d", server.healthzPort)

	server.start()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.healthzPort), nil))
}

// setupHealthzHandlers sets up a readiness and liveness endpoint for kube2sky.
func (server *SlackBotServer) setupHealthzHandlers() {
	http.HandleFunc("/readiness", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "ok\n")
	})
	http.HandleFunc("/cache", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "ok\n")
	})
}

// setupSignalHandlers runs a goroutine that waits on SIGINT or SIGTERM and logs it
// program will be terminated by SIGKILL when grace period ends.
func setupSignalHandlers() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("Received signal: %s, will exit when the grace period ends", <-sigChan)
	}()
}

func (server *SlackBotServer) start() {
	slack.InitSlackLog()
	go server.slackBot.RunSlackRTMServer(server.kubeConfigFile)
}

func (server *SlackBotServer) handleSigterm() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	<-signalChan
	log.Printf("Received SIGTERM, shutting down")

	exitCode := 0
	if err := server.stop(); err != nil {
		log.Printf("Error during shutdown %v", err)
		exitCode = 1
	}

	log.Printf("Exiting with %v", exitCode)
	os.Exit(exitCode)
}

func (server *SlackBotServer) stop() error {
	return nil
}
