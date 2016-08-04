package main

import (
	"image"
	"testing"
)

var (
	lim = 256
	wrk = 6
)

/* Quick way to generate multiple fractals
func TestBuild(t *testing.T) {
	var (
		lim  = 256
		w, h = 7680, 4320
		// c       = complex(.755, -.072)
		fname   = "julia.jpeg"
		workers = runtime.NumCPU()
		data    = make([]pixel, w*h)
		img     = image.NewRGBA(image.Rect(0, 0, w, h))
	)
	for _, c := range []complex128{
		.221 + .713i,
		.545 - .507i,
		.545 - .508i,
		-.384 - .259i,
		.552 - .514i,
		.713 - .252i,
		.747 - .114i,
		.759 - .079i,
		.763 - .072i,
	} {
		fname = fmt.Sprintf("julia%v2.jpeg", -c)
		build(data, w, h)
		lim = calc(data, lim, workers, c)
		colorPx(data, img.Pix, lim, workers)
		err := writeImg(img, fname)
		if err != nil {
			log.Fatal(err)
		}
	}
} //*/

func BenchmarkRun500x400(t *testing.B) {
	w, h := 500, 400
	c := complex(.221, .713)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk, c)
		colorPx(data, img.Pix, max, wrk)
	}
}

func BenchmarkRun1080(t *testing.B) {
	w, h := 1920, 1080
	c := complex(.221, .713)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk, c)
		colorPx(data, img.Pix, max, wrk)
	}
}

func BenchmarkRun4k(t *testing.B) {
	w, h := 4096, 2160
	c := complex(.221, .713)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk, c)
		colorPx(data, img.Pix, max, wrk)
	}
}

func BenchmarkRun8k(t *testing.B) {
	w, h := 7680, 4320
	c := complex(.221, .713)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk, c)
		colorPx(data, img.Pix, max, wrk)
	}
}
