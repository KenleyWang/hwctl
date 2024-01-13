package config

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	AK     string
	SK     string
	Region string
}

func (c *Config) LoadConfig(configPath string) error {

	err := c.parseConfig(configPath)
	if err != nil {
		return err
	}
	// 在考虑环境变量重载
	c.mergeEnv()

	return nil
}

func (c *Config) MustLoadConfig(configPath string) {
	err := c.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
}

func (c *Config) GenerateCredential() *basic.Credentials {
	return basic.NewCredentialsBuilder().
		WithAk(c.AK).
		WithSk(c.SK).
		Build()
}

func (c *Config) parseConfig(conf string) error {
	data, err := os.ReadFile(conf)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) mergeEnv() {
	assign(reflect.ValueOf(c))
}

func assign(val reflect.Value) {
	v := reflect.Indirect(val)
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("env")
		processOne(v.Field(i), key)
	}
}

func processOne(field reflect.Value, key string) {
	envVal, envOk := os.LookupEnv(key)

	switch field.Type().Kind() {
	case reflect.String:
		if !envOk {
			return
		}

		field.SetString(envVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !envOk {
			return
		}

		val, e := strconv.ParseInt(envVal, 0, field.Type().Bits())
		if e != nil {
			return
		}
		field.SetInt(val)
	case reflect.Float64, reflect.Float32:
		if !envOk {
			return
		}

		val, e := strconv.ParseFloat(envVal, field.Type().Bits())
		if e != nil {
			return
		}
		field.SetFloat(val)
	case reflect.Bool:
		if !envOk {
			return
		}

		val, e := strconv.ParseBool(envVal)
		if e != nil {
			return
		}

		field.SetBool(val)
	case reflect.Struct:
		assign(field)
	}
}
