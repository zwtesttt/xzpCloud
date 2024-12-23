package domain

type ActionUrl struct {
	sid  string
	name string
	url  string
}

func NewActionUrl(sid, name, url string) *ActionUrl {
	return &ActionUrl{
		sid:  sid,
		name: name,
		url:  url,
	}
}

func (a *ActionUrl) Sid() string {
	return a.sid
}

func (a *ActionUrl) Name() string {
	return a.name
}

func (a *ActionUrl) Url() string {
	return a.url
}
