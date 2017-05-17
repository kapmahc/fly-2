package forms

import (
	"encoding/csv"
	"fmt"

	"github.com/kapmahc/h2o"
)

// getFormExport csv?
func (p *Plugin) getFormExport(c *h2o.Context) error {
	item := c.Get("item").(*Form)
	header, rows, err := p._exportForm(item)
	if err != nil {
		return err
	}
	c.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=form-%d.ini", item.ID))
	c.SetHeader("Content-Type", "text/plain; charset=utf-8")
	wrt := csv.NewWriter(c.Writer)
	wrt.Write(header)

	for _, row := range rows {
		wrt.Write(row)
	}
	wrt.Flush()
	return nil
}
