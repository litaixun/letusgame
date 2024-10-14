package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	secretNumber int
	tries        int
)

// 初始化游戏
func initGame() {
	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(100) + 1
	tries = 0
}

// 处理首页请求
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Message": "欢迎来到猜数字游戏！请输入1到100之间的数字。",
	}

	tmpl.Execute(w, data)
}

// 处理猜测请求
func guessHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		guessStr := r.FormValue("guess")
		guess, err := strconv.Atoi(guessStr)
		if err != nil || guess < 1 || guess > 100 {
			data := map[string]interface{}{
				"Message": "请输入有效的数字（1到100之间）。",
			}
			tmpl.Execute(w, data)
			return
		}

		tries++
		var message string
		var gameOver bool

		if guess < secretNumber {
			message = fmt.Sprintf("猜小了！你已经猜了 %d 次。", tries)
		} else if guess > secretNumber {
			message = fmt.Sprintf("猜大了！你已经猜了 %d 次。", tries)
		} else {
			message = fmt.Sprintf("恭喜你，猜对了！正确数字是 %d，你一共猜了 %d 次。", secretNumber, tries)
			gameOver = true
			initGame() // 游戏重新开始
		}

		data := map[string]interface{}{
			"Message":  message,
			"GameOver": gameOver,
		}
		tmpl.Execute(w, data)
	}
}

func main() {
	// 初始化游戏
	initGame()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/guess", guessHandler)

	fmt.Println("Server running on localhost:12345")
	log.Fatal(http.ListenAndServe(":12345", nil))
}

