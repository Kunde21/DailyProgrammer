package main

import (
	"fmt"
	"log"
	"math"
)

type animal struct {
	XYpair
	speed float64
}
type XYpair struct {
	x, y float64
}

func (c XYpair) Dot(d XYpair) float64 {
	return c.x*d.x + c.y*d.y
}

func (c XYpair) Norm() float64 {
	return math.Hypot(c.x, c.y)
}

func (c XYpair) Diff(d XYpair) XYpair {
	return XYpair{c.x - d.x, c.y - d.y}
}

func main() {
	var pupper, gopher animal
	_, err := fmt.Scanln(&pupper.x, &pupper.y, &pupper.speed)
	_, errG := fmt.Scanln(&gopher.x, &gopher.y, &gopher.speed)
	if err != nil || errG != nil {
		log.Fatal("Animal parsing failed:", err, errG)
	}

	reps := 0
	_, err = fmt.Scanln(&reps)
	if err != nil {
		log.Fatal("Hole count parsing failed:", err)
	}
	gHoles := make([]XYpair, reps)
	for i := range gHoles {
		_, err = fmt.Scanln(&gHoles[i].x, &gHoles[i].y)
		if err != nil {
			log.Fatalf("Hole %v parsing failed: %s\n", i, err)
		}
	}

	catch := intercept(pupper, gopher, nearest(gopher, gHoles))
	if catch == nil {
		fmt.Println("Gopher got away.")
	}
	fmt.Println(catch.x, catch.y)
}

func nearest(gopher animal, holeList []XYpair) (tgt XYpair) {
	var dist float64
	for i, h := range holeList {
		x, y := gopher.x-h.x, gopher.y-h.y
		tmp := math.Sqrt(float64(x*x + y*y))
		if i == 0 || tmp < dist {
			dist, tgt = tmp, h
		}
	}
	return tgt
}

func intercept(p, g animal, tgt XYpair) (xy *XYpair) {
	pg, tg := p.Diff(g.XYpair), tgt.Diff(g.XYpair)
	gVel := XYpair{g.speed * tg.x / g.Norm(), g.speed * tg.y / g.Norm()}
	arrTime := tgt.Diff(g.XYpair).Norm() / g.speed

	a := g.speed*g.speed - p.speed*p.speed
	b := -2 * gVel.Dot(pg)
	c := pg.Dot(pg)
	time1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	time2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	catchTime := time1
	if time2 > catchTime {
		catchTime = time2
	}

	if catchTime < 0 || catchTime > arrTime {
		return nil
	}
	return &XYpair{g.x + gVel.x*catchTime, g.y + gVel.y*catchTime}
}
