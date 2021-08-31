package configuration

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"reflect"
)

func ParseConfiguration(pathToConf string) Configuration {

	if err := ProcessDotEnv("./", "./conf"); err != nil {
		log.Panic(err)
	}

	c := Configuration{}

	if _, err := toml.DecodeFile(pathToConf, &c); err != nil {
		log.Panicf("Cannot parse configuration: %s", err)
	}

	processEnvVariables(&c)

	return c
}

func processEnvVariables(conf *Configuration) {
	processStruct(reflect.ValueOf(conf).Elem())
}

func processStruct(structValue reflect.Value) {
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		if field.Type.Kind().String() == "struct" {
			processStruct(structValue.Field(i))
			continue
		}

		envKey := field.Tag.Get("env")

		if envKey == "" {
			continue
		}

		envKeyValue := os.Getenv(envKey)

		if envKeyValue == "" {
			log.Panicf("Required env variable [%s] is not set", envKey)
		}

		structValue.Field(i).SetString(envKeyValue)
	}
}
