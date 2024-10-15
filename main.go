package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	currentGame string
)

func main() {
	a := app.New()
	w := a.NewWindow("小游戏合集")

	currentGame = "guessNumber" // 初始游戏设置为猜数字

	// 游戏选择界面
	gameSelect := container.NewVBox(
		widget.NewLabel("选择游戏:"),
		widget.NewButton("猜数字游戏", func() {
			startGuessNumberGame(w)
		}),
		widget.NewButton("猜拳游戏", func() {
			startRockPaperScissorsGame(w)
		}),
		widget.NewButton("打地鼠游戏", func() {
			startWhackAMoleGame(w)
		}),
	)

	w.SetContent(gameSelect)
	w.ShowAndRun()
}

// 猜数字游戏
func startGuessNumberGame(w fyne.Window) {
	secretNumber := rand.Intn(100) + 1
	var guess int

	// 输入框
	input := widget.NewEntry()
	input.SetPlaceHolder("请输入1到100之间的数字")

	// 按钮
	submit := widget.NewButton("提交", func() {
		fmt.Sscanf(input.Text, "%d", &guess)
		if guess < secretNumber {
			dialog.NewInformation("结果", "太小了，再试一次！", w).Show()
		} else if guess > secretNumber {
			dialog.NewInformation("结果", "太大了，再试一次！", w).Show()
		} else {
			celebrate("猜数字游戏成功！", w)
		}
	})

	// 游戏界面
	gameContainer := container.NewVBox(
		widget.NewLabel("猜数字游戏：输入1到100之间的数字"),
		input,
		submit,
	)

	w.SetContent(gameContainer)
}

// 猜拳游戏
func startRockPaperScissorsGame(w fyne.Window) {
	choices := []string{"剪刀", "石头", "布"}
	rand.Seed(time.Now().UnixNano())

	// 输入框
	input := widget.NewSelect([]string{"剪刀", "石头", "布"}, func(choice string) {
		computerChoice := rand.Intn(3)
		if choice == choices[computerChoice] {
			dialog.NewInformation("结果", "平局！", w).Show()
		} else if (choice == "剪刀" && computerChoice == 2) || (choice == "石头" && computerChoice == 0) || (choice == "布" && computerChoice == 1) {
			celebrate("你赢了！", w)
		} else {
			dialog.NewInformation("结果", "你输了！", w).Show()
		}
	})

	// 按钮
	submit := widget.NewButton("提交", func() {
		input.OnChanged(input.Selected)
	})

	// 游戏界面
	gameContainer := container.NewVBox(
		widget.NewLabel("猜拳游戏：选择剪刀、石头或布"),
		input,
		submit,
	)

	w.SetContent(gameContainer)
}

// 打地鼠游戏
func startWhackAMoleGame(w fyne.Window) {
	moles := 0
	hits := 0
	rounds := 10

	for i := 0; i < rounds; i++ {
		moles = rand.Intn(3) + 1
		dialog.NewInformation("回合", fmt.Sprintf("回合 %d：击打 %d 个地鼠！", i+1, moles), w).Show()
		hits += moles
	}

	celebrate(fmt.Sprintf("你总共击打了 %d 个地鼠！", hits), w)
}

// 庆祝效果
func celebrate(message string, w fyne.Window) {
	dialog.NewInformation("庆祝", message, w).Show()
	// 重新开始按钮
	restart := widget.NewButton("重新开始", func() {
		currentGame = "guessNumber"
		startGameSelection(w)
	})

	// 游戏结束界面
	celebrateContainer := container.NewVBox(
		widget.NewLabel(message),
		restart,
	)

	w.SetContent(celebrateContainer)
}

// 游戏选择界面
func startGameSelection(w fyne.Window) {
	// 游戏选择界面
	gameSelect := container.NewVBox(
		widget.NewLabel("选择游戏:"),
		widget.NewButton("猜数字游戏", func() {
			startGuessNumberGame(w)
		}),
		widget.NewButton("猜拳游戏", func() {
			startRockPaperScissorsGame(w)
		}),
		widget.NewButton("打地鼠游戏", func() {
			startWhackAMoleGame(w)
		}),
	)

	w.SetContent(gameSelect)
}

