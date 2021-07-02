package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/aureuneun/bitcoin/blockchain"
)

const templateDir string = "explorer/templates/"

var port string

var templates *template.Template

type context struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(rw, "Hello %s", port)

	// tmpl := template.Must(template.ParseFiles("templates/home.html"))
	// data := context{PageTitle: "Home", Blocks: blockchain.GetBlockchain().AllBlocks()}
	// tmpl.Execute(rw, data)

	data := context{PageTitle: "Home", Blocks: blockchain.Blockchain().Blocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.PostForm.Get("block")
		blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	handler := http.NewServeMux()
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
