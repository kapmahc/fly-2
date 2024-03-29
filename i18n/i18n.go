package i18n

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"github.com/kapmahc/fly/cache"
	"github.com/kapmahc/fly/web"
	"golang.org/x/text/language"
)

// I18n i18n
type I18n struct {
	Store Store        `inject:""`
	Cache *cache.Cache `inject:""`
}

// F format message
func (p *I18n) F(lang, code string, obj interface{}) (string, error) {
	msg, err := p.Store.Get(lang, code)
	if err != nil {
		return "", err
	}
	tpl, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, obj)
	return buf.String(), err
}

//E create an i18n error
func (p *I18n) E(status int, lang, code string, args ...interface{}) error {
	if status == 0 {
		status = http.StatusInternalServerError
	}
	return &web.HTTPError{
		Message: p.T(lang, code, args...),
		Status:  status,
	}
}

//T translate by lang tag
func (p *I18n) T(lang, code string, args ...interface{}) string {
	msg := p.get(lang, code)
	if msg == "" {
		return code
	}
	return fmt.Sprintf(msg, args...)
}

// All all items
func (p *I18n) All(lang string) (map[string]interface{}, error) {
	rt := make(map[string]interface{})

	items, err := p.Store.All(lang)
	if err != nil {
		return nil, err
	}
	for k, v := range items {
		codes := strings.Split(k, ".")
		tmp := rt
		for i, c := range codes {
			if i+1 == len(codes) {
				tmp[c] = v
			} else {
				if tmp[c] == nil {
					tmp[c] = make(map[string]interface{})
				}
				tmp = tmp[c].(map[string]interface{})
			}
		}

	}
	return rt, nil
}

// Load sync records
func (p *I18n) Load(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		const ext = ".ini"
		name := info.Name()
		if info.Mode().IsRegular() && filepath.Ext(name) == ext {
			log.Debugf("Find locale file %s", path)
			if err != nil {
				return err
			}
			lang := name[0 : len(name)-len(ext)]
			if _, err := language.Parse(lang); err != nil {
				return err
			}
			cfg, err := ini.Load(path)
			if err != nil {
				return err
			}

			items := cfg.Section(ini.DEFAULT_SECTION).KeysHash()
			for k, v := range items {
				if err := p.Store.Set(lang, k, v, false); err != nil {
					return err
				}
			}
			log.Infof("find %d items", len(items))
		}
		return nil
	})
}

// Set update locale
func (p *I18n) Set(lang, code, message string) error {
	key := p.key(lang, code)
	if err := p.Store.Set(lang, code, message, true); err != nil {
		return err
	}
	return p.Cache.Set(key, message, defaultTTL)
}

func (p *I18n) key(lang, code string) string {
	return lang + "://locales/" + code
}

func (p *I18n) get(lang, code string) string {
	key := p.key(lang, code)
	var msg string
	err := p.Cache.Get(key, &msg)
	if err == nil {
		return msg
	}
	msg, err = p.Store.Get(lang, code)
	if err == nil {
		p.Cache.Set(key, msg, defaultTTL)
		return msg
	}
	return ""
}

const (
	defaultTTL = time.Hour * 24 * 30
)
