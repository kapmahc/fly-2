package auth

import (
	"github.com/facebookgo/inject"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o/i18n"
	"github.com/kapmahc/h2o/job"
	"github.com/kapmahc/h2o/security"
	"github.com/kapmahc/h2o/settings"
	"github.com/kapmahc/h2o/uploader"
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
	return []stm.URL{
		{"loc": "/users/sign-in"},
		{"loc": "/users/sign-up"},
		{"loc": "/users/forgot-password"},
		{"loc": "/users/confirm"},
		{"loc": "/users/unlock"},
		{"loc": "/users"},
	}, nil
}

func init() {
	web.Register(&Plugin{})
}
