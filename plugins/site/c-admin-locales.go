package site

import (
	"net/http"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

func (p *Plugin) getAdminLocales(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	items, err := p.I18n.Store.All(lng)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func (p *Plugin) deleteAdminLocales(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	if err := p.I18n.Store.Del(lng, c.Param("code")); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

type fmLocale struct {
	Code    string `form:"code" validate:"required,max=255"`
	Message string `form:"message" validate:"required"`
}

func (p *Plugin) postAdminLocales(c *h2o.Context) error {
	var fm fmLocale
	if err := c.Bind(&fm); err != nil {
		return err
	}
	lng := c.Get(i18n.LOCALE).(string)
	if err := p.I18n.Set(lng, fm.Code, fm.Message); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}
