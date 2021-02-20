
Printing Shape Area - Interfaces:GoLang
---------------------------------------
1. Created a square & triangle type that is a type of struct
2. Triangle has a base and height of type float64
3. Square has a sideLength of type float64
4. Function getArea that has receiver of (s square). The function will multiple the receivers sideLength * 2 and return the area as type float64
5. Function getArea that has a receiver of (t triangle). The function will divide the base / 2 & multiple by height and return the area as type float64.
6. Created a shape type that is an interface. The shape has a property that calls the getArea of type float64. Which is common acorss the different getArea functions.
7. Created a function with an argument of (s shape). The shape is passed into function and prints out the area. 
------------------------------------------------------------------
Important Notes:

Note: both the square and triangle can be passed in as a shape as both structs have the same return type of float64. Therefore satisfies the printArea function.