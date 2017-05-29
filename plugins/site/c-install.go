package site

import (
	"net/http"
	"time"

	"github.com/kapmahc/fly/plugins/auth"
	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

type fmInstall struct {
	Title                string `form:"title" validate:"required"`
	SubTitle             string `form:"subTitle" validate:"required"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Plugin) postInstall(c *h2o.Context) error {
	var fm fmInstall
	if err := c.Bind(&fm); err != nil {
		return err
	}

	lng := c.Get(i18n.LOCALE).(string)
	p.I18n.Set(lng, "site.title", fm.Title)
	p.I18n.Set(lng, "site.subTitle", fm.SubTitle)
	user, err := p.Dao.AddEmailUser("root", fm.Email, fm.Password)
	if err != nil {
		return err
	}
	for _, r := range []string{auth.RoleAdmin, auth.RoleRoot} {
		role, er := p.Dao.Role(r, auth.DefaultResourceType, auth.DefaultResourceID)
		if er != nil {
			return er
		}
		if err = p.Dao.Allow(role.ID, user.ID, 50, 0, 0); err != nil {
			return err
		}
	}
	if err = p.Db.Model(user).UpdateColumn("confirmed_at", time.Now()).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) mustDatabaseEmpty(c *h2o.Context) error {
	lng := c.Get(i18n.LOCALE).(string)
	var count int
	if err := p.Db.Model(&auth.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(http.StatusForbidden, lng, "errors.forbidden")
	}
	return nil
}
