package core

type IApi interface {
	GetLinks() []Link
}

type LinkAdder interface {
	Add(link Link)
}

type Api struct {
	Links []Link
}

func MakeApi(links []Link) Api {
	return Api{
		Links: links,
	}
}

func (this *Api) GetLinks() []Link {
	return this.Links
}
