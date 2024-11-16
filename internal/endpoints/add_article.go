package endpoints

import (
	"article/internal/article"
	"log"
	"time"
)

type AddArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int    `json:"authorId"`
}

type AddArticleResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type AddArticle interface {
	Do(r AddArticleRequest) AddArticleResponse
}

type AddArticleBusiness struct {
	Service *article.Service
}

func (e *AddArticleBusiness) Do(r AddArticleRequest) AddArticleResponse {
	a := &article.Article{
		Title:   r.Title,
		Content: r.Content,
		Author: article.Author{
			Id: r.AuthorId,
		},
	}

	resp := AddArticleResponse{
		Id: a.Id,
	}

	err := e.Service.Add(a)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "error"
	} else {
		resp.Status = "OK"
	}

	return resp
}

type AddArticleLog struct {
	Endpoint AddArticle
}

func (e *AddArticleLog) Do(r AddArticleRequest) AddArticleResponse {
	log.Printf("%+v", r)
	resp := e.Endpoint.Do(r)
	log.Printf("%+v", resp)
	return resp
}

type AddArticleTime struct {
	Endpoint AddArticle
}

func (e *AddArticleTime) Do(r AddArticleRequest) AddArticleResponse {
	start := time.Now()
	resp := e.Endpoint.Do(r)
	end := time.Now()
	log.Println("add article completed in", end.Sub(start))
	return resp
}
