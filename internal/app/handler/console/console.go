package console

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func (c *Console) getSelection(min int, max int) int {
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Your selection [%d-%d]: ", min, max)

		text, err := reader.ReadString('\n')
		if err != nil {
			return 0
		}
		text = strings.TrimRight(text, c.newline)

		num, err := strconv.Atoi(text)
		if err != nil || num < min || num > max {
			fmt.Printf("Please select between %d and %d\n", min, max)
			time.Sleep(300 * time.Millisecond)
		} else {
			return num
		}
	}
}

func (c *Console) clearScreen() {
	switch c.environment {
	case "linux":
		{
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	case "windows":
		{
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
}
