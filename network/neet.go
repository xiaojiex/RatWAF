package neet

import (
	gorrm "RatWAF/database"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func ReverseProxy() {
	addr := "127.0.0.1:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//这里打印详细信息
		reqDump, err := httputil.DumpRequest(r, true)
		reqStr := string(reqDump)
		if err != nil {
			fmt.Println(err)
			return
		}

		remote := fmt.Sprintf("http://127.0.0.1:20023%s", r.URL.String())

		// method 请求方法 remote 拼接的url
		req, err := http.NewRequest(r.Method, remote, r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(w, res.Body)
		fmt.Println("Attack ip", r.RemoteAddr)
		fmt.Println("reqstr ", reqStr)
		gorrm.InsertWafLog(r.RemoteAddr, reqStr, remote)
	})

	fmt.Println("gateway runserver", addr)
	http.ListenAndServe(addr, nil)
}
