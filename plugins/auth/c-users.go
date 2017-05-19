package auth

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/fly/i18n"
	"github.com/kapmahc/fly/web"
)

func (p *Plugin) getUsersSignUp(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "auth.users.sign-up.title")
	return nil
}

type fmSignUp struct {
	Name                 string `form:"name" binding:"required,max=255"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersSignUp(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmSignUp)

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

	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")})
	return nil
}

func (p *Plugin) getUsersSignIn(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "auth.users.sign-in.title")
	return nil
}

type fmSignIn struct {
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

func (p *Plugin) postUsersSignIn(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmSignIn)

	user, err := p.Dao.SignIn(l, fm.Email, fm.Password, c.ClientIP())
	if err != nil {
		return err
	}

	cm := jws.Claims{}
	cm.Set(UID, user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*24*7)
	if err != nil {
		return err
	}
	c.SetCookie(TOKEN, string(tkn), 0, "/", "", web.IsProduction(), true)
	return nil
}

func (p *Plugin) getUsersEmailForm(a string) func(*gin.Context, gin.H) error {
	return func(c *gin.Context, v gin.H) error {
		l := c.MustGet(i18n.LOCALE).(string)
		v["title"] = p.I18n.T(l, "auth.users."+a+".title")
		v["action"] = a
		return nil
	}
}

type fmEmail struct {
	Email string `form:"email" binding:"required,email"`
}

func (p *Plugin) getUsersConfirm(c *gin.Context) error {
	l := c.MustGet(i18n.LOCALE).(string)
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

func (p *Plugin) postUsersConfirm(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}

	if user.IsConfirm() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-already-confirm")
	}

	p.sendEmail(l, user, actConfirm)

	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "auth.messages.email-for-confirm")})
	return nil
}

func (p *Plugin) getUsersUnlock(c *gin.Context) error {
	l := c.MustGet(i18n.LOCALE).(string)
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

func (p *Plugin) postUsersUnlock(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.user-not-lock")
	}
	p.sendEmail(l, user, actUnlock)
	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "auth.messages.email-for-unlock")})
	return nil
}

func (p *Plugin) postUsersForgotPassword(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmEmail)
	var user *User
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	p.sendEmail(l, user, actResetPassword)

	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "auth.messages.email-for-reset-password")})
	return nil
}

func (p *Plugin) getUsersResetPassword(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "auth.users.reset-password.title")
	v["token"] = c.Param("token")
	return nil
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postUsersResetPassword(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)

	fm := o.(*fmResetPassword)
	user, err := p.parseToken(l, fm.Token, actResetPassword)
	if err != nil {
		return err
	}
	p.Db.Model(user).Update("password", p.Hmac.Sum([]byte(fm.Password)))
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.reset-password"))
	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "auth.messages.reset-password-success")})
	return nil
}

func (p *Plugin) deleteUsersSignOut(c *gin.Context) error {
	l := c.MustGet(i18n.LOCALE).(string)
	user := c.MustGet(CurrentUser).(*User)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(l, "auth.logs.sign-out"))
	c.JSON(http.StatusOK, gin.H{})
	return nil
}

type fmInfo struct {
	Name string `form:"name" binding:"required,max=255"`
	Home string `form:"home" binding:"max=255"`
	Logo string `form:"logo" binding:"max=255"`
}

func (p *Plugin) postUsersInfo(c *gin.Context, o interface{}) error {
	user := c.MustGet(CurrentUser).(*User)
	fm := o.(*fmInfo)

	if err := p.Db.Model(user).Updates(map[string]interface{}{
		"home": fm.Home,
		"logo": fm.Logo,
		"name": fm.Name,
	}).Error; err != nil {
		return err
	}
	l := c.MustGet(i18n.LOCALE).(string)
	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "success")})
	return nil
}

type fmChangePassword struct {
	CurrentPassword      string `form:"currentPassword" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Plugin) postUsersChangePassword(c *gin.Context, o interface{}) error {
	l := c.MustGet(i18n.LOCALE).(string)
	fm := o.(*fmChangePassword)
	user := c.MustGet(CurrentUser).(*User)

	if !p.Hmac.Chk([]byte(fm.CurrentPassword), user.Password) {
		return p.I18n.E(http.StatusForbidden, l, "auth.errors.bad-password")
	}
	if err := p.Db.Model(user).
		Update("password", p.Hmac.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": p.I18n.T(l, "success")})
	return nil
}

func (p *Plugin) getUsersSelf(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "auth.users.logs.title")
	user := c.MustGet(CurrentUser).(*User)
	var logs []Log
	if err := p.Db.
		Select([]string{"ip", "message", "created_at"}).
		Where("user_id = ?", user.ID).
		Order("id DESC").Limit(120).
		Find(&logs).Error; err != nil {
		return err
	}

	v["logs"] = logs
	return nil
}

func (p *Plugin) indexUsers(c *gin.Context, v gin.H) error {
	l := c.MustGet(i18n.LOCALE).(string)
	v["title"] = p.I18n.T(l, "auth.users.index.title")
	var users []User
	if err := p.Db.
		Select([]string{"name", "logo", "home"}).
		Order("last_sign_in_at DESC").
		Find(&users).Error; err != nil {
		return err
	}
	v["users"] = users
	return nil
}

func (p *Plugin) _signInURL() string {
	return "/users/sign-in"
}
