//Package main is used to serve a page
package main

import (
	"GoExample/GuateGolang/data"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//home redirct the client's petition on / to home.html
func home(w http.ResponseWriter, r *http.Request) {
	j := data.CreateUser("Juan", "Juan356@gmail.com", "23", "no phone", "/common/images/user/mypic.jpg")
	templ, err := template.ParseFiles("html/home.html")
	if err != nil {
		fmt.Println(err)
	}
	templ.Execute(w, j)
}

//GetPort gives you a port  of the current system.
func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

//main is the entry point
//main es el punto de entrada para todo programa en go
func main() {
	//writing a muktiplexer
	//declarando un nuevo  multiplexer
	mux := http.NewServeMux()

	//serving the css file picture and so on
	//sirviendo el archivo css  fotos y demas que estan dentro del directorio  common
	mux.Handle("/common/", http.StripPrefix("/common/", http.FileServer(http.Dir("common"))))

	//when the user request / whe call the home handler
	mux.HandleFunc("/", home)
	//making a serever
	//creando un servidor
	server := &http.Server{
		//listening to port 8080
		//escuchar en el puerto 8080
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
