package web

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin/render"
)

// Render type
type Render map[string]*template.Template

var _ render.HTMLRender = Render{}

// OpenRender load tempaltes
func OpenRender(root string, funcs template.FuncMap) (Render, error) {
	rdr := make(Render)
	const ext = ".html"
	// ---------------
	includes, err := filepath.Glob(path.Join(root, "includes", "*"+ext))
	if err != nil {
		return nil, err
	}
	// -----------------
	layouts, err := filepath.Glob(path.Join(root, "layouts", "*"+ext))
	if err != nil {
		return nil, err
	}
	for _, layout := range layouts {
		lyn := filepath.Base(layout)
		lyn = lyn[:len(lyn)-len(ext)]
		lyr := path.Join(root, "views", lyn)

		if err := filepath.Walk(lyr, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || filepath.Ext(info.Name()) != ext {
				return nil
			}

			name := path[len(lyr)+1 : len(path)-len(ext)]
			tpl := template.
				New(name).
				Funcs(funcs)

			files := append(includes, layout, path)
			log.Debugln("find view", name, "with layout", lyn)
			for _, n := range files {
				buf, err := ioutil.ReadFile(n)
				if err != nil {
					return err
				}
				tpl, err = tpl.Parse(string(buf))
				if err != nil {
					return err
				}
			}
			rdr[name] = tpl
			return nil
		}); err != nil {
			return nil, err
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
