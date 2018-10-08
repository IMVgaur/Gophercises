package handlers

import (
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"

	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", SourceCodeNavigator)
	mux.HandleFunc("/panic/", PanicDemo)
	mux.HandleFunc("/", Welcome)
	return mux
}

func SourceCodeNavigator(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		fmt.Println("Error occured while parsing lineStr", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error occured while opening file, ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	b := bytes.NewBuffer(nil)
	io.Copy(b, file)
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	if err != nil {
		fmt.Println("Error occured while tokenising lexer")
	}
	style := styles.Get("github")
	formatter := html.New(html.TabWidth(2), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	formatter.Format(w, style, iterator)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome!!!</h1>")
	w.WriteHeader(http.StatusOK)
}

func PanicDemo(w http.ResponseWriter, r *http.Request) {
	panic("Panic says : Statue.")
}
