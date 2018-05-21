package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

type ApiServerConfig struct {
	serverConfig ServerConfig
	mqConfig     MqConfig
}

func ApiServerConfigFromContext(c *cli.Context) ApiServerConfig {
	serverCfg, err := getServerConfig(c)
	if err != nil {
		log.Fatalf("[API] Invalid server config: %s", err)
	}
	mqCfg, err := getMqConfig(c)
	if err != nil {
		log.Fatalf("[API] Invalid connection string: %s", err)
	}
	return ApiServerConfig{
		serverCfg,
		mqCfg,
	}
}

// TODO: validate if configs are valid
type ServerConfig struct {
	network string
	port    int
}

func (c *ServerConfig) String() string {
	return fmt.Sprintf("%s:%d", c.network, c.port)
}

const (
	MIN_TCP_PORT = 0
	MAX_TCP_PORT = 65535
)

func (c *ServerConfig) validate() error {
	if c.port < MIN_TCP_PORT || c.port > MAX_TCP_PORT {
		return errors.New(fmt.Sprintf("Invalid port. Port needs to be in the range %s-%s", MIN_TCP_PORT, MAX_TCP_PORT))
	}
	ip := net.ParseIP(c.network)
	if ip == nil {
		return errors.New(fmt.Sprintf("Invalid network address %s", c.network))
	}
	return nil
}

func getServerConfig(c *cli.Context) (ServerConfig, error) {
	cfg := ServerConfig{
		c.String(apiNetworkFlagName),
		c.Int(apiPortFlagName),
	}
	err := cfg.validate()
	if err != nil {
		return ServerConfig{}, err
	}
	return cfg, nil
}

type MqConnectionString struct {
	user     string
	pass     string
	host     string
	port     int
	protocol string
}

func supportedProtocols() []string {
	const (
		MSGQ = "amqp"
	)
	return []string{MSGQ}
}

func isSupportedProtocol(protocol string) bool {
	for _, v := range supportedProtocols() {
		if protocol == v {
			return true
		}
	}
	return false
}

func (c *MqConnectionString) validate() error {
	if c.user == "" {
		return errors.New("Invalid connection string. User cannot be empty")
	}
	if c.pass == "" {
		return errors.New("Invalid connection string. Password cannot be empty")
	}
	if c.host == "" {
		return errors.New("Invalid connection string. Host cannot be empty")
	}
	if c.port < MIN_TCP_PORT || c.port > MAX_TCP_PORT {
		return errors.New(fmt.Sprintf("Invalid connection string. Port needs to be in range %s-%s", MIN_TCP_PORT, MAX_TCP_PORT))
	}
	if !isSupportedProtocol(c.protocol) {
		return errors.New(fmt.Sprintf("Invalid connection string protocol. Expected %v, got %s", supportedProtocols()))
	}
	return nil
}

func (c *MqConnectionString) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/", c.protocol, c.user, c.pass, c.host, c.port)
}

type MqConfig struct {
	connectionStr MqConnectionString
	queueName     string
}

const defaultQueueName = "webhook"

func getMqConfig(c *cli.Context) (MqConfig, error) {
	connectionStr := MqConnectionString{
		c.String(mqUserFlagName),
		c.String(mqPasswordFlagName),
		c.String(mqHostFlagName),
		c.Int(mqPortFlagName),
		c.String(mqProtocolFlagName),
	}
	err := connectionStr.validate()
	if err != nil {
		return MqConfig{}, err
	}
	return MqConfig{
		connectionStr,
		defaultQueueName,
	}, nil
}
