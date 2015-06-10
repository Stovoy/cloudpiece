package model
import "fmt"

type Post struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

func ReadPosts() []Post {
	rows, err := conn.Query("SELECT title, body FROM post")
	if err != nil {
		fmt.Println(err)
	}
	posts := make([]Post, 0)
	for rows.Next() {
		var title string
		var body string
		rows.Scan(&title, &body)
		posts = append(posts, Post{title, body})
	}
	return posts
}
