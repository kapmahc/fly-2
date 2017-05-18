package site

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/i18n"
)

func (p *Plugin) getDashboard(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "site.dashboard.title")
	return nil
}
