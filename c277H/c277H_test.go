package main

import (
	"image"
	"testing"
)

func BenchmarkBuild(t *testing.B) {
	data := make([]pixel, 500*400)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		build(data, 500, 400)
	}
}

func BenchmarkRun500x400(t *testing.B) {
	w, h, lim, wrk := 500, 400, 128, 6
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk)
		colorPx(data, img.Pix, max, wrk)
	}
}

func BenchmarkRun1080(t *testing.B) {
	w, h, lim, wrk := 1920, 1080, 128, 6
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk)
		colorPx(data, img.Pix, max, wrk)
	}
}

func BenchmarkRun4k(t *testing.B) {
	w, h, lim, wrk := 4096, 2160, 128, 6
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk)
		colorPx(data, img.Pix, max, wrk)
	}
}

func BenchmarkRun8k(t *testing.B) {
	w, h, lim, wrk := 7680, 4320, 128, 6
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	data := make([]pixel, w*h)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		build(data, w, h)
		max := calc(data, lim, wrk)
		colorPx(data, img.Pix, max, wrk)
	}
}
