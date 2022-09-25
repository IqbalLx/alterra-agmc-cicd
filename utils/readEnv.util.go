package utils

import (
	"fmt"
	"os"
)

type Default struct {
	Value string
}

type Env interface {
	Read(string) string
}

type env struct {
	EnvFile string
}

func NewEnv(envFile string) *env {
	return &env{envFile}
}

func (e *env) Read(envName string) string {
	return e.ReadWithDefaultVal(envName, Default{Value: ""})
}

func (e *env) ReadWithDefaultVal(envName string, defaultValue Default) string {
	value, isPresent := os.LookupEnv(envName)
	if !isPresent {
		if defaultValue.Value != "" {
			return defaultValue.Value
		}

		msg := fmt.Errorf("%s is not found in %s, please update", envName, e.EnvFile)
		panic(msg)
	}

	return value
}
