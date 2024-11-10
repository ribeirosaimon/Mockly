package server

import (
	"fmt"
	"os"

	tlog "github.com/ribeirosaimon/Mockly/pkg/logs"
	"gopkg.in/yaml.v3"
)

// Environment is a env value
type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type config struct {
	Env  Environment `yaml:"env"`
	Port int         `yaml:"port"`
}

type shortifyEnvironment struct {
	Config config   `yaml:"config"`
	Mongo  DbConfig `yaml:"mongo"`
	Redis  DbConfig `yaml:"redis"`
}

type DbConfig struct {
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	EntryPoint string `yaml:"entryPoint"`
}

var env shortifyEnvironment

func StartEnv(envName Environment) {
	f, err := os.Open(fmt.Sprintf("config.%s.yaml", envName))
	if err != nil {
		tlog.Warn("StartEnv", "Failed to open config file", "err", err)
		// want to stop app, because dont have any environment
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&env); err != nil {
		tlog.Warn("StartEnv", "Failed to parse config file", "err", err)
		panic(err)
	}
	tlog.Info("StartEnv", fmt.Sprintf("Loading environment variables in %s environment in port %d", env.Config.Env, env.Config.Port))
}

func GetEnvironment() config {
	return env.Config
}
func GetMongoConfig() DbConfig {
	return env.Mongo
}
func GetRedisConfig() DbConfig {
	return env.Redis
}
