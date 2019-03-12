package core

type Visitor interface {
	Visit(module IModule) (w Visitor)
}

type inspector func(IModule) bool

func (f inspector) Visit(module IModule) Visitor {
	if f(module) {
		return f
	}
	return nil
}

func Walk(v Visitor, module IModule) {
	if v = v.Visit(module); v == nil {
		return
	}

	for _, m := range module.GetChildren() {
		Walk(v, m)
	}

	v.Visit(nil)
}

func Inspect(module IModule, f func(IModule) bool) {
	Walk(inspector(f), module)
}
