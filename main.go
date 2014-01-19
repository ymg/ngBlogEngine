package main

import (
	"encoding/json"
	"fmt"
	//"unicode"
	//"runtime"
	//"log"
	//"net/http"

	"github.com/astaxie/beego"
	//"github.com/garyburd/redigo/redis"
	"github.com/russross/blackfriday"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SetHeaders() {
	headers := this.Ctx.ResponseWriter.Header()
	headers.Del("Server")
	headers.Add("Strict-Transport-Security", "max-age=631138519")
	headers.Add("X-Content-Type-Options", "nosniff")
	headers.Add("X-Frame-Options", "DENY")
	headers.Add("X-XSS-Protection", "1; mode=block")
}

//api controllers
type AuthenticationController struct {
	BaseController
}
type MainController struct {
	BaseController
}
type PostController struct {
	BaseController
}

type TestType struct {
	name string
	age  string
}

func (this *PostController) Get() {
	this.SetHeaders()

}
func (this *PostController) Post() {
	this.SetHeaders()

	token := this.GetSession("token")

	if token == nil {
		this.Ctx.Output.SetStatus(403)
	} else {
		this.Ctx.Output.SetStatus(200)
	}
}
func (this *PostController) Delete() {
	this.SetHeaders()
}
func (this *PostController) Put() {
	this.SetHeaders()
}

func (this *AuthenticationController) Get() {
	this.SetHeaders()

	fmt.Printf("\nCSRF token value: %s\n", this.CheckXsrfCookie())

	post := &Post{}
	post.Title = "Angular MicroBlog!"
	post.Body = "Hello from Go 1.1!"
	post.Date = "17-Nov-2013"

	DefaultConfig := &BlogConfig{}
	DefaultConfig.Description = "I write about code!"
	DefaultConfig.Gplus = "#"

	output := blackfriday.MarkdownBasic([]byte("##Hello,World\n-------\n_yes_\n\n    js lol epic"))
	fmt.Println(string(output))

	token := this.GetSession("token").(*TestType)

	if token == nil {
		fmt.Println("no auth token..")
	} else {
		tokenDetails := fmt.Sprintf("token value %v\n", token.name)
		fmt.Println(tokenDetails)
	}

	this.Data["json"] = DefaultConfig
	this.ServeJson()
}
func (this *AuthenticationController) Post() {

	this.SetHeaders()

	var cred map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &cred)
	_, usr := cred["user"]
	_, pass := cred["password"]

	if err == nil && usr && pass {
		fmt.Println(cred["user"])
		fmt.Println(cred["password"])
		this.Ctx.Output.SetStatus(200)
	} else {
		this.Ctx.Output.SetStatus(403)
	}
}

//default controller
func (this *MainController) Get() {
	this.SetHeaders()

	sessionTypeTest := &TestType{"Yasser", "28"}
	this.SetSession("token", sessionTypeTest)
	this.XsrfToken()

	this.TplNames = "index.html"
}

func main() {

	beego.Router("/api/authcheck", &PostController{})
	beego.Router("/api/posts", &PostController{})
	beego.Router("/api/posts/:id", &PostController{})
	beego.Router("/api/auth", &AuthenticationController{})
	beego.Router("/api/*", &MainController{})
	beego.Router("/", &MainController{})
	beego.Router("/*", &MainController{})

	beego.SetStaticPath("/content", "content")
	beego.BeegoServerName = "Microsoft-IIS/8.0"
	beego.ViewsPath = "content"
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.SessionOn = true
	beego.EnableXSRF = true
	beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.XSRFExpire = 3600
	beego.EnableGzip = true
	beego.CopyRequestBody = true

	beego.Run()
}
