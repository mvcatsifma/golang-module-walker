package core

type IModule interface {
	Start() error
	Terminate() error
}

