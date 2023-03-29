package broker

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	amqpBackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpBroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	configinstance "github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/locks/eager"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"strings"
)

type AmqpBroker struct {
	machineryBroker
}

var _ queue.IBroker = (*AmqpBroker)(nil)

func NewAmqpBroker(connectionName string) (*AmqpBroker, error) {
	brokerUrl, queueName, amqpConfig := getAmqpConfig(connectionName)

	cnf := &configinstance.Config{
		Broker:          brokerUrl,
		DefaultQueue:    queueName,
		ResultBackend:   brokerUrl,
		ResultsExpireIn: int(config.GetDuration("queue.result_expire").Seconds()),
		AMQP:            &amqpConfig,
	}

	broker := amqpBroker.New(cnf)
	backend := amqpBackend.New(cnf)
	lock := eager.New()

	b := &AmqpBroker{
		machineryBroker{
			server:           machinery.NewServer(cnf, broker, backend, lock),
			jobMap:           container.MustMakeAs("queue.job_map", queue.IJobMap(nil)),
			defaultQueueName: queueName,
			connectionName:   connectionName,
		},
	}
	if err := b.registerTasks(); err != nil {
		return nil, err
	}

	return b, nil
}

func getAmqpConfig(connectionName string) (brokerUrl, queueName string, amqpConfig configinstance.AMQPConfig) {
	keyPrefix := fmt.Sprintf("queue.connections.%s.", connectionName)
	host := config.GetString(keyPrefix + "host")
	username := config.GetString(keyPrefix + "username")
	password := config.GetString(keyPrefix + "password")
	port := config.GetInt(keyPrefix+"port", 0)
	vhost := config.GetString(keyPrefix+"vhost", "/")

	brokerUrl = fmt.Sprintf("amqp://%s:%s@%s:%d/%s", username, password, host, port, strings.TrimLeft(vhost, "/"))
	queueName = GetQueueName(connectionName, "")
	amqpConfig = configinstance.AMQPConfig{
		Exchange:         config.GetString(keyPrefix+"exchange", "machinery_exchange"),
		ExchangeType:     config.GetString(keyPrefix+"exchange_type", "direct"),
		QueueDeclareArgs: nil,
		QueueBindingArgs: nil,
		BindingKey:       config.GetString(keyPrefix+"binding_key", "machinery_task"),
		PrefetchCount:    config.GetInt(keyPrefix+"prefetch_count", 1),
		AutoDelete:       config.GetBool(keyPrefix+"auto_delete", false),
	}

	return
}
