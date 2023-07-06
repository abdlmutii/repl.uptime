package replup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Debug bool   `json:"debug"`
	API   bool   `json:"api"`
	Port  string `json:"port"`
	Path  string `json:"path"`
}

func sender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html dir="rtl">
		<head>
			<link rel="preconnect" href="https://fonts.googleapis.com">
			<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css">
			<link href="https://fonts.googleapis.com/css2?family=Amiri&family=Cairo:wght@200;300;400;500;600;700;800;900;1000&family=Changa:wght@500&family=El+Messiri:wght@500&family=Noto+Kufi+Arabic:wght@900&display=swap" rel="stylesheet">
			<title>Repl.uptime</title>
			<style>
				a:hover {
					color: #212732;
				}
		
				body {
					background: linear-gradient(to right, #000, #0c0f17, #151a24);
					display: flex;
					justify-content: center;
					font-family: 'Cairo', 'Amiri', serif;
					font-weight: 600;
					align-items: center;
					min-height: 100vh;
					margin: 0;
				}
		
				.container {
					text-align: center;
					color: white;
				}
		
				h1 {
					font-size: 1.5rem;
					color: #ffbf00;
				}
		
				p {
					font-size: 0.75rem;
					margin: 0.5rem 1.25rem;
				}
		
				a {
					display: inline-block;
					background-color: white;
					color: #0C0F17;
					padding: 4px 12px;
					font-size: 0.75rem;
					text-decoration: none;
					font-weight: bold;
					border-radius: 5.625px;
				}
			</style>
		</head>
		
		<body>
			<script src="https://stackpath.bootstrapcdn.com/bootstrap/5.0.0-alpha2/js/bootstrap.min.js"></script>
			<div class="container">
				<h1 dir="ltr"><i class="bi bi-lightning-charge-fill"></i> Repl.uptime</h1>
				<p>Don't worry about downtime anymore! With Repl.uptime, your project will work 24 hours a day, 365 days a year. Repl.uptime is a Go version inspired by the npm package of the same name. It is maintained by Abdlmu'tii, who has dedicated efforts to ensure the reliability and seamless operation of the Go implementation.
    Rights reserved for both of Abdlmu'tii for maintaining the inspired-go library & Shuruhatik for maintaining the original npm package 
				<a dir="ltr" href="https://github.com/abdlmutii/repl.uptime" target="_blank" rel="noopener noreferrer"><i class="bi bi-box-arrow-up-right"></i> More info</a>&nbsp&nbsp&nbsp
         <a dir="ltr" href="https://github.com/abdlmutii" target="_blank" rel="noopener noreferrer"><i class="bi bi-box-arrow-up-right"></i>Abdlmu'tii GitHub</a>

         <p></p>
			</div>
		</body>
		</html>`)
}

func Uptime() {
	var conf Config
	data, err := ioutil.ReadFile("uptime.json")
	if err != nil {
		fmt.Println("You have create an configuration file, check the readme again.")
		return
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}

		http.HandleFunc(conf.Path, sender)
		serverErr := http.ListenAndServe(":"+conf.Port, nil)
	if serverErr != nil {
		fmt.Println(serverErr)
	} else {
      fmt.Println("✅ You can rely on this project to stay up and running without any interruptions.")
  }
  
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			url := fmt.Sprintf("https://api.shuruhatik.com/uptime/%s/%s/%s%s", os.Getenv("REPL_ID"), os.Getenv("REPL_SLUG"), os.Getenv("REPL_OWNER"), conf.Path)
			response, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}
			if body != nil && conf.Debug == true {
				fmt.Println(body)
			}
		}
	}()

	select {}
}
