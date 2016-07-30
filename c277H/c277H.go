package main

import (
	"image"
	"image/jpeg"
	"log"
	"math/cmplx"
	"os"
	"sync"
)

type RGBA []uint8

type pixel struct {
	val  complex128
	iter int
}

func main() {
	var (
		lim     = 256
		w, h    = 7680, 4320
		workers = 6
		data    = make([]pixel, w*h)
		img     = image.NewRGBA(image.Rect(0, 0, w, h))
	)
	build(data, w, h)
	lim = calc(data, lim, workers)
	colorPx(data, img.Pix, lim, workers)
	err := writeImg(img, "julia.jpeg")
	if err != nil {
		log.Fatal(err)
	}
}

func build(data []pixel, width, height int) {
	relStep, imgStep := 2.0/float64(width-1), 2.0/float64(height-1)
	var img float64 = 1
	for row := 0; row < height; row++ {
		var rel float64 = -1
		for col := 0; col < width; col++ {
			data[row*width+col].val = complex(rel, img)
			rel += relStep
		}
		img -= imgStep
	}
}

func calc(data []pixel, lim, wrk int) (max int) {
	wg := new(sync.WaitGroup)
	wg.Add(wrk)
	ch := make(chan int)
	ln := len(data) / wrk
	for w := 0; w < wrk-1; w++ {
		go func(w int, ch chan int) {
			max := 0
			for i := range data[w : w+ln] {
				data[w+i].julia(lim)
				if data[w+i].iter > max {
					max = data[w+i].iter
				}
			}
			ch <- max
		}(w*ln, ch)
	}
	go func(ch chan int) {
		w, max := (wrk-1)*ln, 0
		for i := range data[w:] {
			data[w+i].julia(lim)
			if data[w+i].iter > max {
				max = data[w+i].iter
			}
		}
		ch <- max
	}(ch)
	max = <-ch
	for i := 0; i < wrk-1; i++ {
		t := <-ch
		if t > max {
			max = t
		}
	}
	close(ch)
	return max
}

func (px *pixel) julia(lim int) {
	for cmplx.Abs(px.val) < 2.0 && px.iter < lim {
		px.val = px.val*px.val - complex(.221, .713)
		px.iter++
	}
}

func colorPx(data []pixel, clrs []uint8, lim, wrk int) {
	wg := new(sync.WaitGroup)
	wg.Add(wrk)
	ln := len(data) / wrk
	for w := 0; w < wrk-1; w++ {
		go func(w int) {
			for i := range data[w : w+ln] {
				copy(clrs[(w+i)*4:], convert(data[w+i].iter, lim))
			}
			wg.Done()
		}(w * ln)
	}
	go func() {
		w := (wrk - 1) * ln
		for i := range data[w:] {
			copy(clrs[(w+i)*4:], convert(data[w+i].iter, lim))
		}
		wg.Done()
	}()
	wg.Wait()
}

func convert(iter, lim int) (px RGBA) {
	var tmp uint8
	c := float64(iter) / float64(lim)
	switch {
	case c <= 0:
		px = RGBA{0, 0, 0, 255}
	case c <= 0.1:
		tmp = uint8(c / 0.1 * 255)
		px = RGBA{tmp, 0, 0, 255}
	case c <= 0.25:
		tmp = uint8(c / 0.25 * 255)
		px = RGBA{255, tmp, 0, 255}
	case c <= 0.5:
		tmp = uint8(1 - c/0.5*255)
		px = RGBA{tmp, 255, 0, 255}
	case c <= 0.75:
		tmp = uint8(c / 0.75 * 255)
		px = RGBA{0, 255 - tmp, tmp, 255}
	default:
		tmp = uint8(c * 255)
		px = RGBA{tmp, tmp, 255, 255}
	}
	return px
}

func writeImg(img image.Image, fn string) error {
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	return nil
}
