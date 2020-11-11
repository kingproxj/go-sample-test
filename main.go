package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	"net"
)

const middle = "========="

type Config struct {
	Mymap  map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		//panic(err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}

func RemoteIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Original-Forwarded-For"), ",")[0])
	if ip != "" {
		fmt.Printf("RemoteIP X-Original-Forwarded-For: %s\n", ip)
	}

	ip = strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		fmt.Printf("RemoteIP X-Forwarded-For: %s\n", ip)
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		fmt.Printf("RemoteIP X-Real-Ip: %s\n", ip)
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		fmt.Printf("RemoteIP RemoteAddr: %s\n", ip)
		return ip
	}

	return ip
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloHandler begin.")
	fmt.Println(r)
	rJson, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("json.Marshal error: %s\n", err)
	}
	fmt.Printf("rJson is %s\n", rJson)
	/*fmt.Printf("X-Forwarded-For: %s\n", r.Header.Get("X-Forwarded-For"))
	fmt.Printf("HTTP_X_FORWARDED_FOR: %s\n", r.Header.Get("HTTP_X_FORWARDED_FOR"))
	fmt.Printf("REMOTE-HOST: %s\n", r.Header.Get("REMOTE-HOST"))
	fmt.Printf("X-Real-IP: %s\n", r.Header.Get("X-Real-IP"))*/
	RemoteIP(r)

	envTEMP := os.Getenv("TEMP")
	fmt.Fprintf(w, "Image:demo-go-A#####/##### Hello IOP Canary 333 aaa\n")
	if envTEMP != "" {
		fmt.Fprintf(w, "env TEMP is %v\n", envTEMP)
	}

	envTMP := os.Getenv("TMP")
	if envTMP != "" {
		fmt.Fprintf(w, "env TMP is %v\n", envTMP)
	}

	userFile := "/etc/config/BCompare_bak.ini"
	myConfig := new(Config)
	myConfig.InitConfig(userFile)
	installTime := myConfig.Read("BCompare", "InstallTime")
	if installTime != "" {
		fmt.Fprintf(w, "配置信息 installTime is %v\n", installTime)
	}
	/*fmt.Println(myConfig.Read("BCompare", "InstallTime"))
	fmt.Printf("%v\n", myConfig.Mymap)*/


	secret, err := ioutil.ReadFile("/etc/secret/password")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(secret)
	secretstr := string(secret)
	//if secretstr != "" {
	fmt.Fprintf(w, "我的密钥信息：\n %s",secretstr)
	//}

	fmt.Println("helloHandler is called.")
	fmt.Println("helloHandler end.")
}

func helloWhoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWhoHandler begin.")
	fmt.Println(r)
	/*fmt.Printf("X-Forwarded-For: %s\n", r.Header.Get("X-Forwarded-For"))
	fmt.Printf("HTTP_X_FORWARDED_FOR: %s\n", r.Header.Get("HTTP_X_FORWARDED_FOR"))
	fmt.Printf("REMOTE-HOST: %s\n", r.Header.Get("REMOTE-HOST"))
	fmt.Printf("X-Real-IP: %s\n", r.Header.Get("X-Real-IP"))*/
	RemoteIP(r)

	fmt.Fprintf(w, "Image:demo-go-A#####/hello-who#####")
	fmt.Println("helloWhoHandler end.")
	w.Write([]byte("Hello IOP Caleb 0829"))
}

func goodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "Image:demo-go-A#####/good#####")
	w.Write([]byte("Good Caleb 0830"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/hello-who", helloWhoHandler)
	mux.HandleFunc("/good", goodHandler)
	http.ListenAndServe("0.0.0.0:8888", mux)
}
