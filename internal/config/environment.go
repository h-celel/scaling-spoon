package config

import "github.com/h-celel/mapenv"

type Environment struct {
	PostgresUser     string `mpe:"POSTGRES_USER"`
	PostgresPassword string `mpe:"POSTGRES_PASSWORD"`
	PostgresHost     string `mpe:"POSTGRES_HOST"`
	PostgresPort     uint   `mpe:"POSTGRES_PORT"`
	PostgresDB       string `mpe:"POSTGRES_DB"`
	PostgresSslMode  string `mpe:"POSTGRES_SSL_MODE"`
	RabbitmqHost     string `mpe:"RABBITMQ_HOST"`
	RabbitmqUser     string `mpe:"RABBITMQ_USER"`
	RabbitmqPassword string `mpe:"RABBITMQ_PASSWORD"`
	HealthcheckPort  uint   `mpe:"HEALTHCHECK_PORT"`
	GRPCHost         string `mpe:"GRPC_HOST"`
	DBSchemaURL      string `mpe:"DB_SCHEMA_URL"`
}

func NewEnvironment() *Environment {
	env := &Environment{
		PostgresUser:     "postgres",
		PostgresPassword: "password",
		PostgresDB:       "postgres",
		PostgresHost:     "localhost",
		PostgresPort:     5432,
		PostgresSslMode:  "disable",
		HealthcheckPort:  DefaultHealthcheckPort,
		GRPCHost:         DefaultGRPCHost,
		DBSchemaURL:      DefaultDBSchemaURL,
	}
	err := mapenv.Decode(env)
	if err != nil {
		panic(err)
	}
	return env
}
