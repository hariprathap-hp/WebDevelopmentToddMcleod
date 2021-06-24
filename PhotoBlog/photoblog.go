package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/home/hari/go/src/WebDevelopmentTodd/PhotoBlog/pics/", http.StripPrefix("/home/hari/go/src/WebDevelopmentTodd/PhotoBlog/pics/", http.FileServer(http.Dir("./pics"))))
	//http.HandleFunc("/public/", picsHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)

	if r.Method == http.MethodPost {
		fmt.Println("Method is Post")
		r.ParseForm()
		file, fileHeader, err := r.FormFile("newfile")
		if err != nil {
			log.Println("Error fetching file")
		}

		defer file.Close()

		//split the file name and the extension as file name will be converted to hash
		ext := strings.Split(fileHeader.Filename, ".")[1]

		//Now using sha1 function create a hash key for file name
		hash := sha1.New()
		io.Copy(hash, file)
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
		wd, dir_err := os.Getwd()
		if dir_err != nil {
			log.Printf("%s", "Error Getting Current Working Directory")
		}

		path := filepath.Join(wd, "pics", fname)
		fmt.Println(path)

		//so far you have created just a hashed name for the uploaded file along with the checksum for the image
		nf, n_err := os.Create(path)
		if n_err != nil {
			log.Println("File creation failed with error : ", n_err)
		}
		defer nf.Close()
		fmt.Println("Seeking file location")
		file.Seek(0, 0)
		io.Copy(nf, file)
		cookie = appendNewValues(w, cookie, path)
		//http.ServeFile(w, r, path)
	}

	xs := strings.Split(cookie.Value, "|")
	tmpl.ExecuteTemplate(w, "index.gohtml", xs)

}

func appendNewValues(w http.ResponseWriter, c *http.Cookie, pic string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, pic) {
		s += "|" + pic
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	fmt.Println("Inside getCookie")
	cookie, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}
