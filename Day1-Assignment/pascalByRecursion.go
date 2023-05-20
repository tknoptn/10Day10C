package main

import "fmt"

func pascalRecursive(row, col int) int {
	if col == 0 || col == row {
		return 1
	}

	return pascalRecursive(row-1, col-1) + pascalRecursive(row-1, col)
}

func main() {
	numRows := 5

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(pascalRecursive(i, j), " ")
		}
		fmt.Println()
	}
}

/*
Pascal's Triangle using a recursive approach

Main function :
1) we set the variable numRows to 5, indicating that we want to generate Pascal's Triangle with 5 rows.
2) We use nested loops to iterate over each row and column of the triangle.
   The outer loop iterates over i, representing the row index, and the inner loop iterates over j, representing the column index.
3) Within the loops, we call the pascalRecursive function with i and j as arguments to calculate the value at each position in the triangle.
   We then print the calculated value followed by a space using fmt.Print.
4) After printing each row, we use fmt.Println() to move to the next line and print the next row.

Functionlogic : pascalRecursive
Input : two integer parameters - row and col
Output : returns a integer which represents the value at the specified row and column in Pascal's Triangle.

Logic :
1) It checks if col is either 0 or equal to row. If so, it returns 1 since the first and last elements of each row in Pascal's Triangle are always 1.
2) Otherwise, it recursively calls pascalRecursive with row-1 and col-1, as well as with row-1 and col
   to calculate the value based on the elements in the previous row.




*/
