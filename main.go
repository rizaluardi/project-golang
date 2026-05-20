package main
import (
"fmt"
"net/http"
)
func main() {http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "Halo! Ini aplikasi Golang di VPS-mu RIZALUARDI")})
fmt.Println("Server jalan di port 8080")
http.ListenAndServe(":8080", nil)}
