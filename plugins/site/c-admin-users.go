package site

import (
	"net/http"

	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/h2o"
)

func (p *Plugin) indexAdminUsers(c *h2o.Context) error {
	var items []auth.User
	if err := p.Db.
		Order("last_sign_in_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}
