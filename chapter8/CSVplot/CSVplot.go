// Чтение CSV-файлов
// CSV-файлы — это простые форматированные текстовые файлы. В этом разделе
// показано, как прочитать текстовый файл, в котором находятся координаты точек
// на плоскости, — другими словами, в каждой строке содержится пара координат.
// Кроме того, мы также воспользуемся внешней библиотекой Go, которая называется
// Glot, — она поможет нам построить график по точкам, координаты которых будут
// прочитаны из CSV-файла. Обратите внимание, что в Glot используется Gnuplot,
// поэтому, чтобы использовать Glot, придется установить библиотеку Gnuplot на
// ваш компьютер UNIX.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Arafatk/glot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a data file")
		return
	}

	file := os.Args[1]
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("cannot stat", file)
		return
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file %s\n", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	xP := []float64{}
	yP := []float64{}

	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}

	points := [][]float64{}
	points = append(points, xP)
	points = append(points, yP)
	fmt.Println(points)

	dimensions := 2
	persist := true
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	plot.SetTitle("Using Glot with CSV data")
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	style := "circle"
	plot.AddPointGroup("MyPlot", style, points)
	plot.SavePlot("output.png")
}
