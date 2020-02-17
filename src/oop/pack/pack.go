package pack

import "fmt"

type Lang struct {
	name    string
	website string
}

func (l *Lang) GetName() string {
	return l.name
}

func (l *Lang) SetName(name string) {
	l.name = name
}

func (l *Lang) GetWebsite() string {
	return l.website
}

func (l *Lang) SetWebsite(website string) {
	l.website = website
}

func (l *Lang) ToString() string {
	return fmt.Sprintf("Lang:[name = %s,website = %s]", l.name, l.website)
}
