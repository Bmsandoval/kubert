package configs

import (
	"github.com/spf13/viper"
	"reflect"
)

type EkubeConfiguration struct {
	EkubeHost string `json:"EKUBE_HOST"`
	EkubePort string `json:"EKUBE_PORT"`
}

func GetEkubeConfig(vipe viper.Viper) EkubeConfiguration {
	var newEkubeConfiguration EkubeConfiguration
	t := reflect.TypeOf(newEkubeConfiguration)

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("json")

		if tag == "" { continue }
		v := reflect.ValueOf(&newEkubeConfiguration).Elem().FieldByName(field.Name)
		if v.IsValid() {
			tagValue := vipe.GetString(tag)
			v.Set(reflect.ValueOf(tagValue))
		}
	}

	return newEkubeConfiguration
}
