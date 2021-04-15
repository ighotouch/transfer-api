package bank

type ConfigDatabase struct {
	AppName string `env:"APP_NAME" env-default:"BANK"`
	Port    string `env:"MY_APP_PORT" env-default:"5000"`
}

var cfg ConfigDatabase
