# code-challenges
Code challenges: Leetcode, Daily Coding Challenges, etc.

```
package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var sections = []string{"$100", "$200", "$300", "$400", "$500", "$600", "$700", "$800", "$900", "$1000", "Bankrupt", "Lose a Turn"}

func main() {
	rand.Seed(time.Now().UnixNano())

	app := app.New()
	window := app.NewWindow("Wheel of Fortune")

	wheel := canvas.NewImageFromResource(theme.FyneLogo())

	spinButton := canvas.NewText("SPIN", theme.TextColor())
	spinButton.TextSize = 30
	spinButton.TextStyle = fyne.TextStyle{Bold: true}
	spinButton.Resize(spinButton.MinSize()) // resize to fit text

	spinButton.Move(fyne.NewPos(150, 300))

	spinButton.OnTapped = func(event *fyne.PointEvent) {
		result := spinWheel()
		fmt.Printf("You landed on: %s\n", result)
	}

	window.SetContent(container.NewVBox(
		wheel,
		spinButton,
	))

	window.Resize(fyne.NewSize(400, 400))
	window.ShowAndRun()
}

func spinWheel() string {
	// Simulate spinning by generating a random index after a few iterations
	for i := 0; i < 20; i++ {
		index := rand.Intn(len(sections))
		fmt.Printf("\rSpinning... %s", sections[index])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
	// Return the final result
	index := rand.Intn(len(sections))
	return sections[index]
}

```
