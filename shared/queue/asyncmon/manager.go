package asyncmon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/azeezlala/assessment/shared/pubsub"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

type Manager struct {
	client    *asynq.Client
	server    *asynq.Server
	mux       *asynq.ServeMux
	redisOpts *asynq.RedisClientOpt
}

func NewManager(redisUrl string) *Manager {
	redisOpt := asynq.RedisClientOpt{Addr: redisUrl}
	return &Manager{
		client:    asynq.NewClient(redisOpt),
		mux:       asynq.NewServeMux(),
		redisOpts: &redisOpt,
	}
}

func (m Manager) Publish(event string, payload interface{}) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	opts := []asynq.Option{
		asynq.TaskID(uuid.NewString()),
		asynq.Timeout(3 * time.Minute),
		asynq.Queue("high"),
		asynq.Retention(1 * time.Hour),
	}

	task := asynq.NewTask(event, b, opts...)
	_, err = m.client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %v", err)
	}

	return nil
}

func (m Manager) Subscribe(event string, handler pubsub.Handler) {
	wrappedHandler := func(ctx context.Context, task *asynq.Task) error {
		fmt.Println("publishing...")
		var options interface{}
		if err := json.Unmarshal(task.Payload(), &options); err != nil {
			return err
		}
		// Call the original pkg.Handler
		handler(ctx, pubsub.Options{
			Payload: task.Payload(),
		})
		return nil
	}
	m.mux.HandleFunc(event, wrappedHandler)
}

// Start starts the task queue
func (m Manager) Start() {
	m.server = asynq.NewServer(
		m.redisOpts,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"high": 10,
			},
		},
	)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal("there's a problem with asynq server")
			}
		}()
		if err := m.server.Run(m.mux); err != nil {
			log.Fatal(err)
		}
	}()
}

var _ pubsub.IPubSub = Manager{}
