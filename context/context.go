package context

import (

	"context"

	"github.com/spf13/viper"

)

type Context interface {
	context.Context
	viper.Logger
	ProgramName() string
	DataDirectory() (string, error)
	ConfigFile() (string, error)
}
