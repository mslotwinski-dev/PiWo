package labs

import (
	PiWo "PiWo/src"
	fmt "fmt"
	math "math"
)

func Laminarny() {
	PiWo.Start(3, "Laminarny przepływ cieczy")

	const R = 0.04
	const h = 1.4
	uh := 0.005 / math.Sqrt(3)
	ut := 0.2 / math.Sqrt(3)
	const rho_gl = 1261
	const rho_olej = 867
	const um = 0.1e-6
	ud := 0.01e-3 / math.Sqrt(3)
	// const ur = 0.005

	d := PiWo.SI([]float64{3.00, 3.01, 3.04, 2.96, 3.01, 3.01, 3.01, 3.01, 3.00, 3.01}, 1e-3)
	t_gl := []float64{30.38, 30.34, 30.47, 30.06, 30.03, 30.0, 29.94, 29.87, 29.81, 29.82}
	t_olej := []float64{10.06, 10.04, 10.03, 10.09, 10.16, 10.00, 9.97, 10.03, 10.03, 9.97}
	m_p := 1.3007
	m10 := 2.4297

	// Gęstość
	m := ((m10 - m_p) / 10) * 1e-3
	V := math.Pi * math.Pow(PiWo.Mean(d), 3) / 6
	rho := m / V

	uV := math.Pi * math.Pow(PiWo.Mean(d), 2) * ud / 2
	urho := PiWo.SumSquares([]float64{um / V, m * uV / math.Pow(V, 2)})

	// Prędkość graniczna

	v_gl := []float64{}
	v_olej := []float64{}
	for i := range t_gl {
		v_gl = append(v_gl, 0.7/t_gl[i])
		v_olej = append(v_olej, 1/t_olej[i])
	}

	ubv_gl := PiWo.SumSquares([]float64{uh / PiWo.Mean(t_gl), h * ut / math.Pow(PiWo.Mean(t_gl), 2)})

	uv_gl := PiWo.Uncertainty(v_gl, ubv_gl)

	ubv_olej := PiWo.SumSquares([]float64{uh / PiWo.Mean(t_olej), h * ut / math.Pow(PiWo.Mean(t_olej), 2)})
	uv_olej := PiWo.Uncertainty(v_olej, ubv_olej)

	// Lepkość

	r := PiWo.SI(d, 0.5)
	eta_gl := (2 * math.Pow(PiWo.Mean(r), 2) * g * (rho - rho_gl)) / (9 * PiWo.Mean(v_gl) * (1 + 2.4*PiWo.Mean(r)/R) * (1 + 3.1*PiWo.Mean(r)/h))
	eta_olej := (2 * math.Pow(PiWo.Mean(r), 2) * g * (rho - rho_olej)) / (9 * PiWo.Mean(v_olej) * (1 + 2.4*PiWo.Mean(r)/R) * (1 + 3.1*PiWo.Mean(r)/h))

	// O_O

	fmt.Printf("ϱ = %.2f (%0.2f) g/cm3\n", rho/1000, urho/1000)
	fmt.Printf("v = %0.2f (%0.2f) cm/s - gliceryna\n", 100*PiWo.Mean(v_gl), 100*uv_gl)
	fmt.Printf("v = %0.2f (%0.2f) cm/s - olej\n", 100*PiWo.Mean(v_olej), 100*uv_olej)
	fmt.Printf("η = %0.2f (%0.2f) Pas - gliceryna\n", eta_gl, EtaDerivatives(PiWo.Mean(r), rho, PiWo.Mean(v_gl), rho_gl, ud/2, urho, uv_gl))
	fmt.Printf("η = %0.2f (%0.2f) Pas - olej\n", eta_olej, EtaDerivatives(PiWo.Mean(r), rho, PiWo.Mean(v_olej), rho_olej, ud/2, urho, uv_olej))
}

func EtaDerivatives(r, rho, v, c, ur, urho, uv float64) float64 {
	const R = 0.04
	const h = 1.4
	uR := 3e-4 / math.Sqrt(3)
	uh := 0.005 / math.Sqrt(3)

	// Obliczenie pochodnych funkcji eta po poszczególnych zmiennych
	dEtaDr := ur * (4 * math.Pow(r, 2) * 9.8 * (rho - c)) / (9 * v * (1 + 2.4*r/R) * (1 + 3.1*r/h))
	dEtaDrho := urho * (2 * math.Pow(r, 2) * 9.8 * (1 - R/(1+2.4*r/R)) / (9 * v * (1 + 2.4*r/R) * (1 + 3.1*r/h)))
	dEtaDv := uv * (-2 * math.Pow(r, 2) * 9.8 * (rho - c)) / (9 * math.Pow(v, 2) * (1 + 2.4*r/R) * (1 + 3.1*r/h))
	dEtaDR := uR * (-2.4 * math.Pow(r, 3) * 9.8 * (rho - c)) / (9 * v * math.Pow((1+2.4*r/R), 2) * (1 + 3.1*r/h))
	dEtaDh := uh * (-3.1 * math.Pow(r, 2) * 9.8 * (rho - c)) / (9 * v * (1 + 2.4*r/R) * math.Pow((1+3.1*r/h), 2))

	// Suma kwadratów pochodnych
	sumOfSquares := math.Pow(dEtaDr, 2) + math.Pow(dEtaDrho, 2) + math.Pow(dEtaDv, 2) + math.Pow(dEtaDR, 2) + math.Pow(dEtaDh, 2)

	return math.Sqrt(sumOfSquares)
}
