package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Speed int

type FanSpeed struct {
	Speed Speed
}

type Money struct {
	Length float64
}

type Memory struct {
	Count      int
	MemorySize []int
}

type Computer struct {
	SystemName string
	UseNumber  int
	Memory     Memory
	Fan        map[string]FanSpeed
	Money      Money
}

func (s *Computer) Clone() *Computer {
	resume := *s
	return &resume
}

func (s *Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, s); err != nil {
		panic(err.Error())
	}
	return pc
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main() {
	a := Computer{
		SystemName: "Windows10",
		UseNumber:  10,
		Memory: Memory{
			Count:      10,
			MemorySize: []int{},
		},
		Fan: map[string]FanSpeed{
			"Fan1": {Speed: 10},
			"Fan2": {Speed: 10},
		},
		Money: Money{
			Length: 100.00,
		},
	}

	fmt.Println("a: ", a)
	b := a.BackUp()
	c := a.Clone()
	a.SystemName = "Linux"
	fmt.Println("a.SystemName: ", a.SystemName)
	fmt.Println("b.SystemName: ", b.SystemName)
	fmt.Println("b: ", b)
	fmt.Println("c.SystemName: ", c.SystemName)
	fmt.Println("c: ", c)
}
