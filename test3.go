package main
//多路复用
import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,"hello test a")
	})
	mux.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,"hello test b")
	})
	mux.HandleFunc("/c",testc)
	http.ListenAndServe(":8003",mux)
}
func testc(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"hello test c")
}
