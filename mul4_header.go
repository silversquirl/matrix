// +build amd64,!nosse

package matrix

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

func Mul4(a, b mgl.Mat4) mgl.Mat4
