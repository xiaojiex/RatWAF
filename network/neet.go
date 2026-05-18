package neet

import (
	gorrm "RatWAF/database"
	"RatWAF/regex"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func ReverseProxy() {
	addr := "127.0.0.1:8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)
		reqStr := string(reqDump)
		if err != nil {
			fmt.Println(err)
			return
		}
		remote := fmt.Sprintf("http://127.0.0.1:20023%s", r.URL.String())
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
		str := r.RemoteAddr + reqStr + remote
		
		sql_ok, attackTypeSQL := regex.SQL_Injection(str)
		rce_ok, attackTypeRce := regex.RCE(str)
		if sql_ok {
			fmt.Println("TRUE", str)
			fmt.Println("type", attackTypeSQL)
			gorrm.InsertWafLog(r.RemoteAddr, reqStr, remote, attackTypeSQL)
		}
		if rce_ok {
			fmt.Println("TRUE", str)
			fmt.Println("type", attackTypeRce)
			gorrm.InsertWafLog(r.RemoteAddr, reqStr, remote, attackTypeRce)
		}
	})

	fmt.Println("gateway runserver", addr)
	http.ListenAndServe(addr, nil)
}
