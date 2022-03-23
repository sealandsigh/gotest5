package db

import (
	"github.com/sealandsigh/gotest5/day15/blog/model"
	"testing"
	"time"
)

func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:zhujiajun@tcp(localhost:3306)/goday10?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 插入文章

func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "abcdefghigklmn"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "zjj"
	article.ArticleInfo.Username = "zhu"
	article.ArticleInfo.Summary = "abc fd"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("articleId: %d", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("article lens is: %d", len(articleList))
}

func TestGetArticleDetail(t *testing.T) {
	detail, err := GetArticleDetail(1)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("article lens is: %#v\n", detail)
}

func TestGetArticleListByCategoryId(t *testing.T) {
	ar, err := GetArticleListByCategoryId(1, 1, 15)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("article by categoryId is: %#v\n", ar)
	for _, info := range ar {
		t.Logf("ar content is: %#v", info.Title)
	}
}

func TestGetArticleInfo(t *testing.T) {
	articleInfo, err := GetArticleDetail(5)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, article:%#v\n", articleInfo)
}
