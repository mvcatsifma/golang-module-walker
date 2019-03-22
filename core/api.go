package core

type IApi interface {
	IsApi()
}

type Api struct {
	RootLinks []Link
}

func (this *Api) IsApi() {}

func MakeApi(links []Link) Api {
	return Api{
		RootLinks: links,
	}
}

func (this *Api) Get() []Link {
	return this.RootLinks
}
