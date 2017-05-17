package mail

import (
	"net/http"
	"path"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) getReadme(c *h2o.Context) error {
	return c.TEXT(http.StatusOK, path.Join("ops", "vpn", "readme.md"), h2o.H{})
}
