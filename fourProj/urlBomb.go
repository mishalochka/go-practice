package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}

// 		b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 	}
// }

/////////////////////////////////////////////ex///////////2/////////
// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := io.Copy(os.Stdout, resp.Body)
// 		//b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)

// 	}
// }

// /////////////////////////////////////////////ex3
// func main() {
// 	for _, url := range os.Args[1:] {

// 		if !strings.HasPrefix(url, "http://") {
// 			url = "http://" + url
// 		}
// 		fmt.Println(url)

// 		resp, err := http.Get(url)
// 		fmt.Println(resp.Status)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}

// 		_, err = io.Copy(os.Stdout, resp.Body)
// 		resp.Body.Close()

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
// 			os.Exit(1)
// 		}

// 		fmt.Println()

// 	}
// }

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprint("%.2fs %7d %s", secs, nbytes, url)
}
