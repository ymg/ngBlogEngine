// ngBlogEngine project main.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/russross/blackfriday"
)

type AuthenticationController struct {
	beego.Controller
}
type MainController struct {
	beego.Controller
}
type TestType struct {
	name string
	age  string
}

func (this *MainController) Get() {
	SetHeaders(&this.Controller)
	sessionTypeTest := &TestType{"Yasser", "28"}
	this.SetSession("token", sessionTypeTest)
	this.TplNames = "index.html"
}

func (this *AuthenticationController) Get() {
	SetHeaders(&this.Controller)
	//this.Abort("403")
	post := &Post{}
	post.Title = "Angular MicroBlog!"
	post.Body = "Hello from Go 1.1!"
	post.Date = "17-Nov-2013"

	r, _ := json.Marshal(&post)
	c, _ := redis.Dial("tcp", "localhost:6379")
	defer c.Close()
	n, _ := c.Do("APPEND", "post", r)
	fmt.Println(n)

	output := blackfriday.MarkdownBasic([]byte("##Hello,World\n-------\n_yes_\n\n    js lol epic"))
	fmt.Println(string(output))

	this.Data["json"] = post
	this.ServeJson()

	/*token := this.GetSession("token")

	if token == nil {
		this.Ctx.WriteString("no auth token..")
	} else {
		tokenDetails := fmt.Sprintf("token value %v\n", token.(*TestType).name)
		this.Ctx.WriteString(tokenDetails)
	}*/
}

func SetHeaders(this *beego.Controller) {
	headers := this.Ctx.ResponseWriter.Header()
	headers.Del("Server")
	headers.Add("Strict-Transport-Security", "max-age=631138519")
	headers.Add("X-Content-Type-Options", "nosniff")
	headers.Add("X-Frame-Options", "DENY")
	headers.Add("X-XSS-Protection", "1; mode=block")
}

func main() {

	beego.Router("/", &MainController{})
	beego.Router("/auth", &AuthenticationController{})

	beego.SetStaticPath("/content", "content")
	beego.ViewsPath = "content"
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.SessionOn = true

	beego.Run()
}
