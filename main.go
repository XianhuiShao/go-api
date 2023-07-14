package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	//Array string  `json:"array"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.POST("/albums", postAlbums)

	//router.POST("/albums", AddPost)

	router.Run(":50016")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {

	// a := c.Request

	var a string
	a = "edsr"
	c.IndentedJSON(http.StatusOK, a)
	//c.XML(200, a)

}

func postAlbums(c *gin.Context) {
	name := c.PostForm("name")

	fmt.Printf("name: %s", name)

	//a := c.Request.Body

	//c.IndentedJSON(http.StatusOK, name)
	//c.String(http.StatusOK, string(name))

	//len := c.Request.ContentLength

	//c.String(http.StatusOK, len)
	//a1 :=c.Bind()

	decoder := json.NewDecoder(c.Request.Body)
	var al album

	err := decoder.Decode(&al)
	if err != nil {
		panic(err)

	}
	// b, err := json.Marshal(al)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//	str := string(b)
	//c.IndentedJSON(http.StatusOK, al)
	//	c.String(http.StatusOK, str)

	fmt.Printf("name: %s", al.ID)

	// s, _ := ioutil.ReadAll(c.Request.Body)
	// str := string(s)
	// StatusOK := 200
	// if str == "" {
	// 	StatusOK = 400
	// } else {

	// }
	//c.String(StatusOK, str)

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	c.String(http.StatusOK, buf.String())

}

func AddPost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength          // 获取请求实体长度
	body := make([]byte, len)       // 创建存放请求实体的字节切片
	r.Body.Read(body)               // 调用 Read 方法读取请求实体并将返回内容存放到上面创建的字节切片
	io.WriteString(w, string(body)) // 将请求实体作为响应实体返回
}
