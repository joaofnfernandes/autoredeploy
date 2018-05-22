package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/joaofnfernandes/autoredeploy/pkg/mqconfig"
	"github.com/urfave/cli"
)

type ApiServerConfig struct {
	ServerConfig ServerConfig
	MqConfig     mqconfig.MqConfig
}

func ApiServerConfigFromContext(c *cli.Context) ApiServerConfig {
	serverCfg, err := getServerConfig(c)
	if err != nil {
		log.Fatalf("[API] Invalid server config: %s", err)
	}
	mqCfg, err := getMqConfig(c)
	if err != nil {
		log.Fatalf("[API] Invalid message queue config: %s", err)
	}
	return ApiServerConfig{
		serverCfg,
		mqCfg,
	}
}

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

func getMqConfig(c *cli.Context) (mqconfig.MqConfig, error) {
	connectionStr := mqconfig.ConnectionString{
		c.String(mqUserFlagName),
		c.String(mqPasswordFlagName),
		c.String(mqHostFlagName),
		c.Int(mqPortFlagName),
		c.String(mqProtocolFlagName),
	}
	return mqconfig.NewMqConfig(connectionStr)
}
