package mqconfig

const defaultQueueName = "webhook"

type MqConfig struct {
	ConnectionStr ConnectionString
	QueueName     string
}

func NewMqConfig(c ConnectionString) (MqConfig, error) {
	err := c.validate()
	if err != nil {
		return MqConfig{}, err
	}
	return MqConfig{c, defaultQueueName}, nil
}
