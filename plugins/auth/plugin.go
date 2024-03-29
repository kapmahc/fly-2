package auth

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/i18n"
	"github.com/kapmahc/fly/job"
	"github.com/kapmahc/fly/security"
	"github.com/kapmahc/fly/settings"
	"github.com/kapmahc/fly/uploader"
	"github.com/kapmahc/fly/web"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Db       *gorm.DB           `inject:""`
	Jwt      *Jwt               `inject:""`
	Dao      *Dao               `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
	Hmac     *security.Hmac     `inject:""`
	Uploader uploader.Store     `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Dashboard Dashboard
func (p *Plugin) Dashboard(*gin.Context) web.Dropdown {
	return nil
}

// Open open beans
func (p *Plugin) Open(*inject.Graph) error {
	return nil
}

// Atom rss.atom
func (p *Plugin) Atom(lang string) ([]*atom.Entry, error) {
	return []*atom.Entry{}, nil
}

// Sitemap sitemap.xml.gz
func (p *Plugin) Sitemap() ([]stm.URL, error) {
	return []stm.URL{}, nil
}

func init() {
	web.Register(&Plugin{})
}
