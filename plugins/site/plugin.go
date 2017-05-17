package site

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/fly/cache"
	"github.com/kapmahc/fly/i18n"
	"github.com/kapmahc/fly/job"
	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/settings"
	"github.com/kapmahc/fly/web"
	"github.com/spf13/viper"
	"golang.org/x/tools/blog/atom"
)

// Plugin plugin
type Plugin struct {
	Dao      *auth.Dao          `inject:""`
	Db       *gorm.DB           `inject:""`
	Jwt      *auth.Jwt          `inject:""`
	I18n     *i18n.I18n         `inject:""`
	Settings *settings.Settings `inject:""`
	Server   *job.Server        `inject:""`
	Cache    *cache.Cache       `inject:""`
	Wrapper  *web.Wrapper       `inject:""`
}

// Init load config
func (p *Plugin) Init() {}

// Dashboard Dashboard
func (p *Plugin) Dashboard(*gin.Context) web.Dropdown {
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

// Workers background task handlers
func (p *Plugin) Workers() map[string]job.Handler {
	return map[string]job.Handler{}
}

func init() {
	pwd, _ := os.Getwd()
	viper.SetDefault("uploader", map[string]interface{}{
		"dir":  path.Join(pwd, "public", "files"),
		"home": "http://localhost/files",
	})
	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   8,
	})

	viper.SetDefault("rabbitmq", map[string]interface{}{
		"user":     "guest",
		"password": "guest",
		"host":     "localhost",
		"port":     "5672",
		"virtual":  "fly-dev",
	})

	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"user":     "postgres",
			"password": "",
			"dbname":   "fly_dev",
			"sslmode":  "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})

	viper.SetDefault("server", map[string]interface{}{
		"port":  8080,
		"ssl":   false,
		"name":  "localhost",
		"theme": "bootstrap",
	})

	viper.SetDefault("secrets", map[string]interface{}{
		"jwt":     web.Random(32),
		"aes":     web.Random(32),
		"hmac":    web.Random(32),
		"csrf":    web.Random(32),
		"session": web.Random(32),
	})

	viper.SetDefault("elasticsearch", map[string]interface{}{
		"host": "localhost",
		"port": 9200,
	})

	// -----------
	web.Register(&Plugin{})
}
