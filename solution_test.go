package square

import (
	"testing"
)

const (
	ExpectedAreaTriangle float64 = 2.29063719300984
	ExpectedAreaSquare   float64 = 16
	ExpectedAreaCircle   float64 = 180.5045665888
)

func TestCalcSquare(t *testing.T) {
	areaTriangle := CalcSquare(2.3, SidesTriangle)
	if areaTriangle != ExpectedAreaTriangle {
		t.Errorf("areaTriangle %f expected %f", areaTriangle, ExpectedAreaTriangle)
	}

	areaSquare := CalcSquare(4, SidesSquare)
	if areaSquare != ExpectedAreaSquare {
		t.Errorf("areaSquare %f expected %f", areaSquare, ExpectedAreaSquare)
	}

	areaCircle := CalcSquare(7.58, SidesCircle)
	if areaCircle != ExpectedAreaCircle {
		t.Errorf("areaCircle %f expected %f", areaCircle, ExpectedAreaCircle)
	}

	wrongSidesNum := CalcSquare(2.3, 5)
	if wrongSidesNum != 0 {
		t.Errorf("wrongSidesNum %f expected %d", wrongSidesNum, 0)
	}

	wrongSideLen := CalcSquare(0, 5)
	if wrongSideLen != 0 {
		t.Errorf("wrongSideLen %f expected %d", wrongSideLen, 0)
	}
}
