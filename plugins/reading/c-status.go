package reading

import (
	"net/http"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

func (p *Plugin) getStatus(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	data := h2o.H{}
	var bc int
	if err := p.Db.Model(&Book{}).Count(&bc).Error; err != nil {
		return err
	}
	data["book"] = h2o.H{
		p.I18n.T(lng, "reading.admin.status.book-count"): bc,
	}

	dict := h2o.H{}
	for _, dic := range dictionaries {
		dict[dic.GetBookName()] = dic.GetWordCount()
	}
	data["dict"] = dict

	return c.JSON(http.StatusOK, data)

}
