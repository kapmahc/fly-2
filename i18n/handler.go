package i18n

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	// LOCALE locale key
	LOCALE = "locale"
)

// Middleware detect language from http request
func (p *I18n) Middleware() (gin.HandlerFunc, error) {
	langs, err := p.Store.Languages()
	if err != nil {
		return nil, err
	}
	var tags []language.Tag
	for _, l := range langs {
		tags = append(tags, language.Make(l))
	}
	matcher := language.NewMatcher(tags)

	return func(c *gin.Context) {
		lang, written := p.detect(c.Request)
		tag, _, _ := matcher.Match(language.Make(lang))
		if tag.String() != lang {
			written = true
		}
		if written {
			c.SetCookie(LOCALE, tag.String(), 1<<32-1, "/", "", false, false)
		}
		c.Set(LOCALE, tag.String())
	}, nil
}

func (p *I18n) detect(r *http.Request) (lang string, written bool) {
	written = true
	// 1. Check URL arguments.
	if lang = r.URL.Query().Get(LOCALE); lang != "" {
		return
	}

	// 2. Get language information from cookies.
	if ck, er := r.Cookie(LOCALE); er == nil {
		written = false
		lang = ck.Value
		return
	}

	// 3. Get language information from 'Accept-Language'.
	if al := r.Header.Get("Accept-Language"); len(al) > 4 {
		lang = al[:5]
		return // Only compare first 5 letters.
	}

	return
}
