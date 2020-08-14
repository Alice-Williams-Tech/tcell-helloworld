// package main is a "Hello World" demo with the GoLang terminal ui package tcell. 
package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	// Create new screen.
	scrn, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	
	// Initialize the screen for use.
	if err = scrn.Init(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	
	// Configure the screen's styling.
	scrn.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))

	// Go channel for quit call.
	quit := make(chan struct{})

	// Show the screen and start main loop.
	scrn.Show()
	go func() {
		for {
			// Clear the screen each iteration.
			scrn.Clear()

			// Print "Hello World" to the screen.
			displayString(scrn, "Hello World!")
		
            // Poll for event updates.
			ev := scrn.PollEvent()
			switch ev := ev.(type) {
			// Handle keypress events.
			case *tcell.EventKey:
				switch ev.Key() {
				// Handle quit keypresses: Ctrl-c, Escape, or Spacebar.
				case tcell.KeyCtrlC, tcell.KeyESC, tcell.KeyRune:
					close(quit)
					return
				}
			// Handle screen resizes.
			case *tcell.EventResize:
				scrn.Sync()
			}
		}
	}()

	// Wait for quit call from the application.
	<-quit

	// Finalize the screen releasing resources.
	scrn.Fini()
}

// displayString displays a String on a tcell screen.
func displayString(scrn tcell.Screen, strng string) {
	for i, c := range strng {
		scrn.SetContent(i, 0, c, nil, tcell.StyleDefault)
	}
}
