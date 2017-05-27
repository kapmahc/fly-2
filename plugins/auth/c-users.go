package auth

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/h2o"
	"github.com/kapmahc/h2o/i18n"
)

type fmSignUp struct {
	Name                 string `form:"name" validate:"required,max=255"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Plugin) postUsersSignUp(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	var fm fmSignUp
	if err := c.Bind(&fm); err != nil {
		return err
	}

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return p.I18n.E(http.StatusInternalServerError, l, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.Name, fm.Email, fm.Password)
	if err != nil {
		return err
	}

	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.sign-up"))
	p.sendEmail(l, user, actConfirm)

	return c.JSON(http.StatusOK, h2o.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")})
}

type fmSignIn struct {
	Email      string `form:"email" validate:"required,email"`
	Password   string `form:"password" validate:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Plugin) postUsersSignIn(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	var fm fmSignIn
	if err := c.Bind(&fm); err != nil {
		return err
	}

	user, err := p.Dao.SignIn(l, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		return err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	cm.Set("name", user.Name)
	cm.Set("admin", p.Dao.Is(user.ID, RoleAdmin))
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h2o.H{
		"token": string(tkn),
	})
}

type fmEmail struct {
	Email string `form:"email" validate:"required,email"`
}

func (p *Plugin) getUsersConfirm(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(l, token, actConfirm)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-already-confirm")
	}
	p.Db.Model(user).Update("confirmed_at", time.Now())
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.confirm"))

	c.Redirect(http.StatusFound, p._signInURL())
	return nil
}

func (p *Plugin) postUsersConfirm(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}

	if user.IsConfirm() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-already-confirm")
	}

	p.sendEmail(l, user, actConfirm)

	return c.JSON(http.StatusOK, h2o.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")})
}

func (p *Plugin) getUsersUnlock(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	token := c.Param("token")
	user, err := p.parseToken(l, token, actUnlock)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-not-lock")
	}

	p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil})
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.unlock"))

	c.Redirect(http.StatusFound, p._signInURL())
	return nil
}

func (p *Plugin) postUsersUnlock(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)

	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-not-lock")
	}
	p.sendEmail(l, user, actUnlock)
	return c.JSON(http.StatusOK, h2o.H{"message": p.I18n.T(l, "auth.messages.email-for-unlock")})
}

func (p *Plugin) postUsersForgotPassword(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	var fm fmEmail
	if err := c.Bind(&fm); err != nil {
		return err
	}
	var user *User
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	p.sendEmail(l, user, actResetPassword)

	return c.JSON(http.StatusOK, h2o.H{"message": p.I18n.T(l, "auth.messages.email-for-reset-password")})
}

type fmResetPassword struct {
	Token                string `form:"token" validate:"required"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

func (p *Plugin) postUsersResetPassword(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)

	var fm fmResetPassword
	if err := c.Bind(&fm); err != nil {
		return err
	}
	user, err := p.parseToken(l, fm.Token, actResetPassword)
	if err != nil {
		return err
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.reset-password"))
	return c.JSON(http.StatusOK, h2o.H{"message": p.I18n.T(l, "auth.messages.reset-password-success")})
}

func (p *Plugin) deleteUsersSignOut(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)
	user := c.Get(CurrentUser).(*User)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.sign-out"))
	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) getUsersInfo(c *h2o.Context) error {
	user := c.Get(CurrentUser).(*User)
	return c.JSON(http.StatusOK, h2o.H{"name": user.Name, "email": user.Email})
}

type fmInfo struct {
	Name string `form:"name" validate:"required,max=255"`
	// Home string `form:"home" validate:"max=255"`
	// Logo string `form:"logo" validate:"max=255"`
}

func (p *Plugin) postUsersInfo(c *h2o.Context) error {
	user := c.Get(CurrentUser).(*User)
	var fm fmInfo
	if err := c.Bind(&fm); err != nil {
		return err
	}

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		// "home": fm.Home,
		// "logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, h2o.H{})
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" validate:"required"`
	NewPassword          string `form:"newPassword" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=NewPassword"`
}

func (p *Plugin) postUsersChangePassword(c *h2o.Context) error {
	l := c.Get(i18n.LOCALE).(string)

	user := c.Get(CurrentUser).(*User)
	var fm fmChangePassword
	if err := c.Bind(&fm); err != nil {
		return err
	}
	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.bad-password")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h2o.H{})
}

func (p *Plugin) getUsersLogs(c *h2o.Context) error {
	user := c.Get(CurrentUser).(*User)
	var logs []Log
	if err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&logs).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, logs)
}

func (p *Plugin) indexUsers(c *h2o.Context) error {
	var users []User
	if err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (p *Plugin) _signInURL() string {
	return web.Frontend() + "/users/sign-in"
}
