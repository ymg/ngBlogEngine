package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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

var dynamicScriptUrl string

type TestType struct {
	name, age string
}
type BaseController struct {
	beego.Controller
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
type ConfigurationController struct {
	BaseController
}
type DbConfigurationController struct {
	BaseController
}

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
		fmt.Printf("\nReached Single Post\n")

		var result = &Post{Id: "Random Single Post", Body: `
		<div class="html" style=""><h1>Spatiosi concipit odore</h1><h2>More agros</h2><p>Lorem markdownum nubilus aevo formicas, meae tantorum non quem genetrix sanguis certe mentito. Aquis et miserae collo Achaidas longa. Sine sorores vitam pantherarum opus sollicitare rerum aesculus veterumque vigor pavet sibi <strong>costumque</strong>? Quos arentis, nostri amore. Moneri in forem vero, distentae; et arma ferarum.</p><pre><code>(def ^:dynamic chunk-size 17)

(defn next-chunk [rdr]
  (let [buf (char-array chunk-size)
        s (.read rdr buf)]
  (when (pos? s)
    (java.nio.CharBuffer/wrap buf 0 s))))

(defn chunk-seq [rdr]
  (when-let [chunk (next-chunk rdr)]
    (cons chunk (lazy-seq (chunk-seq rdr)))))</code></pre><p>Auctor invita gerebat solent fuga captus si hederae uterum me mihi mox tangit magnanimus agmen, et crine! Inconsumpta longa deus voce, et nescis tempora cupiunt vocat dare <strong>gratata</strong>: subito.</p><h2>Stratis postquam venit passu</h2><p>Primus esse hoc, te arbor unde alimentaque restat perterrita dictu. Agmen et misit, poscat duxit, <a href="http://omgcatsinspace.tumblr.com/">me</a> moenia, ac humanas undis horas vindex. Alti illi quondam, <em>per quaerit</em> crescendo pariterque mane squamis et doctis et mihi suum pro <a href="http://landyachtz.com/">piasque</a>. Ex quoque quod et haec, derant poposcerat leonum.</p><pre><code>var read = key;
var ictPppoe = marketDirectxThin + boot_netiquette.user(cybercrime, 821339) +
        template_bus_xhtml(zoneDigitalReciprocal, 1);
manet_lossless_parity.virtual_boot_bank = file_pop_margin;
raw_so_opengl += exabyte_troll(rich_video_pram(drive_dvd_mbr,
        workstation_character)) - username_ppm_platform.dongle_folder.storage(
        cdfs_tftp);
var sqlRibbonRaid = leopardSynIcs;
</code></pre><p>Et levat spectasse arcem instar insiluit nostro statque longo habent. Manusque gemitu obstantes in audax et possim de nec fuisset imagine. <a href="http://kimjongunlookingatthings.tumblr.com/">Turbine</a> acta silvas feretro superbus Phoebo perque, harenas inque; munera siquid arbore aequalis blandita solvit! Cum Iovis inemptum Pontum: est heros lignoque in lapis, lux. Me donec admirabile pericla quibus.</p><p>Mirantis frustra aetate. Et des <strong>nec quae nigra</strong> colla missis aliena tenebris desiluit, fluctibus. Quae somnia <em>et longum</em> exercent index, regalemque egreditur deum coloribus Lyncus. Sic mihi: <em>eadem</em> candentem lumina requirit saxa, bis Gradivo cupiens! Modo natis, minus desiluit positi pennas vitiaverit primus inexperrectus potitur eluditque caesae nec nos dederis.</p></div>
		`, Title: "Test Post!!!!", Date: "2014/6/12  12:45PM",
			Markdown: `#Hello post editor
		    
	<script> alert('hello, world); </script>
			
			
##This is a first post edit test
---------------------
			
Lets see if this works`}

		this.Data["json"] = result
		this.ServeJson()

	} else {

		fmt.Printf("\nPost View Controller reached\n")

		var postsSlice []*Post
		var n int

		if this.Ctx.Request.Header.Get("page") != "" {
			n, _ = strconv.Atoi(this.Ctx.Request.Header.Get("page"))
		} else {
			n = 2
		}

		for i := 1; i <= n; i++ {
			postsSlice = append(postsSlice, &Post{Id: fmt.Sprintf("my-first-post-%v", i), Body: `
		<div class="html" style=""><h1>Spatiosi concipit odore</h1><h2>More agros</h2><p>Lorem markdownum nubilus aevo formicas, meae tantorum non quem genetrix sanguis certe mentito. Aquis et miserae collo Achaidas longa. Sine sorores vitam pantherarum opus sollicitare rerum aesculus veterumque vigor pavet sibi <strong>costumque</strong>? Quos arentis, nostri amore. Moneri in forem vero, distentae; et arma ferarum.</p><pre><code>(def ^:dynamic chunk-size 17)

(defn next-chunk [rdr]
  (let [buf (char-array chunk-size)
        s (.read rdr buf)]
  (when (pos? s)
    (java.nio.CharBuffer/wrap buf 0 s))))

(defn chunk-seq [rdr]
  (when-let [chunk (next-chunk rdr)]
    (cons chunk (lazy-seq (chunk-seq rdr)))))</code></pre><p>Auctor invita gerebat solent fuga captus si hederae uterum me mihi mox tangit magnanimus agmen, et crine! Inconsumpta longa deus voce, et nescis tempora cupiunt vocat dare <strong>gratata</strong>: subito.</p><h2>Stratis postquam venit passu</h2><p>Primus esse hoc, te arbor unde alimentaque restat perterrita dictu. Agmen et misit, poscat duxit, <a href="http://omgcatsinspace.tumblr.com/">me</a> moenia, ac humanas undis horas vindex. Alti illi quondam, <em>per quaerit</em> crescendo pariterque mane squamis et doctis et mihi suum pro <a href="http://landyachtz.com/">piasque</a>. Ex quoque quod et haec, derant poposcerat leonum.</p><pre><code>var read = key;
var ictPppoe = marketDirectxThin + boot_netiquette.user(cybercrime, 821339) +
        template_bus_xhtml(zoneDigitalReciprocal, 1);
manet_lossless_parity.virtual_boot_bank = file_pop_margin;
raw_so_opengl += exabyte_troll(rich_video_pram(drive_dvd_mbr,
        workstation_character)) - username_ppm_platform.dongle_folder.storage(
        cdfs_tftp);
var sqlRibbonRaid = leopardSynIcs;
</code></pre><p>Et levat spectasse arcem instar insiluit nostro statque longo habent. Manusque gemitu obstantes in audax et possim de nec fuisset imagine. <a href="http://kimjongunlookingatthings.tumblr.com/">Turbine</a> acta silvas feretro superbus Phoebo perque, harenas inque; munera siquid arbore aequalis blandita solvit! Cum Iovis inemptum Pontum: est heros lignoque in lapis, lux. Me donec admirabile pericla quibus.</p><p>Mirantis frustra aetate. Et des <strong>nec quae nigra</strong> colla missis aliena tenebris desiluit, fluctibus. Quae somnia <em>et longum</em> exercent index, regalemque egreditur deum coloribus Lyncus. Sic mihi: <em>eadem</em> candentem lumina requirit saxa, bis Gradivo cupiens! Modo natis, minus desiluit positi pennas vitiaverit primus inexperrectus potitur eluditque caesae nec nos dederis.</p></div>
		`, Title: fmt.Sprintf("Post # %v", i), Date: "2014/6/12  12:45PM",
				Markdown: `#Hello post editor
		    
	<script> alert('hello, world); </script>
			
			
##This is a first post edit test
---------------------
			
Lets see if this works`})
		}

		//time.Sleep(2 * time.Second)

		this.Data["json"] = postsSlice
		this.ServeJson()
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

		i, _ := p["title"]
		i2, _ := p["markdown"]

		fmt.Println(strings.Replace(strings.ToLower(i.(string)), " ", "-", -1), i2)

		output := blackfriday.MarkdownBasic([]byte(i2.(string)))

		fmt.Println(string(output))

		this.Ctx.Output.SetStatus(200)
	}
}

