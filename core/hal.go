package core

type Link map[string]string

type ILinkAdder interface {
	Add(link Link)
}

type ILinkGetter interface {
	Get() []Link
}
