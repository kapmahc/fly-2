package auth

import (
	"net/http"

	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

type fmAttachmentNew struct {
	Type string `form:"type" validate:"required,max=255"`
	ID   uint   `form:"id"`
}

func (p *Plugin) createAttachment(c *h2o.Context) error {
	user := c.Get(CurrentUser).(*User)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return err
	}

	url, size, err := p.Uploader.Save(header)
	if err != nil {
		return err
	}

	// http://golang.org/pkg/net/http/#DetectContentType
	buf := make([]byte, 512)
	file.Seek(0, 0)
	if _, err = file.Read(buf); err != nil {
		return err
	}

	a := Attachment{
		Title:        header.Filename,
		URL:          url,
		UserID:       user.ID,
		MediaType:    http.DetectContentType(buf),
		Length:       size / 1024,
		ResourceType: DefaultResourceType, //fm.Type,
		ResourceID:   DefaultResourceID,   //fm.ID,
	}
	if err := p.Db.Create(&a).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

type fmAttachmentEdit struct {
	Title string `form:"title" validate:"required,max=255"`
}

func (p *Plugin) updateAttachment(c *h2o.Context) error {
	a := c.Get("item").(*Attachment)
	var fm fmAttachmentEdit
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if err := p.Db.Model(a).Update("title", fm.Title).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, a)
}

func (p *Plugin) destroyAttachment(c *h2o.Context) error {
	a := c.Get("item").(*Attachment)
	if err := p.Db.Delete(a).Error; err != nil {
		return err
	}
	if err := p.Uploader.Remove(a.URL); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, a)
}

func (p *Plugin) showAttachment(c *h2o.Context) error {
	a := c.Get("item").(*Attachment)
	return c.JSON(http.StatusOK, a)
}

func (p *Plugin) indexAttachments(c *h2o.Context) error {
	user := c.Get(CurrentUser).(*User)
	isa := c.Get(IsAdmin).(bool)
	var items []Attachment
	qry := p.Db
	if !isa {
		qry = qry.Where("user_id = ?", user.ID)
	}
	if err := qry.Order("updated_at DESC").Find(&items).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, items)
}

func (p *Plugin) canEditAttachment(c *h2o.Context) error {
	user := c.Get(CurrentUser).(*User)
	lng := c.Get(i18n.LOCALE).(string)

	var a Attachment
	if err := p.Db.Where("id = ?", c.Param("id")).First(&a).Error; err != nil {
		return err
	}

	if user.ID == a.UserID || c.Get(IsAdmin).(bool) {
		c.Set("item", &a)
		return nil
	}

	return p.I18n.E(http.StatusForbidden, lng, "auth.errors.not-allow")
}
