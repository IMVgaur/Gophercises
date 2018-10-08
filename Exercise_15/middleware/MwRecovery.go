package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
)

func RecoveryMid(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>%v</h1><pre>%s</pre>", err, ErrLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func ErrLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for index, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		filePath := ""
		for i, ch := range line {
			if ch == ':' {
				filePath = line[1:i]
				break
			}
		}
		var lineStr strings.Builder
		for count := len(filePath) + 2; count < len(line); count++ {
			if line[count] < '0' || line[count] > '9' {
				break
			}
			lineStr.WriteByte(line[count])
		}
		v := url.Values{}
		v.Set("path", filePath)
		v.Set("line", lineStr.String())
		lines[index] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + filePath + ":" + lineStr.String() + "</a>"
	}
	return strings.Join(lines, "\n")
}
