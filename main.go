package main

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	//"strconv"
	"strings"
	"time"
	//"io/ioutil"
	//"unicode"
	//"runtime"
	//"net/http"

	"github.com/astaxie/beego"
	"github.com/howeyc/fsnotify"
	"github.com/russross/blackfriday"
	//"github.com/garyburd/redigo/redis"
)

const oneYear = time.Hour * 24 * 365

var authenticationDAO AuthDAO
var blogDAO BlogDAO
var hashUtils Hasher
var dynamicScriptUrl string

type TestType struct {
	name, age string
}
type BaseController struct {
	beego.Controller
}

//api controllers
type AuthenticationController struct{ BaseController }
type AdminController struct{ BaseController }
type MainController struct{ BaseController }
type PostController struct{ BaseController }
type ConfigurationController struct{ BaseController }
type DbConfigurationController struct{ BaseController }

func (this *BaseController) SetHeaders() {
	headers := this.Ctx.ResponseWriter.Header()
	//headers.Del("Server")
	headers.Add("Strict-Transport-Security", "max-age=631138519; includeSubDomains")
	headers.Add("X-Content-Type-Options", "nosniff")
	headers.Add("X-Frame-Options", "SAMEORIGIN")
	headers.Add("X-XSS-Protection", "1; mode=block")
}

func AuthCheck(c *BaseController) {

}

//authentication = [optional]
func (this *PostController) Get() {
	this.SetHeaders()

	/*ctxHeader := this.Ctx.ResponseWriter.Header()

	ctxHeader.Add("Content-Type", "application/x-javascript")
	ctxHeader.Add("Expires", time.Now().Add(oneYear).String())

	this.Ctx.WriteString(`(function x(){
	return function(){
	console.log('Hello, Script handler!');
	}})()();`)*/

	if this.Ctx.Input.Param(":id") != "" {

		post, err := blogDAO.Get(&Post{Id: this.Ctx.Input.Param(":id")})

		if err == nil {
			this.Data["json"] = post
			this.ServeJson()
		}

	} else {

		if this.Ctx.Request.Header.Get("page") != "" {
			n, _ := strconv.Atoi(this.Ctx.Request.Header.Get("page"))

			post, err := blogDAO.GetAll(n)

			if err != nil {
				fmt.Println(err)
				this.Ctx.Output.SetStatus(500)
			}

			this.Data["json"] = post
			this.ServeJson()

		} else {

			post, err := blogDAO.GetAll(1)

			if err != nil {
				fmt.Println(err)
				this.Ctx.Output.SetStatus(500)
			}

			this.Data["json"] = post
			this.ServeJson()
		}

	}

}

//authentication = [true]
func (this *PostController) Post() {
	this.SetHeaders()
	token := this.GetSession("token")

	if token == nil {
		this.Ctx.Output.SetStatus(401)
	} else {

		var p map[string]interface{}

		json.Unmarshal(this.Ctx.Input.RequestBody, &p)

		title, _ := p["title"]
		md, _ := p["markdown"]

		currentTime := time.Now()

		newP := new(Post)

		newP.Id = strings.Replace(strings.ToLower(title.(string)), " ", "-", -1)
		newP.Body = string(blackfriday.MarkdownBasic([]byte(md.(string))))
		newP.Markdown = md.(string)
		newP.Date = fmt.Sprintf("%v/%v/%v - %v:%v", currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
		newP.Title = title.(string)

		blogDAO.NewPost(newP)

		this.Ctx.Output.SetStatus(200)
	}
}

//authentication = [true]
func (this *PostController) Delete() {
	this.SetHeaders()

	if err := blogDAO.DeletePost(&Post{Id: this.Ctx.Input.Param(":id")}); err != nil {
		this.Ctx.Output.SetStatus(500)
	} else {
		this.Ctx.Output.SetStatus(200)
	}

}

//authentication = [true]
func (this *PostController) Put() {
	this.SetHeaders()

	var p map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &p)

	title, _ := p["title"]
	md, _ := p["markdown"]
	id, _ := p["Id"]

	currentTime := time.Now()

	updateP := new(Post)

	updateP.Id = id.(string)
	updateP.Date = fmt.Sprintf("%v/%v/%v - %v:%v", currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
	updateP.Body = string(blackfriday.MarkdownBasic([]byte(md.(string))))
	updateP.Markdown = md.(string)
	updateP.Title = title.(string)

	blogDAO.UpdatePost(updateP)

	this.Ctx.Output.SetStatus(200)
}

func (this *AuthenticationController) Post() {

	this.SetHeaders()

	///////////////////////
	//////// test /////////
	/*token := this.GetSession("token")

	if token == nil {
		fmt.Println("no auth token..")
	} else {
		tokenDetails := fmt.Sprintf("token value %v\n", token.(*TestType).name)
		fmt.Println(tokenDetails)
	}*/
	///////////////////////////

	var cred map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &cred)
	usrName, usr := cred["user"]
	usrPass, pass := cred["password"]

	if err == nil && usr && pass {

		fmt.Println("Authenticating Now...")

		credentials := struct {
			Username string
			Password string
		}{
			usrName.(string),
			usrPass.(string),
		}

		fmt.Println("Comparing with db data now...")

		if err := authenticationDAO.AuthenticateUser(credentials); err == nil {
			//adding auth token
			sessionTypeTest := &TestType{name: "Yasser", age: "28"}
			fmt.Println("Successfully Authenticated User...")

			this.SetSession("token", sessionTypeTest)

			//fmt.Println("Successfully Authenticated User...")
			this.Ctx.Output.SetStatus(200)

		} else {
			this.Ctx.Output.SetStatus(500)
		}

	} else {
		this.Ctx.Output.SetStatus(401)
	}
}
func (this *AuthenticationController) Delete() {
	this.SetHeaders()

	this.DestroySession()

	this.Ctx.Output.SetStatus(200)

}
func (this *AuthenticationController) Put() {
	this.SetHeaders()
	token := this.GetSession("token")

	if token != nil {
		this.Ctx.Output.SetStatus(200)
		fmt.Printf("Token details : %v", token.(*TestType).name)
	} else {
		/* 403 == lack of sufficient privileges / roles */
		this.Ctx.Output.SetStatus(401)
	}
}

