package main
import "fmt"

func main(){
	var numbers [5] int
	var maximumElement int
	var j int

	for i :=0; i < len(numbers); i++ {
		fmt.Println("Insert number:")
		fmt.Scan(&numbers[i])

		for _, val := range numbers {
			if val > j {
				j=val
				maximumElement = j
			}
		}
	}
	fmt.Println("Array:", numbers)
	fmt.Println("Maximum/largest element is", maximumElement)
}
