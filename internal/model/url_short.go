package model

type UrlShortDataBase struct {
	Name     string
	UrlDict  map[string]string
	FilePath string
}

func (d *UrlShortDataBase) Add(original, short string) {
	d.UrlDict[short] = original
}

func (d *UrlShortDataBase) GetUrl(short string) (string, bool) {
	val, has := d.UrlDict[short]
	if !has {
		return "", has
	}
	return val, has
}
