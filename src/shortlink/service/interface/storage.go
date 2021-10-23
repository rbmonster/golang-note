package _interface

type storage interface {
	Shorten(url string, exp int64) (string, error)
	ShortLinkInfo(eid string) (interface{}, error)
	UnShorten(eid string) (string, error)
}
