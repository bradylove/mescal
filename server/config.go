package main

type Config struct {
	Port string
}

func NewConfig() Config {
	return Config{
		Port: "8899",
	}
}
