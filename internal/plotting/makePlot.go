package plotting

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func CreateLineChartByPlotter(arrX []float64, arrY, arrYWorst, arrYBest, arrYAlmost []int64, quantity int, SortName string) {
	fileName := fmt.Sprintf("./internal/plotting/plots/pngPlots/%s.png", SortName)

	plt := plot.New()

	plt.Title.Text = SortName
	plt.X.Label.Text = "sizes"
	plt.Y.Label.Text = "nanoseconds"

	//  random arr
	randomCasePoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		randomCasePoints[i].X = float64(arrX[i])
		randomCasePoints[i].Y = float64(arrY[i])
	}
	//	reversed arr
	reversedSortedPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		reversedSortedPoints[i].X = float64(arrX[i])
		reversedSortedPoints[i].Y = float64(arrYWorst[i])
	}
	//	sorted arr
	alreadySortedPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		alreadySortedPoints[i].X = float64(arrX[i])
		alreadySortedPoints[i].Y = float64(arrYBest[i])
	}
	//	almost sorted arr
	almostSortedPoints := make(plotter.XYs, len(arrX))
	for i := range arrX {
		almostSortedPoints[i].X = float64(arrX[i])
		almostSortedPoints[i].Y = float64(arrYAlmost[i])
	}
	averagePlot, err := plotter.NewLine(randomCasePoints)
	if err != nil {
		panic(err)
	}
	worstPlot, err := plotter.NewLine(reversedSortedPoints)
	if err != nil {
		panic(err)
	}
	BestPlot, err := plotter.NewLine(alreadySortedPoints)
	if err != nil {
		panic(err)
	}
	AlmostPlot, err := plotter.NewLine(almostSortedPoints)
	if err != nil {
		panic(err)
	}

	averagePlot.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	worstPlot.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	BestPlot.LineStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	AlmostPlot.LineStyle.Color = color.RGBA{R: 255, G: 127, B: 80, A: 255}
	plt.Add(averagePlot, worstPlot, BestPlot, AlmostPlot)

	plt.Legend.Add("Random Case", averagePlot)
	plt.Legend.Add("Reversed", worstPlot)
	plt.Legend.Add("Sorted", BestPlot)
	plt.Legend.Add("Almost sorted", AlmostPlot)

	plt.X.Tick.Marker = plot.DefaultTicks{}
	plt.Y.Tick.Marker = plot.DefaultTicks{}

	//plt.X.Tick.Marker = plot.ConstantTicks(ticks)

	if err := plt.Save(10*vg.Inch, 5*vg.Inch, fileName); err != nil {
		panic(err)
	}
}
