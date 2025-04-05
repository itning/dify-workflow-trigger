package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/tmaxmax/go-sse"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

type Config struct {
	Name  string `json:"name"`
	Cron  string `json:"cron"`
	URL   string `json:"url"`
	Token string `json:"token"`
	Body  Body   `json:"body"`
}

type Body struct {
	Inputs       map[string]interface{} `json:"inputs"`
	ResponseMode string                 `json:"response_mode"`
	User         string                 `json:"user"`
}

type Task struct {
	Config Config
	job    gocron.Job
}

type AppContext struct {
	Tasks     []Task
	Scheduler gocron.Scheduler
}

func (c *AppContext) Init() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("Create scheduler failed: %v", err)
		return
	}
	c.Scheduler = s
}

func (c *AppContext) Start() {
	c.Scheduler.Start()
}

func (c *AppContext) Shutdown() {
	err := c.Scheduler.Shutdown()
	if err != nil {
		log.Printf("Shutdown scheduler failed: %v", err)
		return
	}
	log.Println("Scheduler shutdown")
}

func (c *AppContext) New(task Task) {
	j, err := (c.Scheduler).NewJob(
		gocron.CronJob(task.Config.Cron, true),
		gocron.NewTask(task.Execution),
		gocron.WithName(task.Config.Name),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
		gocron.WithEventListeners(
			gocron.BeforeJobRuns(
				func(jobID uuid.UUID, jobName string) {
					log.Printf("[%s][%s] Job started", jobID, jobName)
				},
			),
			gocron.AfterJobRuns(
				func(jobID uuid.UUID, jobName string) {
					log.Printf("[%s][%s] Job finished", jobID, jobName)
				},
			),
		),
	)
	if err != nil {
		log.Fatalf("[][%s] Create job failed: %v", task.Config.Name, err)
		return
	}
	task.job = j
	c.Tasks = append(c.Tasks, task)
}

func (t *Task) Execution() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jsonBytes, err := json.Marshal(t.Config.Body)
	if err != nil {
		log.Printf("[%s][%s] Failed to parse the request body: %v", t.job.ID(), t.Config.Name, err)
		return
	}
	reqBody := bytes.NewBuffer(jsonBytes)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, t.Config.URL, reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+t.Config.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[%s][%s] Request failed: %v", t.job.ID(), t.Config.Name, err)
		return
	}
	defer res.Body.Close()
	mediaType, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
	if err != nil {
		log.Printf("[%s][%s] Failed to parse Content-Type [%s] %v", t.job.ID(), t.Config.Name, res.Header.Get("Content-Type"), err)
	}
	if mediaType == "text/event-stream" {
		for ev, err := range sse.Read(res.Body, &sse.ReadConfig{MaxEventSize: 1024 * 1024}) {
			if err != nil {
				log.Printf("[%s][%s] SSE Error: %v", t.job.ID(), t.Config.Name, err)
				break
			}
			log.Printf("[%s][%s] %s", t.job.ID(), t.Config.Name, ev.Data)
		}
	} else {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("[%s][%s] Failed to read response body: %v", t.job.ID(), t.Config.Name, err)
			return
		}
		log.Printf("[%s][%s] %s", t.job.ID(), t.Config.Name, bodyBytes)
	}
}

func ParseConfigurationFiles(configFilePath string) []Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("Can not open file %s: %v\n", configFilePath, err)
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Can not read file content: %v", err)
		return nil
	}
	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	_, err = file.Read(data)
	if err != nil {
		log.Printf("Read file failed: %v", err)
		return nil
	}

	var configs []Config
	err = json.Unmarshal(data, &configs)
	if err != nil {
		log.Printf("JSON analysis failed: %v", err)
		return nil
	}

	return configs
}

func GetConfigPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Unable to get the current directory: %v", err)
		return ""
	}
	return filepath.Join(currentDir, "config.json")
}

func main() {
	configFilePtr := flag.String("config", GetConfigPath(), "path to the configuration file")

	flag.Parse()

	configFile := *configFilePtr
	if configFile == "" {
		log.Fatalf("Configuration file path is required")
		return
	}
	configs := ParseConfigurationFiles(configFile)
	if configs == nil {
		return
	}
	appContext := AppContext{}
	appContext.Init()

	for _, config := range configs {
		appContext.New(Task{Config: config})
	}
	appContext.Start()

	for _, task := range appContext.Tasks {
		t, _ := task.job.NextRun()
		log.Printf("Created job: Id [%s] Name [%s] NextRunTime [%s]", task.job.ID(), task.job.Name(), t)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		appContext.Shutdown()
		os.Exit(0)
	}()

	select {}
}
