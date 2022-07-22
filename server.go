package main

import (
	"fmt"
	"net/http"
  "strings"
  "os"
  "encoding/hex"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

  uri := r.URL.String()
  uri2 := strings.Replace(uri, "/","",1)
  decoded, _ := hex.DecodeString(uri2)

  p, perr := os.OpenFile("outfile",
        os.O_WRONLY|os.O_APPEND|os.O_CREATE,
        0666)
      if perr != nil {
				fmt.Println("[-] Error creating the outfile...")
			}
			defer p.Close()

      p.Write(decoded)

}
