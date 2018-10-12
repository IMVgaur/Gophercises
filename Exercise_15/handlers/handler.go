package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

//Handler function will handle all the routes
//Return handler func
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", sourceCodeNavigator)
	mux.HandleFunc("/panic", PanicDemo)
	mux.HandleFunc("/", welcome)
	return mux
}

//This sourceCodeNavigator function rendera the link to the program where error
//was occured, with line number
func sourceCodeNavigator(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		line = -1
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	io.Copy(b, file)
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	formatter := html.New(html.TabWidth(2), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	formatter.Format(w, style, iterator)

}

//welcome function displays the welcome message
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome!!!</h1>")
	w.WriteHeader(http.StatusOK)
}

//PanicDemo function is used to generate panic for server.
func PanicDemo(w http.ResponseWriter, r *http.Request) {
	panic("Error occured!!!")
}
