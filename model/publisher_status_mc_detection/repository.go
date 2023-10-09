package publisher_status_mc_detection

import (
	"context"
	"encoding/json"

	"github.com/aditya3232/gatewatchApp-services.git/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	CreateQueueStatusMcDetection(input PublisherStatusMcInput) (StatusMcDetection, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueStatusMcDetection(input PublisherStatusMcInput) (StatusMcDetection, error) {
	ch, err := connection.RabbitMQ().Channel()
	if err != nil {
		return StatusMcDetection{}, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"StatusMcDetectionQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return StatusMcDetection{}, err
	}

	// body from input
	var inputReadytoMarshal = PublisherStatusMcInput{
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		StatusSignal:  input.StatusSignal,
		StatusStorage: input.StatusStorage,
		StatusRam:     input.StatusRam,
		StatusCpu:     input.StatusCpu,
	}

	// Convert the StatusMc struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return StatusMcDetection{}, err
	}

	ctx := context.Background() // Create a context
	err = ch.PublishWithContext(ctx,
		"",
		"StatusMcDetectionQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return StatusMcDetection{}, err
	}

	// Return the sent StatusMcDetection struct
	return StatusMcDetection{}, err
}
