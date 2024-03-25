package src

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Plot(title string, x string, y string, X []float64, Y []float64) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = x
	p.Y.Label.Text = y

	pts := make(plotter.XYs, 10)
	for i := range pts {
		pts[i].X = X[i]
		pts[i].Y = Y[i]
	}

	err := plotutil.AddLinePoints(p,
		x, pts)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "plots/"+title+y+".png"); err != nil {
		panic(err)
	}

}

func PlotSquares(title string, x string, y string, X []float64, Y []float64, squares Squares) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = x
	p.Y.Label.Text = y

	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = X[i]
		pts[i].Y = Y[i]
	}

	err := plotutil.AddLinePoints(p,
		x, pts)
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "assets/plots/"+title+y+".png"); err != nil {
		panic(err)
	}

}
