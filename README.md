# Go-Dijkstra
This is my first Go project. I took the opportunity to have a refresher on data structures and algorithms by creating a random adjacency matrix and inputting it to a function that performs Dijkstra's shortest path algorithm to find the shortest distance from 1 node, to any other node.

## Instructions for Use
To demonstrate how the project works, issue the command  
```golang
go run .
```
The following text will be displayed
```
Welcome to Davis Henckel's implementation of Dijkstra's Algorithm written in Go.
================================================================================
Enter the number of nodes for this graph: 
```
Enter any number between `10-1000` and the menu will be displayed
```
Please choose from the following options
1: Print Adjacency Matrix
2: Find Shortest Path from Node A to B
3: Find Shortest Path to all Nodes
4: Create New Adjacency Matrix
```
As the menu states, enter 1 to display the matrix, in the example, number of nodes for this graph is 10
```
--------------------------------------
| 0  1  2  3  4  5  6  7  8  9 |#####|
--------------------------------------
| 0  6  2  0  6  0  0  0  0  4 |  0  |
| 6  0  0  7  0  0  0  0  5  6 |  1  |
| 2  0  0  0  0  4  7  8  0  0 |  2  |
| 0  7  0  0  1  3  0  0  0  0 |  3  |
| 6  0  0  1  0  0  0  0  0  0 |  4  |
| 0  0  4  3  0  0  0  6  0  0 |  5  |
| 0  0  7  0  0  0  0  4  0  0 |  6  |
| 0  0  8  0  0  6  4  0  2  0 |  7  |
| 0  5  0  0  0  0  0  2  0  0 |  8  |
| 4  6  0  0  0  0  0  0  0  0 |  9  |
--------------------------------------

Please choose from the following options
1: Print Adjacency Matrix
2: Find Shortest Path from Node A to B
3: Find Shortest Path to all Nodes
4: Create New Adjacency Matrix
```
The menu will be shown again. From here choose any of the below options. Option 2 will show data like
```
Pick starting node: 2
Pick destination node: 8
Shortest distance from node 2 to node 8 is 10
The path is: 2->7->8
```
Option 3 will show data like 
```
Pick starting node: 2
Total distances from 2 to each node
Node: 0 -- Distance: 2
Node: 1 -- Distance: 8
Node: 2 -- Distance: 0
Node: 3 -- Distance: 7
Node: 4 -- Distance: 8
Node: 5 -- Distance: 4
Node: 6 -- Distance: 7
Node: 7 -- Distance: 8
Node: 8 -- Distance: 10
Node: 9 -- Distance: 6
```
Select option 4 to create a new adjacency matrix and recreate the process. During development I used an awesome tool [Graph Online](https://graphonline.ru/en/) to input the adjacency matrices to visualize the graph and double check the solution provided.  
  
  Select option 5 to print the graph in an easy copy/paste format to easily input it on [Graph Online](https://graphonline.ru/en/)!
  
### Thanks for checking out my first Go Project!