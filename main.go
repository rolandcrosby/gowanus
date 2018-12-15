// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/fogleman/ln/ln"
)

func querySelector(sel string) js.Value {
	return js.Global().Get("document").Call("querySelector", sel)
}

func addClass(el js.Value, className string) {
	el.Get("classList").Call("add", className)
}

func removeClass(el js.Value, className string) {
	el.Get("classList").Call("remove", className)
}

func doSomething(_ []js.Value) {
	el := querySelector("#out")
	w := el.Get("offsetWidth").Float()
	h := el.Get("offsetHeight").Float()
	addClass(el, "rendering")
	el.Set("innerHTML", "<h1>rendering...</h1>")
	el.Set("innerHTML", functionPaths(w, h).ToSVG(w, h))
	removeClass(el, "rendering")
}

func registerCallbacks() {
	js.Global().Set("doSomething", js.NewCallback(doSomething))
}

func functionPaths(width, height float64) ln.Paths {
	scene := ln.Scene{}
	box := ln.Box{ln.Vector{-2, -2, -4}, ln.Vector{2, 2, 2}}
	scene.Add(ln.NewFunction(func(x, y float64) float64 {
		return -1 / (x*x + y*y)
	}, box, ln.Below))
	eye := ln.Vector{3, 0, 3}
	center := ln.Vector{1.1, 0, 0}
	up := ln.Vector{0, 0, 1}
	return scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
}

func cubePaths(width, height float64) ln.Paths {
	// create a scene and add a single cube
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// define camera parameters
	eye := ln.Vector{4, 3, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}     // up direction

	// define rendering parameters
	fovy := 50.0 // vertical field of view, degrees
	znear := 0.1 // near z plane
	zfar := 10.0 // far z plane
	step := 0.01 // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	return scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	fmt.Println("we've got callbacks")
	button := querySelector("#button")
	button.Set("disabled", false)
	<-c
}
