package configs

import (
	"github.com/spf13/viper"
	"reflect"
)

type ServerConfiguration struct {
	SrvPort string `json:"SERVER_PORT"`
	SrvSecretsFile string `json:"SECRETS_FILE"`
}

func GetServerConfig(vipe viper.Viper) ServerConfiguration {
	var newServerConfiguration ServerConfiguration
	t := reflect.TypeOf(newServerConfiguration)

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("json")

		if tag == "" { continue }
		v := reflect.ValueOf(&newServerConfiguration).Elem().FieldByName(field.Name)
		if v.IsValid() {
			tagValue := vipe.GetString(tag)
			v.Set(reflect.ValueOf(tagValue))
		}
	}

	return newServerConfiguration
}
