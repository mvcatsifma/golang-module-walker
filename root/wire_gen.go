// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package root

// Injectors from wiring.go:

func BuildModule() *Root {
	rootApi := NewApi()
	root := NewRoot(rootApi)
	return root
}
