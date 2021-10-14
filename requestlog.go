package requestlog

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"net/url"
)

func Log(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		u, _ := url.Parse(r.RequestURI)

		method := r.Method
		methodLen := len(method) + 2

		path := u.Path
		pathLen := len(path) + 2

		timeLen := 9

		query := strings.Join(strings.Split(u.RawQuery, "&"), ", ")
		queryLen := len(query) + 3

		t := time.Now()

		h := t.Hour()
		m := t.Minute()
		s := t.Second()

		requestTime := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

		topLine := "╭" +
			strings.Repeat("─", methodLen) +
			"┬" +
			strings.Repeat("─", pathLen) +
			strings.Repeat("─", queryLen) +
			"┬" +
			strings.Repeat("─", timeLen) +
			"─╮"

		bottomLine := "╰" +
			strings.Repeat("─", methodLen) +
			"┴" +
			strings.Repeat("─", pathLen) +
			strings.Repeat("─", queryLen) +
			"┴" +
			strings.Repeat("─", timeLen) +
			"─╯"

		centerLine := fmt.Sprintf("│ \033[33m%s\033[0m │ \033[36m%s\033[0m [\033[31m%s\033[0m] │ \033[34m%s\033[0m │", r.Method, path, query, requestTime)

		fmt.Printf("%s\n%s\n%s\n", topLine, centerLine, bottomLine)

		next.ServeHTTP(w, r)
	})
}
