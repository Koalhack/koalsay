package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func replaceTabToSpace(lines []string) []string {
	var newLines []string
	for _, line := range lines {
		line := strings.Replace(line, "\t", "  ", -1)
		newLines = append(newLines, line)
	}

	return newLines
}

func calculateMaxWidth(lines []string) int {
	var width = 0
	for _, line := range lines {
		var lineLength = len(line)
		if lineLength > width {
			width = lineLength
		}
	}

	return width
}

func normalizelineWidth(lines []string, maxWidth int) []string {
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, string(line)+strings.Repeat(" ", maxWidth-len(line)))
	}

	return newLines
}

func buildBalloon(lines []string, maxWidth int) string {
	var borders = []string{"/", "\\", "\\", "/", "|", "<", ">", "-", "_"}
	var borderTop = " " + strings.Repeat(borders[8], maxWidth+2)
	var borderBottom = " " + strings.Repeat(borders[7], maxWidth+2)

	var balloon []string
	var count = len(lines)

	balloon = append(balloon, borderTop)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		balloon = append(balloon, s)
	} else {
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		balloon = append(balloon, s)
		for i := 1; i < count; i++ {
			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			balloon = append(balloon, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[2], lines[0], borders[3])
		balloon = append(balloon, s)

	}
	balloon = append(balloon, borderBottom)

	return strings.Join(balloon, "\n")
}

func main() {

	//INFO: Get data from Stdin

	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		var line, _, err = reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(line))
	}

	//INFO: Display final Result

	lines = replaceTabToSpace(lines)
	maxWidth := calculateMaxWidth(lines)
	messages := normalizelineWidth(lines, maxWidth)
	balloon := buildBalloon(messages, maxWidth)

	fmt.Println(balloon)
}
