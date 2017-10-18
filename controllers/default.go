package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
	"strings"
)

type MainController struct {
	beego.Controller
}

//@router / [get]
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//@router /access [get]
func (c *MainController) Access() {
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	token := "lianghuchen"
	listslice := []string{timestamp, nonce, token}

	sort.Strings(listslice)

	list := strings.Join(listslice, "")

	t := sha1.New()
	io.WriteString(t, list)
	xhashcode := fmt.Sprintf("%x", t.Sum(nil))

	hashcode := string(xhashcode)

	if hashcode == signature {
		c.Ctx.WriteString(echostr)
	}
}
