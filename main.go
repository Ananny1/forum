package main

import (
	"fmt"
	auth "forum/Auth"
	"forum/database"
	handlerfuncitons "forum/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// no reason for adding this.. i just want to push my stuff

func main() {

	// Open the SQLite database file (creates it if it doesn't exist)
	// Create tables
	database.Database()

	// handlers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) 
	//http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))
	http.HandleFunc("/", handlerfuncitons.Homepage)
	http.HandleFunc("/register", auth.Registerhandler)
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/logout", auth.LogoutHandler)
	http.HandleFunc("/post", handlerfuncitons.Post)
	http.HandleFunc("/category", handlerfuncitons.Category)
	http.HandleFunc("/welcome", handlerfuncitons.Welcomepage)
	http.HandleFunc("/createpost", handlerfuncitons.Createpostshandler)
	http.HandleFunc("/like", handlerfuncitons.LikesHandler)
	http.HandleFunc("/dislike",handlerfuncitons.DisLikesHandler)
	http.HandleFunc("/comment",handlerfuncitons.Comment)
	http.HandleFunc("/commentlike", handlerfuncitons.CommentLikeHandler)
	http.HandleFunc("/commentdislike", handlerfuncitons.CommentDisLikeHandler)
	http.HandleFunc("/filter", handlerfuncitons.FilterHandler)
	http.HandleFunc("/createcomment", handlerfuncitons.CreateComment)
	http.HandleFunc("/profile", handlerfuncitons.Profilehandler)
	http.HandleFunc("/auth/google/login" , auth.GoogleLoginHandler)
	http.HandleFunc("/auth/google/callback" , auth.GoogleCallbackHandler)
	http.HandleFunc("/auth/github/login" , auth.GitHubHandler)
	http.HandleFunc("/auth/github/callback" , auth.GitHubCallbackHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
