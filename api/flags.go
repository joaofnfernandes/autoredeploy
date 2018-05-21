package main

import (
	"github.com/urfave/cli"
)

var (
	apiNetworkFlagName = "api-network"
	apiNetworkFlag     = cli.StringFlag{
		Name:   apiNetworkFlagName,
		EnvVar: "API_NETWORK",
		Value:  "0.0.0.0",
		Usage:  "Network where API serve listens",
	}
	apiPortFlagName = "api-port"
	apiPortFlag     = cli.IntFlag{
		Name:   apiPortFlagName,
		EnvVar: "API_PORT",
		Value:  8000,
		Usage:  "Port where API serve listens",
	}
	mqUserFlagName = "mq-user"
	mqUserFlag     = cli.StringFlag{
		Name:   mqUserFlagName,
		EnvVar: "MQ_USER",
		Value:  "admin",
		Usage:  "Username to use when connecting to the message queue",
	}
	mqPasswordFlagName = "mq-password"
	mqPasswordFlag     = cli.StringFlag{
		Name:   mqPasswordFlagName,
		EnvVar: "MQ_PASSWORD",
		Value:  "password",
		Usage:  "Password to use when connecting to the message queue",
	}
	mqHostFlagName = "mq-host"
	mqHostFlag     = cli.StringFlag{
		Name:   mqHostFlagName,
		EnvVar: "MQ_HOST",
		Value:  "msgq",
		Usage:  "Host to connect for message queue service",
	}
	mqPortFlagName = "mq-port"
	mqPortFlag     = cli.IntFlag{
		Name:   mqPortFlagName,
		EnvVar: "MQ_PORT",
		Value:  5672,
		Usage:  "Port to use when connecting to message queue",
	}
	mqProtocolFlagName = "mq-protocol"
	mqProtocolFlag     = cli.StringFlag{
		Name:   mqProtocolFlagName,
		EnvVar: "MQ_PROTOCOL",
		Value:  "msgq",
		Usage:  "Protocol to use when connecting to message queue",
	}
	flags = []cli.Flag{
		apiNetworkFlag,
		apiPortFlag,
		mqUserFlag,
		mqPasswordFlag,
		mqHostFlag,
		mqPortFlag,
		mqProtocolFlag,
	}
)
