package main
import (
    "fmt"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "i am a go")
}
func main(){
    fmt.Println("basic web server is starting on port 8080")
http.HandleFunc("/",indexHandler)
http.ListenAndServe(":8080",nil)
}
