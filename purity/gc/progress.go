package gc

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// newProgressBar returns a channel to a progress bar that animates until the channel is closed.
// The caller should push ints onto the channel indicating incremental progress.
func newProgressBar(message string, total int) chan<- int {
	bar := &progressBar{
		message: message,
		total:   total,
		ch:      make(chan int, 10),
	}
	if cols, _, err := terminal.GetSize(int(os.Stderr.Fd())); err == nil {
		bar.cols = cols
	}
	go bar.Animate()
	return bar.ch
}

type progressBar struct {
	message              string
	total, current, cols int
	ch                   chan int
}

func (bar *progressBar) Animate() {
	for inc := range bar.ch {
		bar.current += inc
		proportion := float64(bar.current) / float64(bar.total)
		percentage := 100.0 * proportion
		if bar.cols == 0 {
			fmt.Fprintf(os.Stderr, "%0.1%%\n", percentage)
		} else {
			bar.render(proportion, percentage)
		}
	}
}

func (bar *progressBar) render(proportion, percentage float64) {
	totalCols := bar.cols - 40 // Pretty arbitrary amount of overhead to make sure we have space.
	currentPos := int(proportion * float64(totalCols))
	if currentPos > totalCols {
		currentPos = totalCols
	}
	before := strings.Repeat("=", currentPos)
	after := strings.Repeat(" ", totalCols-currentPos)
	fmt.Fprintf(os.Stderr, "${RESETLN}${BOLD_WHITE}%s: ${GREY}[%s>%s] ${BOLD_WHITE}%0.1f%%${RESET}", bar.message, before, after, percentage)
}
