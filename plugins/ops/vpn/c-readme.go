package vpn

import (
	"net/http"
	"path"

	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o"
	"github.com/spf13/viper"
)

func (p *Plugin) getReadme(c *h2o.Context) error {
	data := h2o.H{}
	data["user"] = c.Get(auth.CurrentUser)
	data["name"] = viper.Get("server.name")
	data["home"] = web.Backend()
	data["port"] = 1194
	data["network"] = "10.18.0"

	token, err := p.generateToken(10)
	if err != nil {
		return err
	}
	data["token"] = string(token)
	return c.TEXT(http.StatusOK, path.Join("ops", "vpn", "readme.md"), data)
}
