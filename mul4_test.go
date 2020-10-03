package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	mgl "github.com/go-gl/mathgl/mgl32"
)

func floatClose32(a, b float32, ulps int32) bool {
	ai := int32(math.Float32bits(a))
	bi := int32(math.Float32bits(b))
	return ai>>31 == bi>>31 && ai-bi <= ulps && ai-bi >= -ulps
}

func genMatrix(r *rand.Rand) mgl.Mat4 {
	return mgl.Mat4{
		r.Float32(), r.Float32(), r.Float32(), r.Float32(),
		r.Float32(), r.Float32(), r.Float32(), r.Float32(),
		r.Float32(), r.Float32(), r.Float32(), r.Float32(),
		r.Float32(), r.Float32(), r.Float32(), r.Float32(),
	}
}

func TestAgainstMathGL(t *testing.T) {
	r := rand.New(rand.NewSource(0))

	for i := 0; i < 10000; i++ {
		a := genMatrix(r)
		b := genMatrix(r)

		mglRes := a.Mul4(b)
		sseRes := Mul4(a, b)

		for k := range mglRes {
			if !floatClose32(mglRes[k], sseRes[k], 3) {
				fmt.Print(a, "\n", b, "\n\n")
				fmt.Println(mglRes, "\n", sseRes)
				t.Fatalf("values non-close at index %d: %.20g != %.20g", k, mglRes[k], sseRes[k])
			}
		}
	}
}

func BenchmarkMathGL(b *testing.B) {
	id := mgl.Ident4()
	for i := 0; i < b.N; i++ {
		id.Mul4(id)
	}
}

func BenchmarkSSE(b *testing.B) {
	id := mgl.Ident4()
	for i := 0; i < b.N; i++ {
		Mul4(id, id)
	}
}
