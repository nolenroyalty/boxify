package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var borders = map[string]lipgloss.Border{
	"normal":         lipgloss.NormalBorder(),
	"rounded":        lipgloss.RoundedBorder(),
	"block":          lipgloss.BlockBorder(),
	"outer-half":     lipgloss.OuterHalfBlockBorder(),
	"inner-half":     lipgloss.InnerHalfBlockBorder(),
	"thick":          lipgloss.ThickBorder(),
	"double":         lipgloss.DoubleBorder(),
	"hidden":         lipgloss.HiddenBorder(),
	"markdown":       lipgloss.MarkdownBorder(),
	"ascii":          lipgloss.ASCIIBorder(),
}

func main() {
	borderType := flag.String("border", "rounded", "border style: "+borderNames())
	paddingStr := flag.String("padding", "0,1", "padding: \"N\" for all sides, or \"Y,X\" (like CSS shorthand)")
	flag.Parse()

	padY, padX, err := parsePadding(*paddingStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid padding %q: %v\n", *paddingStr, err)
		os.Exit(1)
	}

	border, ok := borders[*borderType]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown border type: %q\navailable: %s\n", *borderType, borderNames())
		os.Exit(1)
	}

	var input []byte
	if flag.NArg() > 0 {
		input, err = os.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
			os.Exit(1)
		}
	} else {
		var err error
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
			os.Exit(1)
		}
	}

	text := strings.TrimRight(string(input), "\n")

	style := lipgloss.NewStyle().
		Border(border).
		Padding(padY, padX)

	fmt.Println(style.Render(text))
}

func parsePadding(s string) (y, x int, err error) {
	parts := strings.Split(s, ",")
	switch len(parts) {
	case 1:
		v, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return 0, 0, fmt.Errorf("expected integer, got %q", parts[0])
		}
		return v, v, nil
	case 2:
		y, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return 0, 0, fmt.Errorf("expected integer for Y, got %q", parts[0])
		}
		x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return 0, 0, fmt.Errorf("expected integer for X, got %q", parts[1])
		}
		return y, x, nil
	default:
		return 0, 0, fmt.Errorf("expected \"N\" or \"Y,X\", got %q", s)
	}
}

func borderNames() string {
	names := make([]string, 0, len(borders))
	for name := range borders {
		names = append(names, name)
	}
	return strings.Join(names, ", ")
}
