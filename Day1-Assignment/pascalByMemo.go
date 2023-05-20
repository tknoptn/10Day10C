package main

import "fmt"

var memo map[int]map[int]int

func calculatePascalValue(row, col int) int {
	if col == 0 || col == row {
		return 1
	}

	if val, ok := memo[row][col]; ok {
		return val
	}

	memo[row][col] = calculatePascalValue(row-1, col-1) + calculatePascalValue(row-1, col)
	return memo[row][col]
}

func generatePascalTriangleMemoized(numRows int) [][]int {
	triangle := make([][]int, numRows)
	memo = make(map[int]map[int]int)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		memo[i] = make(map[int]int)

		for j := 0; j <= i; j++ {
			triangle[i][j] = calculatePascalValue(i, j)
		}
	}

	return triangle
}

func main() {
	numRows := 5
	triangle := generatePascalTriangleMemoized(numRows)

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(triangle[i][j], " ")
		}
		fmt.Println()
	}
}

/*
Main --> generatePascalTriangleMemoized(numRows) --> calculatePascalValue(i,j)

Main function :
 1) In the main function, we set the variable numRows to 5, indicating that we want to generate Pascal's Triangle with 5 rows.
 2) We call the generatePascalTriangleMemoized function with numRows as the argument, which generates and returns the triangle using memoization
 3) Finally, we use nested loops to print the generated triangle. The outer loop iterates over each row, and the inner loop
    prints the elements in that row, separated by a space. After printing each row, we use `fmt.Println()
	to move to the next line and print the next row.

At Global level:  declare a variable called memo of type map[int]map[int]int. This map will be used for memoization,
   which is a technique to store and reuse previously calculated values.

generatePascalTriangleMemoized function  :
Input : takes an integer parameter numRows
Output : returns a two-dimensional slice representing Pascal's Triangle.

Logic :

1) Inside the function, we create an empty two-dimensional slice called triangle with a length equal to numRows.
   We also initialize the memo map to store the calculated values.

6) Next, we iterate over each row and column of the triangle using nested loops.
   Within the loops, we assign the calculated value obtained by calling calculatePascalValue with the current row and column indices
   to the corresponding position in the triangle slice.

calculatePascalValue function:
I/p : two integer parameters: row and col
O/p : It returns an integer representing the value at the specified row and column in Pascal's Triangle

Logic :

1)This function uses memoization to optimize the recursive calculations.
It first checks if the value for the given row and col exists in the memo map.
If it does, it returns the stored value, avoiding redundant calculations.

2) If the value is not found in the memo map, it recursively calls calculatePascalValue with row-1 and col-1, as well as with row-1 and col
   to calculate the value based on the elements in the previous row. It then stores the calculated value in the memo map before returning it.


*/
