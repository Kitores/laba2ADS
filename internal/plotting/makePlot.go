package plotting

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
	"os"
	"strconv"
)

func CreateLineChartByPlotter(arrX []float64, arrY []int, typeOfTree string) {
	fileName := fmt.Sprintf("./internal/plotting/%s.png", typeOfTree)

	plt := plot.New()

	plt.Title.Text = typeOfTree
	plt.X.Label.Text = "keys"
	plt.Y.Label.Text = "height"

	// height
	points := make(plotter.XYs, len(arrX))
	for i := range arrX {
		points[i].X = arrX[i]
		points[i].Y = float64(arrY[i])
	}
	// log2n for avl
	log2nPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		log2nPoints[i].X = arrX[i]
		log2nPoints[i].Y = Log2n(arrX[i])
	}
	// logpPiN either for avl
	logHIPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		logHIPoints[i].X = arrX[i]
		logHIPoints[i].Y = LogHI(arrX[i])
	}
	// (log2n + 1) * 2 for rbt
	log2n1Points := make(plotter.XYs, len(arrX))
	for i := range arrX {
		log2n1Points[i].X = arrX[i]
		log2n1Points[i].Y = Log2n_1(arrX[i])
	}
	// log2n * k for BST
	log2nKPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		log2nKPoints[i].X = arrX[i]
		log2nKPoints[i].Y = Log2nK(arrX[i])
	}

	Plot, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}

	log2nPlot, err := plotter.NewLine(log2nPoints)
	if err != nil {
		panic(err)
	}

	logHIPlot, err := plotter.NewLine(logHIPoints)
	if err != nil {
		panic(err)
	}

	log2n1Plot, err := plotter.NewLine(log2n1Points)
	if err != nil {
		panic(err)
	}

	log2nKPlot, err := plotter.NewLine(log2nKPoints)
	if err != nil {
		panic(err)
	}

	Plot.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	log2nPlot.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	logHIPlot.LineStyle.Color = color.RGBA{R: 255, G: 127, B: 80, A: 255}
	log2n1Plot.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	log2nKPlot.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	if typeOfTree == "AVL" {
		plt.Add(Plot, log2nPlot, logHIPlot)
		plt.Legend.Add("Dependency Key/Height", Plot)
		plt.Legend.Add("log2n", log2nPlot)
		plt.Legend.Add("logφn", logHIPlot)
	} else if typeOfTree == " RBT" {
		plt.Add(Plot, log2nPlot)
		plt.Legend.Add("Dependency Key/Height", Plot)
		plt.Legend.Add("(log2n + 1) * 2", log2n1Plot)
	} else {
		plt.Add(Plot, log2nKPlot)
		plt.Legend.Add("Dependency Key/Height", Plot)
		plt.Legend.Add("log2n * k", log2nKPlot)
	}

	//plt.X.Tick.Marker = plot.DefaultTicks{}
	//plt.Y.Tick.Marker = plot.DefaultTicks{}

	//plt.X.Tick.Marker = plot.ConstantTicks(ticks)

	if err := plt.Save(8*vg.Inch, 4*vg.Inch, fileName); err != nil {
		panic(err)
	}
}
func Log2n_1(x float64) float64 {
	return math.Log2(x+1) * 2
}
func Log2nK(x float64) float64 {
	return math.Log2(x) * 1.9
}
func Log2n(x float64) float64 {
	return math.Log2(x)
}
func LogHI(x float64) float64 {
	return math.Log(x) / math.Log(1.618)
}

// generate data for line chart
func generateLineAverageItems(arrY []int, quantity int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < quantity; i++ {
		items = append(items, opts.LineData{Name: "heights", Value: arrY[i]})
	}
	//fmt.Println(items)
	return items
}

func CreateLineChart(arrX []float64, arrY []int, typeOfTree string) {
	fileName := fmt.Sprintf("./internal/plotting/%s.html", typeOfTree)

	// create a new line
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWonderland,
			//Theme: types.ThemeRoma,
			//Theme: types.ThemeRomantic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    typeOfTree,
			Subtitle: "This is fun to use!",
		}),
	)
	arrXInt := make([]int, len(arrX))
	for i := range arrX {
		arrXInt[i] = int(arrX[i])
	}

	tr := false
	// Put data into instance

	stringArrayX := make([]string, len(arrX))
	// Преобразуем каждый элемент массива целых чисел в строку
	for i, v := range arrXInt {
		stringArrayX[i] = strconv.Itoa(v) // Используем strconv.Itoa для преобразования
	}

	line.SetXAxis(stringArrayX).
		AddSeries("Dependency Height/Keys", generateLineAverageItems(arrY, len(arrY))).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: &tr, ConnectNulls: &tr}))
	f, _ := os.Create(fileName)
	_ = line.Render(f)
}
