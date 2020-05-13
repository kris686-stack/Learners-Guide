// MockSample1 project main.go

package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func array() {
	fmt.Println("Array of data")
	var a, b [2]string
	a[0] = "Golang"
	a[1] = "JVT"
	fmt.Print(a)
	for i := 0; i < len(b); i++ {
		b[i] = a[i]
	}
	var ab [2][3]int = [2][3]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(ab, "is the 2d array")
}
func slice() {
	fmt.Println("Slice Of Data")
	var s []int
	// give nil o/p
	fmt.Println(s)
	s = make([]int, 12) // slice reference an array so make returns an address to s by calling s it dereference the value
	fmt.Println("Slice with initialization", s)
	s1 := []int{1, 2, 3, 5, 6, 7, 6}
	/*s1[0] = 12 this will give panic
	s1[1] = 34
	s1[4] = 345
	s1[5] = 6
	s1[3] = 78*/
	fmt.Println(s1)
	for i, e := range s1 {
		fmt.Println("Dereferenced an unsorted slice using range", i, e)
	}
	fmt.Println("sorting S1")
	sort.Ints(s1)
	fmt.Println(s1)
}
func mapSam() {
	var m map[int]string
	fmt.Println("nil map", m)
	m1 := map[int]string{1: "Sony", 2: "surya", 3: "Vasanthi", 4: "Lekshmi", 5: "Aiswarya", 6: "Krishna"}
	fmt.Println("map using make", m1)
	v, ok := m1[3]
	fmt.Println(v, ok)
	delete(m1, 6)
	fmt.Println("jvt gobatch 1 ", m1)
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}

type data struct {
	Name           string      `json:"name"`
	Age            int         `json:"age"`
	Job            string      `json:"job"`
	Qualifications []education `json:"qualifications"`
}
type education struct {
	Course string `json:"course"`
	Type   string `json:"type"`
}
type sampleOne struct {
	name string
	age  int
}

type details interface {
	personDetails()
}

func (sam1 sampleOne) personDetails() {
	fmt.Println("partial data of person", sam1)
}
func callDetails(detail []details) {
	for _, v := range detail {
		v.personDetails()

	}
}

func main() {
	//array()
	//slice()
	//mapSam()
	//var sam1 sampleOne = sampleOne{"ramu", 30}
	//	var sam2 sampleOne = sampleOne{"reghu", 32}
	//var detail []details = []details{sam1, sam2}
	//callDetails(detail)
	jsdata := []byte(`
	{
		"name":"Krishna",
		"age":30,
		"job":"Go Coder",
		"qualifications":[{
			"course":"BA",
			"type":"PG Diploma"
			},
			{
				"course":"B TECH",
				"Type":"UG"
		    }
			]
	}
				
	`)
	var d data
	err := json.Unmarshal(jsdata, &d)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(d)
	js, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(js))

}
