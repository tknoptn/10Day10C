package main

import "fmt"

func pascalIterative(numRows int) [][]int {
	pascaltriangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		pascaltriangle[i] = make([]int, i+1)
		pascaltriangle[i][0] = 1
		pascaltriangle[i][i] = 1

		for j := 1; j < i; j++ {
			pascaltriangle[i][j] = pascaltriangle[i-1][j-1] + pascaltriangle[i-1][j]
		}
	}

	return pascaltriangle
}

func main() {
	numRows := 5
	triangle := pascalIterative(numRows)

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(triangle[i][j], " ")
		}
		fmt.Println()
	}
}

/*
Main function :

1) In the main function, we set the variable numRows to 5, indicating that we want to generate Pascal's Triangle with 5 rows.
2) We call the pascalIterative function with numRows as the argument, which generates and returns the triangle.
3) Finally, we use nested loops to print the generated triangle. The outer loop iterates over each row,
   and the inner loop prints the elements in that row,  separated by a space.
4) After printing each row, we use fmt.Println() to move to the next line.

Called function : pascalIterative
Inputs :  integer parameter (which is no of rows in the pascaltriangle)
Output : returns a two dimensional array to represent pascal triangle

Steps followed in the function logic:
1) In the function create a empty new two dimensional slice to return pascaltriangle
2) start the for loop which iterates from 0 to numRows-1, representing each row in the triangle
3) Within the loop, we initialize each row by creating an inner slice with a length of i+1.
4) The inner slice represents the elements in that row.
5) Assign 1 to the first and last elements of each row (triangle[i][0] and triangle[i][i])
6) Nested loop calculates the values for the remaining elements in each row. It starts from 1 and continues until i-1 (excluding the first and last elements).
   Each element is calculated by adding the corresponding elements from the previous row (pascaltriangle[i-1][j-1] and pascaltriangle[i-1][j]).
7) After completing the loop, we return the fully constructed pascaltriangle as the result of the function.

*/
