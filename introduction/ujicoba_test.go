package introduction

import "testing"

var (
	cube          Cube    = Cube{4}
	volumeTest    float64 = 64
	areaTest      float64 = 96
	perimeterTest float64 = 48
)

func TestCubeVolume(t *testing.T) {
	t.Logf("Volume : %.2f", cube.Volume())

	if cube.Volume() != volumeTest {
		t.Errorf("Wrong ! Value Volume != %.2f", volumeTest)
	}
}

func TestCubeArea(t *testing.T) {
	t.Logf("Surface Area Of Cube : %.2f", cube.SurfaceArea())

	if cube.SurfaceArea() != areaTest {
		t.Errorf("Wrong ! Value Surface Area Of Cube != %.2f", areaTest)
	}
}

func TestCubePerimeter(t *testing.T) {
	t.Logf("Perimeter : %.2f", cube.Perimeter())

	if cube.Perimeter() != perimeterTest {
		t.Errorf("Wrong ! Value Perimeter != %.2f", perimeterTest)
	}
}

func BenchmarkVolume(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cube.Volume()
	}
}

func BenchmarkSurfaceArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cube.SurfaceArea()
	}
}
