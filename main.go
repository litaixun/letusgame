package main

import (
    "math/rand"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
    "net/http"
)

var (
    targetNumber  int
    mu            sync.Mutex
    gameStartTime time.Time
)

func main() {
    router := gin.Default()
    router.Static("/static", "./static")

    router.LoadHTMLGlob("static/*.html")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // 各游戏的路由
    router.POST("/play/:game", playGame)
    router.POST("/play/guess", guessNumber)
    router.POST("/play/rps", playRockPaperScissors)
    router.POST("/play/timer", startTimer)

    router.Run(":12345") // 修改端口为 12345
}

// playGame 初始化游戏
func playGame(c *gin.Context) {
    game := c.Param("game")

    if game == "guess" {
        mu.Lock()
        targetNumber = rand.Intn(100) + 1 // 生成1到100之间的随机数
        mu.Unlock()
        c.JSON(http.StatusOK, gin.H{
            "message": "已开始猜数字游戏，目标数字是一个1到100之间的随机数！",
        })
    } else if game == "rps" {
        mu.Lock()
        c.JSON(http.StatusOK, gin.H{"message": "已开始石头剪刀布游戏！请提交你的选择。"})
        mu.Unlock()
    } else if game == "timer" {
        mu.Lock()
        gameStartTime = time.Now()
        mu.Unlock()
        c.JSON(http.StatusOK, gin.H{"message": "计时器游戏开始！请在 5 秒内提交！"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"message": "未识别的游戏"})
    }
}

// guessNumber 处理用户猜测
func guessNumber(c *gin.Context) {
    var guess struct {
        Number int `json:"number"`
    }

    if err := c.ShouldBindJSON(&guess); err == nil {
        mu.Lock()
        if guess.Number == targetNumber {
            mu.Unlock()
            c.JSON(http.StatusOK, gin.H{"message": "恭喜你，猜对了！"})
        } else if guess.Number < targetNumber {
            mu.Unlock()
            c.JSON(http.StatusOK, gin.H{"message": "猜小了，再试试！"})
        } else {
            mu.Unlock()
            c.JSON(http.StatusOK, gin.H{"message": "猜大了，再试试！"})
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"message": "请求数据不合法"})
    }
}

// playRockPaperScissors 处理石头剪刀布
func playRockPaperScissors(c *gin.Context) {
    var choice struct {
        PlayerChoice string `json:"choice"`
    }

    if err := c.ShouldBindJSON(&choice); err == nil {
        options := []string{"rock", "paper", "scissors"}
        computerChoice := options[rand.Intn(3)]

        var result string
        if choice.PlayerChoice == computerChoice {
            result = "平局！"
        } else if (choice.PlayerChoice == "rock" && computerChoice == "scissors") ||
            (choice.PlayerChoice == "scissors" && computerChoice == "paper") ||
            (choice.PlayerChoice == "paper" && computerChoice == "rock") {
            result = "你赢了！"
        } else {
            result = "你输了！"
        }

        c.JSON(http.StatusOK, gin.H{
            "computer": computerChoice,
            "result":   result,
        })
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"message": "请求数据不合法"})
    }
}

// startTimer 处理计时器游戏
func startTimer(c *gin.Context) {
    mu.Lock()
    elapsed := time.Since(gameStartTime).Seconds()
    mu.Unlock()

    if elapsed < 5 {
        c.JSON(http.StatusOK, gin.H{"message": "成功在 5 秒内提交！", "time": elapsed})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": "时间已过，请再试一次！", "time": elapsed})
    }
}

