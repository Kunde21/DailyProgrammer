package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBase(t *testing.T) {
	threes(929, 0, nil)
	fmt.Println("fin")
	wg.Wait()
	fmt.Println("waited")
}

func TestBigger(t *testing.T) {
	//	i, _ := big.NewInt(0).SetString("18446744073709551614", 10)
	threes(3, 0, nil)
	wg.Wait()

	fmt.Println(33, "Start")
	threes(33, 0, nil)
	wg.Wait()

	fmt.Println(929, "Start")
	threes(929, 0, nil)
	wg.Wait()

	fmt.Println("Starting")
	tmp := os.Stdout
	f, err := os.Open(os.DevNull)
	if err == nil {
		os.Stdout = f
	}
	threes(18446744073709551614, 0, nil)
	wg.Wait()
	os.Stdout = tmp
	fmt.Println("Done")
}