// authentication = [true]
func (this *AdminController) Put() {
	this.SetHeaders()

	token := this.GetSession("token")

	if token != nil {

		var cred map[string]interface{}
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &cred)

		currentPassword, cPass := cred["currentpass"]
		newPassword, newPass := cred["newpass"]
		_, confNewPass := cred["confirmnewpass"]

		if err == nil && cPass && newPass && confNewPass && newPass == confNewPass {
			Pass := struct {
				NewPassword     string
				CurrentPassword string
			}{
				newPassword.(string),
				currentPassword.(string),
			}

			if err := authenticationDAO.EditAdminPassword(Pass); err == nil {
				this.Ctx.Output.SetStatus(200)
			} else {
				this.Ctx.Output.SetStatus(403)
			}
		}
	} else {
		/* 403 == lack of sufficient privileges / roles */
		this.Ctx.Output.SetStatus(401)
	}
}

//authentication = [true]
func (this *ConfigurationController) Get() {
	this.SetHeaders()

	fmt.Printf("\nCSRF token value: %s\n", this.CheckXsrfCookie())

	DefaultConfig := &BlogConfig{}
	DefaultConfig.Description = "I write about code!"
	DefaultConfig.Gplus = "#"
	DefaultConfig.BlogTitle = "YMG's Random Thoughts!"

	this.Data["json"] = DefaultConfig
	this.ServeJson()
}

//authentication = [true]
func (this *ConfigurationController) Put() {}

//authentication = [true]
func (this *DbConfigurationController) Get() {

	this.SetHeaders()

	InitDbConfig()

	this.Data["json"] = GlobalCfg

	this.ServeJson()

}

//authentication = [true]
func (this *DbConfigurationController) Put() {

	this.SetHeaders()

	this.Data["json"] = GlobalCfg

	this.ServeJson()

}

//default controller
func (this *MainController) Get() {
	this.SetHeaders()

	this.TplNames = "index.html"
}

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Process events
	go func() {
		var s int = 0
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev, ev)
				s = s + 1
				dynamicScriptUrl = fmt.Sprintf("generating new url, %v", s)
				log.Println(dynamicScriptUrl)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("minified[js]&[css]")
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	/////
	config := flag.String("c", "", "path of json configuration file - default: local directory")

	flag.Parse()

	if len(*config) > 0 {
		InitDbConfigWithPath(*config)
	} else {
		InitDbConfig()
	}

	blogDAO.InitBlogDao(GlobalCfg)
	authenticationDAO.InitAuthDao(GlobalCfg)

	/////

	/*beego.RunMode = "prod"
	beego.HttpsPort = 443
	beego.EnableHttpTLS = true
	beego.HttpCertFile = "cert.pem"
	beego.HttpKeyFile = "key.pem"

	go func() {

		if err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			//sslURL := fmt.Sprintf("https://%s%s", req.URL.Host, req.URL.RequestURI)
			http.Redirect(w, req, "https://ymg.io", http.StatusMovedPermanently)

		})); err != nil {
			log.Fatal(err)
		}

	}()*/

	beego.Router("/api/posts", &PostController{})
	beego.Router("/api/posts/:id", &PostController{})
	beego.Router("/api/user/update", &AdminController{})
	beego.Router("/api/login", &AuthenticationController{})
	beego.Router("/api/logout", &AuthenticationController{})
	beego.Router("/api/config", &ConfigurationController{})
	beego.Router("/api/dbconfig", &DbConfigurationController{})
	beego.Router("/api/*", &MainController{})
	beego.Router("/", &MainController{})
	beego.Router("/*", &MainController{})

	beego.SetStaticPath("/content", "content")
	beego.RunMode = "test"
	beego.BeegoServerName = "Microsoft-IIS/8.0"
	beego.ViewsPath = "content"
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.MaxMemory = 5 << 20
	beego.SessionOn = true
	beego.SessionName = "ymg.account"
	//beego.SessionHashKey = rndString(KEYLENGTH)
	beego.EnableXSRF = true
	beego.XSRFKEY = rndString(KEYLENGTH)
	beego.XSRFExpire = 60 * 60 * 24 // 60 sec * 60 min * 24 hours
	beego.EnableGzip = true
	beego.CopyRequestBody = true

	beego.Run()
}

func rndString(n int) string {
	const alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789~!@#$%^&{}[]:,.*()_+-/?><."
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i := 0; i < n; i++ {
		bytes[i] = alpha[bytes[i]%byte(len(alpha))]
	}
	return string(bytes)
}
