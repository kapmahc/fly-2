package forms

import (
	"net/http"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

type fmCancel struct {
	Who string `form:"who" validate:"required,max=255"`
}

func (p *Plugin) postFormCancel(c *h2o.Context) error {
	var fm fmCancel
	if err := c.Bind(&fm); err != nil {
		return err
	}

	lng := c.Get(i18n.LOCALE).(string)
	item := c.Get("item").(*Form)

	if item.Expire() {
		return p.I18n.E(http.StatusForbidden, lng, "forms.errors.expired")
	}
	var record Record
	if err := p.Db.Where("form_id = ? AND (phone = ? OR email = ?)", item.ID, fm.Who, fm.Who).First(&record).Error; err != nil {
		return err
	}

	if err := p.Db.Delete(&record).Error; err != nil {
		return err
	}
	p._sendEmail(lng, item, &record, actCancel)
	return c.JSON(http.StatusOK, h2o.H{})
}
