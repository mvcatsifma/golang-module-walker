// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package config

// Injectors from wiring.go:

func BuildModule() *module {
	configApi := NewApi()
	configModule := NewConfig(configApi)
	return configModule
}
