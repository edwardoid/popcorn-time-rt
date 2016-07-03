package pp

import (
	"www"
	"log"
	"fmt"
)

type Tracker struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Cookies map[string]string `json:"cookies"`
}

func (this *Tracker) Fetch() []Media {
	return nil
}

func (this *Tracker) Auth() bool {
	form := map[string]string{}
	form["login_username"] = "edwardoid"
	form["login_password"] = "azazello"
	form["login"] = "%E2%F5%EE%E4"
	cxt, err := www.PostForm("http://rutracker.org/forum/login.php", form)
	if (err != nil) {
		log.Fatal("Can't authenticate to tracker website")
		return false;
	}

	this.Cookies = cxt.Cookies

	return false
}

func StartTracker(name string) {
	var t Tracker = GetConfiguration().Trackers[name]

	t.Auth()
	go t.Search("1080p")
}

type Forum struct {
	Name string
	Id int
	Subforums []Forum
}

func (this *Forum) AddSubForum(sub Forum) {
	if this.Subforums == nil {
		this.Subforums = []Forum{sub}
		return
	}

	this.Subforums = append(this.Subforums, sub)
}

func (this *Tracker) GetForumsList() ([]Forum, error) {
	c := www.Context{}
	c.URL = "http://rutracker.org/forum/index.php?map=1"
	c.Method = "POST"
	c.Cookies = this.Cookies
	r, e := www.Do(&c)
	if e != nil {
		log.Fatal(fmt.Print("Can't fetch forums list"))
		return nil, e
	}

	mapXml, mapErr := www.SaveResponse(r)
	if mapErr != nil {
		log.Fatal("Can't save response html")
		return nil, mapErr
	}

	root, rootE := www.ParseAndGet(mapXml, "//*[@id=\"f-map\"]/ul")
	if rootE != nil || root == nil {
		log.Fatal("Can't parse response xml")
		return nil, rootE
	}

	forums := []Forum{}

	for root.Next() {
		forum := root.Node()
		titleNode := www.GetChilds("/li/span/span", forum)
		if titleNode == nil {
			log.Fatal("Can't find forum title node")
			continue
		}
		titleNode.Next()
		topForumName :=www.GetValue(titleNode.Node())
		
		top := Forum{}
		top.Name = topForumName
		top.Id = 0

		lvl1ForumNode := www.GetChilds("/li/ul/li/span/a", forum)

		if lvl1ForumNode == nil || !lvl1ForumNode.Next() {
			log.Fatal("Can't find first level forum name node")
			continue
		}

		lvl1ForumName := www.GetValue(lvl1ForumNode.Node())

	}

	return forums, nil
}

func (this *Tracker) Search(what string) {
	c := www.Context{}
	c.URL = "http://rutracker.org/forum/tracker.php?nm=" + what
	c.Method = "POST"
	c.Cookies = this.Cookies
	c.Data = map[string]string{ "nm" : what, "f" : "-1" }
	c.Parameters = map[string]string{ "nm" : what }
	r, e := www.Do(&c)
	if e != nil {
		log.Fatal(fmt.Sprintf("Can't search for '%s'", what))
		return
	}
	searchXml, saveErr := www.SaveResponse(r)
	if saveErr != nil {
		log.Fatal(fmt.Sprintf("Can't save result for '%s'", what))
		return	
	}

	root, parseE := www.ParseAndGet(searchXml, "//*[@id='tor-tbl']/tbody/tr/td/a")
	if parseE != nil {
		log.Fatal(fmt.Sprintf("Can't parse search result for '%s'", what))
		return		
	}


	for root.Next() {
		log.Print(www.GetAttr("href", root.Node()))
	}
}