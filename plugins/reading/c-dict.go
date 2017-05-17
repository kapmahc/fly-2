package reading

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

type fmDict struct {
	Keywords string `form:"keywords" validate:"required,max=255"`
}

func (p *Plugin) postDict(c *h2o.Context) error {

	var fm fmDict
	if err := c.Bind(&fm); err != nil {
		return err
	}
	rst := h2o.H{}
	for _, dic := range dictionaries {
		for _, sen := range dic.Translate(fm.Keywords) {
			var items []h2o.H
			for _, pat := range sen.Parts {
				items = append(items, h2o.H{"type": pat.Type, "body": string(pat.Data)})
			}
			rst[dic.GetBookName()] = items
		}
	}

	return c.JSON(http.StatusOK, rst)
}
