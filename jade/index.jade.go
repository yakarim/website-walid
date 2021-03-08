// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"io"

	"github.com/yakarim/website-walid/database"
)

const (
	index__0 = `<!DOCTYPE html><html><head><meta charset="UTF-8"/><title>`
	index__1 = `</title><link rel="stylesheet" href="/css/main.css"/><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-BmbxuPwQa2lc/FVzBcNJ7UAyJxM6wuqIj61tLrc4wSX0szH/Ev+nYRRuWlolflfl" crossorigin="anonymous"/><script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/js/bootstrap.bundle.min.js" integrity="sha384-b5kHyXgcpbZJO/tY9Ul7kGkf1S0CWuKcCD38l8YkeH8z8QjE0GmW1gYU5S9FOnJ0" crossorigin="anonymous"></script><link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css"/></head><body><nav class="navbar navbar-expand-lg navbar-light bg-light"><div class="container"><a class="navbar-brand" href="./">Website Walid</a><button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button><div id="navbarSupportedContent" class="collapse navbar-collapse"><ul class="navbar-nav me-auto mb-2 mb-lg-0"><li class="nav-item"><a class="nav-link active" aria-current="page" href="./">Home</a></li>`
	index__2 = `</ul><form class="d-flex"><input class="form-control me-2" type="search" placeholder="Search" aria-label="Search"/><button class="btn btn-outline-success" type="submit">Search</button></form></div></div></nav><div class="container"><h1>My Article</h1><div></div><table class="table-bordered"><thead><tr><th>id</th><th>name</th></tr></thead><tbody>`
	index__3 = `</tbody></table><div class="footer">2021</div></div></body></html>`
	index__4 = `<li class="nav-item"></li><a class="nav-link" href="./admin">Admin</a><a class="nav-item nav-link" href="./login">Logout (`
	index__5 = `)</a>`
	index__6 = `<a class="nav-item nav-link" href="./login">Login</a>`
	index__7 = `<tr><td>`
	index__8 = `</td><td>`
	index__9 = `</td></tr>`
)

func Index(pageTitle, user string, signIn bool, data []database.Images, wr io.Writer) {
	buffer := &WriterAsBuffer{wr}

	buffer.WriteString(index__0)
	WriteEscString(pageTitle, buffer)
	buffer.WriteString(index__1)

	if signIn {
		buffer.WriteString(index__4)
		WriteEscString(user, buffer)
		buffer.WriteString(index__5)

	} else {
		buffer.WriteString(index__6)

	}
	buffer.WriteString(index__2)

	for _, post := range data {
		buffer.WriteString(index__7)
		WriteAll(post.UID, true, buffer)
		buffer.WriteString(index__8)
		WriteAll(post.Name, true, buffer)
		buffer.WriteString(index__9)

	}
	buffer.WriteString(index__3)

}
