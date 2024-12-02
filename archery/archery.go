package archery

import "context"

type Archery struct {
	ctx context.Context

	Cookies map[string]string
}

func NewArchery() *Archery {
	return &Archery{
		Cookies: make(map[string]string),
	}
}

func (a *Archery) SetCookie(key, value string) {
	a.Cookies[key] = value
}

func (a *Archery) GetCookie() map[string]string {
	return a.Cookies
}
