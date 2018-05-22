package mqconfig

import (
	"errors"
	"fmt"
)

const (
	MIN_TCP_PORT = 0
	MAX_TCP_PORT = 65535
)

const PROTOCOL_AMQP = "amqp"

var supportedProtocols = []string{PROTOCOL_AMQP}

type ConnectionString struct {
	User     string
	Pass     string
	Host     string
	Port     int
	Protocol string
}

func isSupportedProtocol(protocol string) bool {
	for _, v := range supportedProtocols {
		if protocol == v {
			return true
		}
	}
	return false
}

func (c *ConnectionString) validate() error {
	if c.User == "" {
		return errors.New("Invalid connection string. User cannot be empty")
	}
	if c.Pass == "" {
		return errors.New("Invalid connection string. Password cannot be empty")
	}
	if c.Host == "" {
		return errors.New("Invalid connection string. Host cannot be empty")
	}
	if c.Port < MIN_TCP_PORT || c.Port > MAX_TCP_PORT {
		return errors.New(fmt.Sprintf("Invalid connection string. Port needs to be in range %s-%s", MIN_TCP_PORT, MAX_TCP_PORT))
	}
	if !isSupportedProtocol(c.Protocol) {
		return errors.New(fmt.Sprintf("Invalid connection string protocol. Expected %v, got %s", supportedProtocols))
	}
	return nil
}

func (c *ConnectionString) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/", c.Protocol, c.User, c.Pass, c.Host, c.Port)
}
