package web

import (
	"io"
	"net/http"
	"path"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

// Wrapper wrapper
type Wrapper struct {
	Render *render.Render `inject:""`
}

// Text render text template
func (p *Wrapper) Text(w io.Writer, n string, v gin.H) error {
	t, e := template.New("").ParseFiles(path.Join("templates", n))
	if e != nil {
		return e
	}
	return t.Execute(w, v)
}

// Handle wrapper handler
func (p *Wrapper) Handle(f func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if e := f(c); e != nil {
			log.Error(e)
			s := http.StatusInternalServerError
			if he, ok := e.(*HTTPError); ok {
				s = he.Status
			}
			p.Render.JSON(c.Writer, s, e.Error())
		}
	}
}

// Form wrap form handler
func (p *Wrapper) Form(o interface{}, f func(*gin.Context, interface{}) error) gin.HandlerFunc {
	return p.Handle(func(c *gin.Context) error {
		if e := c.Bind(o); e != nil {
			return e
		}
		return f(c, o)
	})
}

// // JSON render json
// func (p *Wrapper) JSON(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
// 	return p.Handle(func(c *gin.Context) error {
// 		v, e := f(c)
// 		if e != nil {
// 			return e
// 		}
// 		p.Render.JSON(c.Writer, http.StatusOK, v)
// 		return nil
// 	})
// }
//
// // XML render xml
// func (p *Wrapper) XML(f func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
// 	return p.Handle(func(c *gin.Context) error {
// 		v, e := f(c)
// 		if e != nil {
// 			return e
// 		}
// 		p.Render.XML(c.Writer, http.StatusOK, v)
// 		return nil
// 	})
// }

// HTML render html
func (p *Wrapper) HTML(t string, f func(*gin.Context, gin.H) error) gin.HandlerFunc {
	return p.Handle(func(c *gin.Context) error {
		v := gin.H{}
		if e := f(c, v); e != nil {
			return e
		}
		v["ctx"] = c.Keys
		p.Render.HTML(c.Writer, http.StatusOK, t, v)
		return nil
	})
}
