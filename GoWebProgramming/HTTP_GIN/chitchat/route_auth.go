package main

import (
	"WebDevelopmentTodd/GoWebProgramming/HTTP_GIN/chitchat/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authenticate(ctx *gin.Context) {
	ctx.Request.ParseForm()
	user, err := data.UserbyEmail(ctx.Request.FormValue("email"))
	if err != nil {
		danger(err, "Cannot find the user")
	}
	if user.Password == data.Encrypt(ctx.Request.FormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(ctx.Writer, &cookie)
		http.Redirect(ctx.Writer, ctx.Request, "/", 302)
	} else {
		http.Redirect(ctx.Writer, ctx.Request, "/login", 302)
	}
}

func login(ctx *gin.Context) {
	parseTemplateFiles("login.layout", "public.navbar", "login")
	ctx.HTML(200, "layout", nil)
}

func signup(ctx *gin.Context) {
	generateHTML(ctx, nil, "login.layout", "public.navbar", "signup")
}

func logout(ctx *gin.Context) {
	fmt.Println("Logout")
	cookie, err := ctx.Request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		fmt.Println("Error is -- ", err)
		//warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(ctx.Writer, ctx.Request, "/", 302)
}

func signupAccount(ctx *gin.Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		danger(err, "cannot parse form")
	}
	name := ctx.Request.FormValue("name")
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	user := data.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	fmt.Println("User is -- ", user)
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(ctx.Writer, ctx.Request, "/login", 302)
}
