package model

type URL struct {
	Original  string
	Shortened string
}

func NewUrl(original, shortened string) URL {
	return URL{
		Original:  original,
		Shortened: shortened,
	}
}
