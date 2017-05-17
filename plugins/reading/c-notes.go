package reading

import (
	"net/http"

	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

func (p *Plugin) myNotes(c *h2o.Context) error {

	user := c.Get(auth.CurrentUser).(*auth.User)
	isa := c.Get(auth.IsAdmin).(bool)
	var notes []Note
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&notes).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, notes)
}

func (p *Plugin) indexNotes(c *h2o.Context) error {

	var total int64
	var pag *web.Pagination
	if err := p.Db.Model(&Note{}).Count(&total).Error; err != nil {
		return err
	}

	pag = web.NewPagination(c.Request, total)
	var notes []Note
	if err := p.Db.
		Limit(pag.Limit()).Offset(pag.Offset()).
		Find(&notes).Error; err != nil {
		return err
	}

	for _, it := range notes {
		pag.Items = append(pag.Items, it)
	}

	return c.JSON(http.StatusOK, pag)
}

type fmNoteNew struct {
	Type   string `form:"type" validate:"required,max=8"`
	Body   string `form:"body" validate:"required,max=2000"`
	BookID uint   `form:"bookId"`
}

func (p *Plugin) createNote(c *h2o.Context) error {

	user := c.Get(auth.CurrentUser).(*auth.User)

	var fm fmNoteNew
	if err := c.Bind(&fm); err != nil {
		return err
	}
	item := Note{
		Type:   fm.Type,
		Body:   fm.Body,
		BookID: fm.BookID,
		UserID: user.ID,
	}
	if err := p.Db.Create(&item).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func (p *Plugin) showNote(c *h2o.Context) error {
	var item Note
	if err := p.Db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}

type fmNoteEdit struct {
	Type string `form:"type" validate:"required,max=8"`
	Body string `form:"body" validate:"required,max=2000"`
}

func (p *Plugin) updateNote(c *h2o.Context) error {
	note := c.Get("item").(*Note)

	var fm fmNoteEdit
	if err := c.Bind(&fm); err != nil {
		return err
	}

	if err := p.Db.Model(note).
		Updates(map[string]interface{}{
			"body": fm.Body,
			"type": fm.Type,
		}).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) destroyNote(c *h2o.Context) error {
	n := c.Get("item").(*Note)
	if err := p.Db.Delete(n).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) canEditNote(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	user := c.Get(auth.CurrentUser).(*auth.User)

	var n Note
	if err := p.Db.Where("id = ?", c.Param("id")).First(&n).Error; err != nil {
		return err
	}
	if user.ID == n.UserID || c.Get(auth.IsAdmin).(bool) {
		c.Set("item", &n)
		return nil
	}
	return p.I18n.E(http.StatusForbidden, lng, "errors.forbidden")
}
