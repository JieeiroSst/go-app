package integration_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	domain "hithub.com/JIeeiroSst/go-app/domain"
)

//author
func AuthorById(t *testing.T){
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:8080/author/%d",id), nil)
	assert.NoError(t, err)

	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	
	bodys, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	author := domain.Author{
		ID: 1,
		Name: "chmp",
	}	
	
	jsonData, _ := json.Marshal(author)
	fmt.Println(string(bodys))
	fmt.Println( string(jsonData)+"\n")
	assert.Equal(t, string(jsonData)+"\n", string(bodys))
}

func AuthorByAll(t *testing.T){
	req, err := http.NewRequest(echo.GET, "http://localhost:8080/author/", nil)
	assert.NoError(t, err)

	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	
	bodys, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	authors=[]domain.Author{
		{
			ID: 1,
			Name: "chmp",
		},
	}	
	
	jsonData, _ := json.Marshal(authors)
	fmt.Println(string(bodys))
	fmt.Println( string(jsonData)+"\n")
	assert.Equal(t, string(jsonData)+"\n", string(bodys))
}

func CreateAuthor(t *testing.T) {
	endpoint := "http://localhost:8080/author"
	data:=url.Values{}
	data.Set("id","1")
	data.Set("name","chmp")


	client := &http.Client{}
	r, _ := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(r)
	defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t,create_message, string(body))
}

func UpdateAuthor(t *testing.T) {
	endpoint := fmt.Sprintf("http://localhost:8080/author/%d",id)
	data:=url.Values{}
	data.Set("name","chmp1")

	client := &http.Client{}
	r, _ := http.NewRequest("PUT", endpoint, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(r)
	defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t,update_message, string(body))
}

func DeleteAuthor(t *testing.T) {
	req, err := http.NewRequest(echo.DELETE,fmt.Sprintf("http://localhost:8080/author/%d",id), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	bodys, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t,delete_message, string(bodys))
}

func CreateArticle(t *testing.T) {
	endpoint := "http://localhost:8080/article"
	data:=url.Values{}
	data.Set("id","1")
	data.Set("title","chmp")
	data.Set("body","chmp")
	data.Set("author_id","1")


	client := &http.Client{}
	r, _ := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(r)
	defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t,create_message, string(body))
}

func UpdateArticle1(t *testing.T) {
	endpoint := fmt.Sprintf("http://localhost:8080/article/%d",id)
	data:=url.Values{}
	data.Set("title","chmp")
	data.Set("body","chmp")
	data.Set("author_id","1")

	client := &http.Client{}
	r, _ := http.NewRequest("PUT", endpoint, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := client.Do(r)
	defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t,update_message, string(body))
}

func DeleteArticle(t *testing.T) {
	req, err := http.NewRequest(echo.DELETE,fmt.Sprintf("http://localhost:8080/article/%d",id), nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	bodys, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t,delete_message, string(bodys))
}

func ArticleAll(t *testing.T) {
	req, err := http.NewRequest(echo.GET, "http://localhost:8080/article/", nil)
	assert.NoError(t, err)

	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	
	bodys, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	listArticle=[]domain.Article{
		{
			ID: 1,
			Title: "Name",
			Body: "name",
			AuthorId: 1,
			Author: domain.Author{
				ID: 1,
				Name:"chmp",
		},
	},
	}
	
	jsonData, _ := json.Marshal(listArticle)
	fmt.Println(string(bodys))
	fmt.Println( string(jsonData)+"\n")
	assert.Equal(t, string(jsonData)+"\n", string(bodys))
}

func ArticleById(t *testing.T){
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:8080/article/%d",id), nil)
	assert.NoError(t, err)

	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	
	bodys, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	Article=domain.Article{
		ID: 1,
		Title: "Name",
		Body: "name",
		AuthorId: 1,
		Author: domain.Author{
			ID: 1,
			Name:"chmp",
		},
	}
	
	jsonData, _ := json.Marshal(Article)
	fmt.Println(string(bodys))
	fmt.Println( string(jsonData)+"\n")
	assert.Equal(t, string(jsonData)+"\n", string(bodys))
}