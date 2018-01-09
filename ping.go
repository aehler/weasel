package main

import (
	"net/http"
	"fmt"
	"math/rand"
	"io/ioutil"
	"time"
	"strings"
)

func main() {

	fmt.Println("Pinging weasel...")

	for i := 0; i < 100; i++ {

		v := fmt.Sprintf("%d", rand.Int())

		go ping(v)

	}

	time.Sleep(time.Second * 5)

}

func ping(v string) {

	c, err := http.Get("http://127.0.0.1:8082/pong/?v=" + v)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	body, err := ioutil.ReadAll(c.Body)

	c.Body.Close()

	fmt.Println("Success", fmt.Sprintf(`"%s"`, v), "=>", string(body), fmt.Sprintf(`"%s"`, v) == strings.TrimSpace(string(body)))

}