package main  
  
import (  
    "fmt"  
    "log"  
    "math/rand"  
    "net/http"  
    "time"  
)  
  
var headers = []string{  
    "Server: Apache/2.4.29 (Ubuntu)",  
    "X-Powered-By: PHP/7.2.24-0ubuntu0.18.04.7",  
    "Link: <http://example.com/wp-json/>; rel=\"https://api.w.org/\"",  
    "Content-Type: text/html; charset=UTF-8",  
    "X-Cache: MISS",  
}  
  
func main() {  
    rand.Seed(time.Now().UnixNano())  
  
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {  
        fmt.Fprintln(w, "<html><head><title>WordPress Site</title></head><body><h1>Welcome to WordPress</h1></body></html>")  
    })  
  
    http.HandleFunc("/wp-login.php", func(w http.ResponseWriter, r *http.Request) {  
        if r.Method == "POST" {  
            username := r.FormValue("log")  
            password := r.FormValue("pwd")  
            ip := r.RemoteAddr  
            fmt.Printf("Hacker tried to login with username %s and password %s from IP %s\n", username, password, ip)  
        }  
        fmt.Fprintln(w, "<html><head><title>WordPress Admin Login</title></head><body><h1>WordPress Admin Login</h1><form method=\"post\" action=\"/wp-login.php\"><label>Username:</label><input type=\"text\" name=\"log\"><br><label>Password:</label><input type=\"password\" name=\"pwd\"><br><input type=\"submit\" value=\"Log In\"></form></body></html>")  
    })  
  
    log.Fatal(http.ListenAndServe(":80", randomHeader(http.DefaultServeMux)))  
}  
  
func randomHeader(handler http.Handler) http.Handler {  
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  
        headerIndex := rand.Intn(len(headers))  
        header := headers[headerIndex]  
        w.Header().Set("Server", header)  
        w.Header().Set("X-Powered-By", header)  
        w.Header().Set("Link", header)  
        w.Header().Set("Content-Type", header)  
        w.Header().Set("X-Cache", header)  
        handler.ServeHTTP(w, r)  
    })  
}  
