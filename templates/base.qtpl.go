// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line templates/base.qtpl:3
package templates

//line templates/base.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/base.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/base.qtpl:4
type Page interface {
//line templates/base.qtpl:4
	Title() string
//line templates/base.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line templates/base.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line templates/base.qtpl:4
	Body() string
//line templates/base.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line templates/base.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line templates/base.qtpl:4
}

// Page prints a page implementing Page interface.

//line templates/base.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line templates/base.qtpl:12
	qw422016.N().S(`
<html>
	<head>
	 <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/css/bootstrap.min.css" rel="stylesheet"
    integrity="sha384-BmbxuPwQa2lc/FVzBcNJ7UAyJxM6wuqIj61tLrc4wSX0szH/Ev+nYRRuWlolflfl" crossorigin="anonymous">
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta2/dist/js/bootstrap.bundle.min.js"
    integrity="sha384-b5kHyXgcpbZJO/tY9Ul7kGkf1S0CWuKcCD38l8YkeH8z8QjE0GmW1gYU5S9FOnJ0"
    crossorigin="anonymous"></script>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css">
		<title>`)
//line templates/base.qtpl:21
	p.StreamTitle(qw422016)
//line templates/base.qtpl:21
	qw422016.N().S(`</title>
	</head>
	<body>
	<div class="container">
		<div>
			<a href="/">return to main page</a>
		</div>
		`)
//line templates/base.qtpl:28
	p.StreamBody(qw422016)
//line templates/base.qtpl:28
	qw422016.N().S(`
	</div>
	</body>
</html>
`)
//line templates/base.qtpl:32
}

//line templates/base.qtpl:32
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line templates/base.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/base.qtpl:32
	StreamPageTemplate(qw422016, p)
//line templates/base.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line templates/base.qtpl:32
}

//line templates/base.qtpl:32
func PageTemplate(p Page) string {
//line templates/base.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/base.qtpl:32
	WritePageTemplate(qb422016, p)
//line templates/base.qtpl:32
	qs422016 := string(qb422016.B)
//line templates/base.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/base.qtpl:32
	return qs422016
//line templates/base.qtpl:32
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line templates/base.qtpl:37
type BasePage struct{}

//line templates/base.qtpl:38
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/base.qtpl:38
	qw422016.N().S(`This is a base title`)
//line templates/base.qtpl:38
}

//line templates/base.qtpl:38
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/base.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/base.qtpl:38
	p.StreamTitle(qw422016)
//line templates/base.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line templates/base.qtpl:38
}

//line templates/base.qtpl:38
func (p *BasePage) Title() string {
//line templates/base.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/base.qtpl:38
	p.WriteTitle(qb422016)
//line templates/base.qtpl:38
	qs422016 := string(qb422016.B)
//line templates/base.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/base.qtpl:38
	return qs422016
//line templates/base.qtpl:38
}

//line templates/base.qtpl:39
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/base.qtpl:39
	qw422016.N().S(`This is a base body`)
//line templates/base.qtpl:39
}

//line templates/base.qtpl:39
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/base.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/base.qtpl:39
	p.StreamBody(qw422016)
//line templates/base.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line templates/base.qtpl:39
}

//line templates/base.qtpl:39
func (p *BasePage) Body() string {
//line templates/base.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/base.qtpl:39
	p.WriteBody(qb422016)
//line templates/base.qtpl:39
	qs422016 := string(qb422016.B)
//line templates/base.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/base.qtpl:39
	return qs422016
//line templates/base.qtpl:39
}