//authentication = [true]
func (this *PostController) Delete() {
	this.SetHeaders()

	fmt.Println(this.Ctx.Input.Param(":id"))

	this.Ctx.Output.SetStatus(200)
}

//authentication = [true]
func (this *PostController) Put() {
	this.SetHeaders()

	var p map[string]interface{}

	json.Unmarshal(this.Ctx.Input.RequestBody, &p)

	i, _ := p["title"]
	i2, _ := p["markdown"]
	i3, _ := p["Id"]

	if i != nil && i2 != nil && i3 != nil {
		fmt.Printf("Title: %v\n", i)
		fmt.Printf("MD: %v\n", i2)
		fmt.Printf("Id: %v\n", i3)
	}

	this.Ctx.Output.SetStatus(200)
}

func (this *AuthenticationController) Post() {

	this.SetHeaders()

	///////////////////////
	//////// test /////////
	token := this.GetSession("token")

	if token == nil {
		fmt.Println("no auth token..")
	} else {
		tokenDetails := fmt.Sprintf("token value %v\n", token.(*TestType).name)
		fmt.Println(tokenDetails)
	}
	///////////////////////////

	var cred map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &cred)
	usrName, usr := cred["user"]
	usrPass, pass := cred["password"]

	if err == nil && usr && pass && token != nil {
		fmt.Printf("User: %v  |  Password: %v", usrName, usrPass)
		this.Ctx.Output.SetStatus(200)
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

//authentication = [true]
func (this *ConfigurationController) Get() {
	this.SetHeaders()

	fmt.Printf("\nCSRF token value: %s\n", this.CheckXsrfCookie())

	post := &Post{}
	post.Title = "Angular MicroBlog!"
	post.Body = "Hello from Go 1.3!"
	post.Date = "17-Nov-2013"

	DefaultConfig := &BlogConfig{}
	DefaultConfig.Description = "I write about code!"
	DefaultConfig.Gplus = "#"
	DefaultConfig.BlogTitle = "YMG.com"

	output := blackfriday.MarkdownBasic([]byte("##Hello,World\n-------\n_yes_\n\n    js lol epic"))
	fmt.Println(string(output))

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

//default controller
func (this *MainController) Get() {
	this.SetHeaders()

	sessionTypeTest := &TestType{name: "Yasser", age: "28"}
	this.SetSession("token", sessionTypeTest)

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

	beego.Router("/api/posts", &PostController{})
	beego.Router("/api/posts/:id", &PostController{})
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
