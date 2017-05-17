package site

import (
	"net/http"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

func (p *Plugin) getLocales(c *h2o.Context) error {
	items, err := p.I18n.All(c.Param("lang"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func (p *Plugin) getSiteInfo(c *h2o.Context) error {
	langs, err := p.I18n.Store.Languages()
	if err != nil {
		return err
	}
	lng := c.Get(i18n.LOCALE).(string)
	data := h2o.H{"locale": lng, "languages": langs}
	for _, k := range []string{"title", "subTitle", "keywords", "description", "copyright"} {
		data[k] = p.I18n.T(lng, "site."+k)
	}
	author := h2o.H{}
	for _, k := range []string{"name", "email"} {
		author[k] = p.I18n.T(lng, "site.author."+k)
	}
	data["author"] = author
	return c.JSON(http.StatusOK, data)
}
