package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("Before...")
}

func teardown(){
	fmt.Println("Affter ...")
}

func TestAdd2(t *testing.T) {
	fmt.Println("testing")
}

func TestMain(m *testing.M)	{
	fmt.Println("main test")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}