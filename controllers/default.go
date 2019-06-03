package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

var atc *ArticlesOfType

type Controller struct {
	beego.Controller
	Tag           []*Tag
	Lable         []*Label
	LabelTypeList []int
}

type Tag struct {
	TagName  string
	Path     string
	IsActive bool
}

type Label struct {
	TagName string
	Types   int
}

func (c *Controller) setTagActive(tagPath string) {
	for _, v := range c.Tag {
		v.IsActive = false
		if v.Path == tagPath {
			v.IsActive = true
		}
	}
}

type ArtTim struct {
	Y   int
	M   int
	D   int
	H   int
	Min int
}

type Article struct {
	ID    int
	Title string
	Brief string
	Img   string
	Type  int
	Text  string
	IsNew bool
	Tim   *ArtTim
}

type Articles struct {
	ID2Article map[int]*Article
}

type ArticlesOfType struct {
	Type2Articles map[int]*Articles
	TotalNum      int
}

func (a *Articles) getArticleByID(aid int) *Article {
	if v, ok := a.ID2Article[aid]; ok {
		return v
	}
	return nil
}

func (a *ArticlesOfType) getArticlesByType(aty int) *Articles {
	if v, ok := a.Type2Articles[aty]; ok {
		return v
	}
	return nil
}

func (a *ArticlesOfType) getAllArticles() *Articles {
	art := &Articles{
		ID2Article: map[int]*Article{},
	}
	for _, v := range a.Type2Articles {
		for k, as := range v.ID2Article {
			art.ID2Article[k] = as
		}
	}
	return art
}

func (a *ArticlesOfType) initArticles(ty int) {
	art := &Articles{
		ID2Article: map[int]*Article{},
	}
	t := &ArtTim{
		Y:   2018,
		M:   10,
		D:   1,
		H:   19,
		Min: 12,
	}
	atc := &Article{
		ID:    ty*100 + 1,
		Title: "空间立体效果图，完美呈现最终效果1",
		Brief: "室内设计作为一门新兴的学科，尽管还只是近数十年的事，但是人们有意识地对自己生活、生产活动的室内进行安排布置，甚至美化装饰，赋予室内环境以所祈使的气氛，却早巳从人类文明伊始的时期就已存在",
		Img:   "sy_img1.jpg",
		Type:  ty,
		IsNew: true,
		Tim:   t,
	}
	art.ID2Article[atc.ID] = atc
	t2 := &ArtTim{
		Y:   2019,
		M:   12,
		D:   8,
		H:   19,
		Min: 12,
	}
	a2 := &Article{
		ID:    ty*100 + 2,
		Title: "空间立体效果图，完美呈现最终效果2",
		Brief: "室内设计作为一门新兴的学科，尽管还只是近数十年的事，但是人们有意识地对自己生活、生产活动的室内进行安排布置，甚至美化装饰，赋予室内环境以所祈使的气氛，却早巳从人类文明伊始的时期就已存在",
		Img:   "sy_img1.jpg",
		Type:  ty,
		IsNew: false,
		Tim:   t2,
	}
	art.ID2Article[a2.ID] = a2
	a.Type2Articles[ty] = art
	a.TotalNum = 2
}

func (c *Controller) initData() {
	c.Data["MyNickname"] = beego.AppConfig.String("myNickname")
	c.Data["MyEmail"] = beego.AppConfig.String("myEmail")
	c.Data["PhoneNum"] = beego.AppConfig.String("phoneNum")
	c.Data["WeChatNum"] = beego.AppConfig.String("weChatNum")
	c.Data["ViewsNum"] = 999
	c.Data["ArticleNum"] = 88
	tagl := beego.AppConfig.Strings("tag")
	pathl := beego.AppConfig.Strings("pagePath")
	if len(tagl) != len(pathl) {
		beego.Error("tag len != path len, inspect the conf/app.conf")
	}
	for i := 0; i < len(tagl); i++ {
		tag := &Tag{}
		tag.TagName = tagl[i]
		tag.Path = pathl[i]
		c.Tag = append(c.Tag, tag)
	}

	lb := beego.AppConfig.Strings("lable")
	lp := beego.AppConfig.Strings("lableLink")
	if len(lb) != len(lp) {
		beego.Error("tag len != path len, inspect the conf/app.conf, err:")
	}
	for i := 0; i < len(lb); i++ {
		lab := &Label{}
		lab.TagName = lb[i]
		li, err := strconv.Atoi(lp[i])
		if err != nil {
			beego.Error(err)
		}
		lab.Types = li
		c.LabelTypeList = append(c.LabelTypeList, li)
		c.Lable = append(c.Lable, lab)
	}
	c.setTagActive("/")

	atc = &ArticlesOfType{
		Type2Articles: map[int]*Articles{},
	}
	atc.initArticles(1)
	atc.initArticles(3)
}

func (c *Controller) Get() {
	c.initData()
	//var aix int
	//var err error
	//artIndex := c.Ctx.Input.Param(":id")
	//if artIndex != "" {
	//	c.setLableActive(artIndex)
	//	aix, err = strconv.Atoi(artIndex)
	//	if err != nil {
	//		beego.Error("Controller's Get func string to int err=%s, path=%s", err, artIndex)
	//	}
	//}
	//switch aix {
	//case consts.ProjectDisplay:
	//	c.Data["Articles"] = atc.getArticlesByType(consts.ProjectDisplay)
	//case consts.LifeMiscellany:
	//	c.Data["Articles"] = atc.getArticlesByType(consts.LifeMiscellany)
	//case consts.PersonalQuotations:
	//	c.Data["Articles"] = atc.getArticlesByType(consts.PersonalQuotations)
	//default:
	//	c.Data["Articles"] = atc.getArticlesByType(consts.TechBlog)
	//}

	c.Data["Articles"] = atc
	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Data["LabelList"] = c.LabelTypeList
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *Controller) Index() {
	c.initData()
	c.setTagActive("index")

	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Layout = "layout.html"
	c.TplName = "index.html"
}

func (c *Controller) About() {
	c.initData()
	c.setTagActive("about")
	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Layout = "layout.html"
	c.TplName = "about.html"
}

func (c *Controller) Album() {
	c.initData()
	c.setTagActive("album")
	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Layout = "layout.html"
	c.TplName = "album.html"
}

func (c *Controller) Details() {
	c.initData()
	c.setTagActive("details")
	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Layout = "layout.html"
	c.TplName = "details.html"
}

func (c *Controller) Leacots() {
	c.initData()
	c.setTagActive("leacots")
	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Layout = "layout.html"
	c.TplName = "leacots.html"
}

func (c *Controller) Whisper() {
	c.initData()
	c.setTagActive("whisper")
	for _, v := range c.Tag {
		beego.Debug(v.Path, v.TagName, v.IsActive)
	}
	c.Data["Tag"] = c.Tag
	c.Data["Lable"] = c.Lable
	c.Layout = "layout.html"
	c.TplName = "whisper.html"
}
