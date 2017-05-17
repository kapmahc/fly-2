package web

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin/render"
)

// Render type
type Render map[string]*template.Template

var _ render.HTMLRender = Render{}

// OpenRender load tempaltes
func OpenRender(r string, m template.FuncMap) (Render, error) {
	rdr := make(Render)
	layouts, err := ioutil.ReadDir(r)
	if err != nil {
		return nil, err
	}

	const layout = "layout.html"
	const ext = ".html"

	for _, lyt := range layouts {
		if lyt.IsDir() {
			root := filepath.Join(r, lyt.Name())
			if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				fn := info.Name()
				if info.IsDir() || filepath.Ext(fn) != ext || fn == layout {
					return nil
				}

				name := path[len(root)+1 : len(path)-len(ext)]
				tpl := template.
					New(name).
					Funcs(m)
				for _, n := range []string{filepath.Join(root, layout), path} {
					buf, err := ioutil.ReadFile(n)
					if err != nil {
						return err
					}
					tpl, err = tpl.Parse(string(buf))
					if err != nil {
						return err
					}
				}

				log.Debugln("find view", name, "with layout", lyt.Name())
				rdr[name] = tpl
				return nil
			}); err != nil {
				return nil, err
			}
		}

	}
	return rdr, nil
}

// Instance supply render string
func (r Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r[name],
		Data:     data,
	}
}
