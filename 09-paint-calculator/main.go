/*Prompt for length and width of ceiling.
constant: 1 gallon = 350 square ft.
Output: You will need to purchase 2 gallons of paint to cover 360 square feet
Need to round up to the next whole gallon of paint*/

package main

import (
	"fmt"
	"math"

	"github.com/milessabine/prompt"
)

const g = 350

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	p := prompt.New()
	l := p.ScanInt("What is the length of the ceiling?")
	w := p.ScanInt("What is the width of the ceiling?")
	a := l * w
	c := float64(a) / g
	//e := round(c)
	e := int(math.Ceil(c))

	fmt.Printf("You will need %d gallons of paint to cover %d square feet\n", e, a)

}

/*func promptAndScan(n string, scanner *bufio.Scanner) string {
	fmt.Print(n, "")
	scanner.Scan()
	return scanner.Text()
}

func promptAndScanInt(n string, scanner *bufio.Scanner) int {
	s := promptAndScan(n, scanner)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Please input valid number", s, err)
		return 0
	}
	return i

}*/

func round(c float64) int {
	i := int(c)         //c was 1.3 and now i =1
	t := i + 1          //t should now be equal to 2
	k := float64(t) - c //now k = 0.7
	if k == 1.0 {
		return i
	}
	return t

}

//func ParseFloat(s string, bitSize int) (f float64, err error)
/*func promptAndScanFloat64(n string, scanner *bufio.Scanner) float64 {
	s := promptAndScan(n, scanner)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Please input valid number", s, err)
		return 0.0
	}
	return f
}*/
