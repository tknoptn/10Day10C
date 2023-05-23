# 10Day10C with #Telsuko

#10Day10Challenges #Telsuko #NavinReddy
#Day1 : Concepts -- Iteration,Recursion,Memoization
#Tech Stack Used : Golang
#Day1-Work -- Done
#Day1-Assignment -- Done

/\* pascal - Iteration
Main function :

1. In the main function, we set the variable numRows to 5, indicating that we want to generate Pascal's Triangle with 5 rows.
2. We call the pascalIterative function with numRows as the argument, which generates and returns the triangle.
3. Finally, we use nested loops to print the generated triangle. The outer loop iterates over each row,
   and the inner loop prints the elements in that row, separated by a space.
4. After printing each row, we use fmt.Println() to move to the next line.

Called function : pascalIterative
Inputs : integer parameter (which is no of rows in the pascaltriangle)
Output : returns a two dimensional array to represent pascal triangle

Steps followed in the function logic:

1. In the function create a empty new two dimensional slice to return pascaltriangle
2. start the for loop which iterates from 0 to numRows-1, representing each row in the triangle
3. Within the loop, we initialize each row by creating an inner slice with a length of i+1.
4. The inner slice represents the elements in that row.
5. Assign 1 to the first and last elements of each row (triangle[i][0] and triangle[i][i])
6. Nested loop calculates the values for the remaining elements in each row. It starts from 1 and continues until i-1 (excluding the first and last elements).
   Each element is calculated by adding the corresponding elements from the previous row (pascaltriangle[i-1][j-1] and pascaltriangle[i-1][j]).
7. After completing the loop, we return the fully constructed pascaltriangle as the result of the function.

\*/

/\*
Pascal's Triangle using a recursive approach

Main function :

1. we set the variable numRows to 5, indicating that we want to generate Pascal's Triangle with 5 rows.
2. We use nested loops to iterate over each row and column of the triangle.
   The outer loop iterates over i, representing the row index, and the inner loop iterates over j, representing the column index.
3. Within the loops, we call the pascalRecursive function with i and j as arguments to calculate the value at each position in the triangle.
   We then print the calculated value followed by a space using fmt.Print.
4. After printing each row, we use fmt.Println() to move to the next line and print the next row.

Function : pascalRecursive
Input : two integer parameters - row and col
Output : returns a integer which represents the value at the specified row and column in Pascal's Triangle.

Logic :

1. It checks if col is either 0 or equal to row. If so, it returns 1 since the first and last elements of each row in Pascal's Triangle are always 1.
2. Otherwise, it recursively calls pascalRecursive with row-1 and col-1, as well as with row-1 and col
   to calculate the value based on the elements in the previous row.

\*/

Memoization :
/\*
Main --> generatePascalTriangleMemoized(numRows) --> calculatePascalValue(i,j)

Main function :

1.  In the main function, we set the variable numRows to 5, indicating that we want to generate Pascal's Triangle with 5 rows.
2.  We call the generatePascalTriangleMemoized function with numRows as the argument, which generates and returns the triangle using memoization
3.  Finally, we use nested loops to print the generated triangle. The outer loop iterates over each row, and the inner loop
    prints the elements in that row, separated by a space. After printing each row, we use `fmt.Println()
    to move to the next line and print the next row.

At Global level: declare a variable called memo of type map[int]map[int]int. This map will be used for memoization,
which is a technique to store and reuse previously calculated values.

Function : generatePascalTriangleMemoized -
Input : takes an integer parameter numRows
Output : returns a two-dimensional slice representing Pascal's Triangle.

Logic :

1. Inside the function, we create an empty two-dimensional slice called triangle with a length equal to numRows.
   We also initialize the memo map to store the calculated values.

2. Next, we iterate over each row and column of the triangle using nested loops.
   Within the loops, we assign the calculated value obtained by calling calculatePascalValue with the current row and column indices
   to the corresponding position in the triangle slice.

Function: calculatePascalValue
I/p : two integer parameters: row and col
O/p : It returns an integer representing the value at the specified row and column in Pascal's Triangle

Logic :

1.This function uses memoization to optimize the recursive calculations.
It first checks if the value for the given row and col exists in the memo map.
If it does, it returns the stored value, avoiding redundant calculations.

2. If the value is not found in the memo map, it recursively calls calculatePascalValue with row-1 and col-1, as well as with row-1 and col
   to calculate the value based on the elements in the previous row. It then stores the calculated value in the memo map before returning it.

\*/

/\*

**\* Day2 \*\***
Product managment solution for a productmanager

used :Golang, slices,strings package, stuct , struct methods Looping mechanism, conditonal statements

Features :

1. Add products
2. list products
3. search functionality based on text in products list
4. Get the products based on the location
5. warranty checking of the products

**_Day 3 _**
Product managment solution for a productmanager

Used : Created product Table for storage, Tools used: pgadmin4 ,
Database : PostgreSQL
Features:
Added DB Interactivity with PostgreSQL Database
and Implemented the features to store the products and
Added search functionality from Database to make the solution robust.
