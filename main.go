package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"mazegen/maze"
	"os"
	"strconv"
	"strings"
)

var (
	wallColor     = color.RGBA{165, 42, 42, 255}
	cellColor     = color.RGBA{255, 255, 255, 255}
	startColor    = color.RGBA{0, 0, 255, 255}
	endColor      = color.RGBA{0, 255, 0, 255}
	solutionColor = color.RGBA{32, 32, 32, 255}
)

func main() {
	// Setup and parse optional CLI arguments
	path := flag.String("path", "simple", "path length: simple, convoluted")
	wall_size := flag.Int("wall", 4, "wall thickness")
	cell_size := flag.Int("cell", 10, "cell size")
	solved := flag.Bool("solve", false, "generate solution along with maze")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s [flags] <columns> <rows> <output_file>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	var cols, rows int
	filename := ""
	algo := maze.AlgoKruskal

	// Verify path type
	switch *path {
	case "simple":
		algo = maze.AlgoKruskal
	case "convoluted":
		algo = maze.AlgoRandWalk
	default:
		debug("None of the options match for path length.")
		helpMsgExit()
	}

	// Parse positional CLI arguments
	switch flag.NArg() {
	case 3:
		cols = int(parseUintArg("columns", 2, 1<<16, flag.Arg(0)))
		rows = int(parseUintArg("rows", 2, 1<<16, flag.Arg(1)))
		// Strip .png extension (if there)
		filename, _ = strings.CutSuffix(flag.Arg(2), ".png")

	default:
		debug("Wrong number of positional arguments.")
		helpMsgExit()
	}

	debug("Generating maze...")
	// ----------------------------------------------------
	my_maze := maze.GenerateMaze(
		algo,
		int(cols), int(rows),
		maze.TopLeft, maze.BottomRight,
	)

	debug("Rendering image(s)...")
	// ----------------------------------------------------
	img := my_maze.ToImage(
		*cell_size, *wall_size, false,
		wallColor, cellColor,
		startColor, endColor, solutionColor,
	)

	output, err := os.Create(filename + ".png")
	if err == nil {
		png.Encode(output, img)
	} else {
		debug("Error creating file: %v", err.Error())
		os.Exit(1)
	}

	if !*solved {
		return
	}

	// Generate solution image if requested.
	// ----------------------------------------------------
	sol_img := my_maze.ToImage(
		*cell_size, *wall_size, true,
		wallColor, cellColor,
		startColor, endColor, solutionColor,
	)

	sol_output, err := os.Create(filename + "-solved.png")
	if err == nil {
		png.Encode(sol_output, sol_img)
	} else {
		debug("Error creating file: %v", err.Error())
		os.Exit(1)
	}

	debug("Done.")
}

// Parse as unsigned integer and exit after printing
// error message along with flags help message if any error occurs.
func parseUintArg(name string, min, max uint, raw_arg string) uint {
	ret, ok := strconv.ParseInt(raw_arg, 10, 32)
	val := uint(ret)

	if ok != nil || val < min || val > max {
		debug(`invalid value "%v" for argument %v.`, raw_arg, name)
		helpMsgExit()
	}

	return val
}

func helpMsgExit() {
	flag.Usage()
	os.Exit(2)
}

func debug(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
