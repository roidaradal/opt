## Allocation Problems 

### Resource Allocation 
* Input: 
  * List of items, each with count, cost, and value assigned to it 
  * Budget limit
* Task: **Assign a number** for each item, that does not exceed its count, such that the _total cost of all items does not exceed the budget_
* Goal: **Maximize** the total value of selected items 

### Scene Allocation 
* Input:
  * Number of days, with each day having a min/max number of scenes to shoot
  * List of actors, each with their daily cost 
  * List of scenes, each with list of actors involved in the scene
* Task: **Assign scenes to days** such that _number of daily scenes do not violate the min/max limits_
* Goal: **Minimize** the total cost of production: sum of total actors' fees for each day

## Bin Cover Problems

### Bin Cover
* Input: 
    * Number of bins 
    * Minimum capacity of each bin 
    * Weight of each item
* Task: **Distribute** the items to the bins such that each used bin has _at least minimum capacity_
* Goal: **Maximize** the number of bins used

## Bin Packing Problems 

### Bin Packing
* Input: 
    * Number of bins 
    * Maximum capacity of each bin 
    * Weight of each item
* Task: **Distribute** the items to the bins such that each bin _does not exceed maximum capacity_
* Goal: **Minimize** the number of bins used

## Car Painting Problems

### Car Painting 
* Input: sequence of cars with colors, and integer maxShift
* Task: **Arrange the cars in a sequence**, such that _each car is within <maxShift> distance from its original order_
* Goal: **Minimize** the number of color changes in the sequence

### Binary Paint Shop
* Input: a sequence of 2N cars, from N types of cars
* Task: **Assign a binary color** to each type of car - this will be the color of the first car of this type in the sequence;
  the second car of this type will be colored the opposite.
* Goal: **Minimize** the number of color changes in the full color sequence of all cars

## Car Sequencing Problems 

### Car Sequencing 
* Input: 
  * List of car types, with counts for each type 
  * List of car options, each with a maximum count of cars for a specified window size
  * For each car type, list of car options needed
* Task: **Create a sequence** of all the cars, such that _each car option's window maximum count is not violated_

## Clique Problems

### Clique 
* Input: undirected graph
* Task: **Select a subset** of vertices such that the _selected vertices are all connected to each other_
* Goal: **Maximize** the number of selected vertices

### K-Clique
* Input: undirected graph and integer K
* Task: **Select a subset** of vertices of size K, such that the _selected vertices are all connected to each other_

## Clique Cover Problems

### Clique Cover
* Input: undirected graph
* Task: **Partition the vertices** into groups, such that _each group is a clique_
* Goal: **Minimize** the number of cliques

## Dominating Set Problems

### Dominating Set 
* Input: undirected graph
* Task: **Select a subset** of vertices such that _all vertices are either in the set or has neighbor in the set_
* Goal: **Minimize** the number of selected vertices

### Edge Dominating Set
* Input: undirected graph
* Task: **Select a subset** of edges such that _all edges have at least one endpoint covered by an edge in the set_
* Goal: **Minimize** the number of selected edges

### Efficient Dominating Set
* Input: undirected graph
* Task: **Select a subset** of vertices such that _all vertices are dominated exactly once_
* Goal: **Minimize** the number of selected vertices

## Edge Coloring Problems

### Edge Coloring 
* Input: undirected graph and list of colors
* Task: **Assign colors to edges**, such that _for each vertex, all edges connected to it have different colors_
* Goal: **Minimize** the number of colors used

## Edge Cover Problems 

### Edge Cover
* Input: undirected graph
* Task: **Select a subset** of edges, such that _all vertices are covered by at least one selected edge endpoint_
* Goal: **Minimize** the number of selected edges

## Flow Shop Scheduling Problems 

### Flow Shop Scheduling 
* Input:
  * List of machines
  * List of jobs, with each job having list of tasks that have associated durations at given machines
* Task: **Arrange the sequence** of jobs to do, subject to the ff:
  * Each job can only process one task at a time 
  * Each machine can only process one task at a time
* Goal: **Minimize** the total makespan of the schedule

## Graph Matching Problems

### Max Cardinality Matching
* Input: undirected graph
* Task: **Select a subset** of edges to produce a **matching**: vertices used by the edges are only covered once
* Goal: **Maximize** the number of selected edges

### Max Weight Matching
* Input: undirected graph, with edge weights
* Task: **Select a subset** of edges to produce a **matching**: vertices used by the edges are only covered once
* Goal: **Maximize** the sum of selected edge weights

### Rainbow Matching
* Input: undirected graph, with colored edges
* Task: **Select a subset** of edges to produce a **rainbow matching**: a matching where all edges have different colors
* Goal: **Maximize** the number of selected edges

## Graph Partition Problems

### Graph Partition
* Input:
  * Number of partitions N
  * Min partition size
  * Weighted, undirected graph
* Task: **Partition** the vertices up to _N groups_, such that each _group size is at least the minimum partition size_
* Goal: **Minimize** the sum of the weights of _crossing edges_ (v1 and v2 of edge belong in different groups)

## Graph Path Problems 

### Longest Path 
* Input: list of vertices, start and end vertex, and vertex distance matrix
* Task: **Create a path** from start to end vertex, where no vertex is visited twice
* Goal: **Maximize** the sum of edge distances in the path

### Shortest Path
* Input: list of vertices, start and end vertex, and vertex distance matrix
* Task: **Create a path** from start to end vertex, where no vertex is visited twice
* Goal: **Minimize** the sum of edge distances in the path

### Minimax Path
* Input: list of vertices, start and end vertex, and vertex distance matrix
* Task: **Create a path** from start to end vertex, where no vertex is visited twice
* Goal: **Minimize** the max-weight edge in the path

### Widest Path
* Input: list of vertices, start and end vertex, and vertex distance matrix
* Task: **Create a path** from start to end vertex, where no vertex is visited twice
* Goal: **Maximize** the min-weight edge in the path

## Graph Tour Problems 

### Eulerian Path 
* Input: undirected graph
* Task: **Create a sequence** of edges that forms a Eulerian path: _visits each edge exactly once_

### Eulerian Cycle
* Input: undirected graph
* Task: **Create a sequence** of edges that forms a Eulerian cycle: _visits each edge exactly once_, and _ends where it started_

### Hamiltonian Path
* Input: undirected grap
* Task: **Create a sequence** of vertices that forms a Hamiltonian path: _visits each vertex exactly once_

### Hamiltonian Cycle
* Input: undirected graph
* Task: **Create a sequence** of vertices that forms a Hamiltonian cycle: _visits each vertex exactly once_, and _ends where it started_

## Independent Set Problems

### Independent Set 
* Input: undirected graph 
* Task: **Select a subset** of vertices, such that _none are connected to each other_
* Goal: **Maximize** the number of selected vertices

### Rainbow Independent Set
* Input: undirected graph, where each vertex has an associated color
* Task: **Select a subset** of vertices, such that _none are connected to each other_, and _each selected vertex has a different color_
* Goal: **Maximize** the number of selected vertices

## Induced Path Problems 

### Max Induced Path 
* Input: undirected graph
* Task: **Select a sequence** of vertices to form an induced path: _adjacent vertices in the path have an edge in the graph,
   while non-adjacent vertices are not connected by an edge in the graph_
* Goal: **Maximize** the length of the induced path

## Interval Problems

### Activity Selection
* Input: 
    * List of activities
    * List of start times
    * List of end times
* Task: **Select a subset** of _non-overlapping_ activities
* Goal: **Maximize** the number of selected activities

### Weighted Activity Selection
* Input: 
    * List of activities
    * List of start times
    * List of end times
    * Weight of each activity
* Task: **Select a subset** of _non-overlapping_ activities
* Goal: **Maximize** the sum of weights of selected activities

## K-Center Problems 

### K-Center 
* Input: undirected graph, with edge weights, and integer K
* Task: **Select a subset** of vertices that will act as the _centers_
* Goal: **Minimize** the maximum shortest distance of any vertex to the selected centers

## K-Cut Problems

### Min K-Cut
* Input: undirected graph, with edge weights, and integer K
* Task: **Select a subset** of edges to cut to _produce at least K connected components_
* Goal: **Minimize** the sum of selected edge weights 

### Max K-Cut
* Input: undirected graph, with edge weights, and integer K
* Task: **Select a subset** of edges to cut to _produce exactly K connected components_
* Goal: **Maximize** the sum of selected edge weights

## Knapsack Problems

### 0-1 Knapsack 
* Input: 
    * Knapsack capacity
    * List of items
    * List of item weights
    * List of item values
* Task: **Select a subset** of the items such that the _sum of the selected item weights do not exceed the knapsack capacity_
* Goal: **Maximize** the sum of the selected item's values

### Quadratic Knapsack
* Input: 
    * Knapsack capacity
    * List of items
    * List of item weights
    * List of item values
    * Bonus value for item pairs
* Task: **Select a subset** of the items such that the _sum of the selected item weights do not exceed the knapsack capacity_
* Goal: **Maximize** the total value of the selected items:
    * Sum of the selected item's values
    * SUm of the bonus values for item pairs

## Max Coverage Problems

### Max Coverage
* Input:
  * Universal set of items
  * List of subsets
  * Limit number N
* Task: **Select a combination** of up to N subsets
* Goal: **Maximize** the number of universal items covered by the selected subsets

## Number Coloring Problems 

### Sum Coloring 
* Input: undirected graph and list of numbers 
* Task: **Assign numbers to vertices**, such that _all adjacent vertices in the graph have different numbers_
* Goal: **Minimize** the sum of numbers used in the assignment

## Number Partition Problems

### Number Partition
* Input: list of numbers
* Task: **Split** the numbers into 2 groups such that the _sum of the 2 groups are equal_
* Goal: If optimization, **minimize** the difference between the 2 partition sums

## Nurse Scheduling Problems

### Nurse Scheduling 
* Input: 
  * List of days
  * List of shifts, each shift having min/max limits
  * List of nurses, each nurse having preferred days and shifts
  * Global limits: MaxConsecutive, MaxTotal, MaxDaily
* Task: **Assign nurses to shifts** such that:
  * the number of assigned nurses per shift fall within min/max limit
  * all nurse shift counts do not exceed the MaxTotal 
  * all nurse daily shifts do not exceed the MaxDaily 
  * all nurse schedules do not exceed the MaxConsecutive
* Goal: **Minimize** the penalty incurred by not following the nurses' preferred days and shifts

## Satisfaction Problems

### Exact Cover
* Input:
  * Universal set of items
  * List of subsets
* Task: **Select a combination** of the subsets such that each item in the universal set is _covered exactly once_ by one of the selected subsets

### Langford Pair
* Input: N
* Task: **Arrange** the numbers 1,1,2,2,...,N,N such that for each number pair, their _gap in the sequence == the number_ 

### Magic Series
* Input: N
* Task: **Assign** a value [0,N] to the slots [0,N] such that the _number assigned at slot X also appears X times in the sequence_

### N-Queens
* Input: N 
* Task: **Arrange** the N queens in the NxN board such that _no queens attack each other_ (horizontal, vertical, diagonal) 

### Topological Sort 
* Input: directed graph
* Task: **Arrange** the vertices in topological order: _all directed edges should be pointing forward_

## Set Cover Problems 

### Set Cover 
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Select a combination** of the subsets such that _each universal item is covered at least once_ by the selected subsets
* Goal: **Minimize** the number of selected subsets

## Set Packing Problems

### Set Packing
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Select a combination** of the subsets such that _each covered item is only covered once_ (no overlap)
* Goal: **Maximize** the number of selected subsets

## Set Splitting Problems

### Set Splitting
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Split** the items into _two groups_
* Goal: **Maximize** the number of subsets that are split by the partition (subset has items in both groups)

## Spanning Tree Problems

### Minimum Spanning Tree
* Input: undirected graph, with edge weights
* Task: **Select a subset** of edges to form a spanning tree: _all vertices are reachable and the tree is connected_
* Goal: **Minimize** the sum of edge weights

### Min-Degree Spanning Tree
* Input: undirected graph
* Task: **Select a subset** of edges to form a spanning tree: _all vertices are reachable and the tree is connected_
* Goal: **Minimize** the max-degree vertex from the spanning tree

### K-Minimum Spanning Tree
* Input: undirected graph, with edge weights, and integer K
* Task: **Select a subset** of edges to form a k-spanning tree: _k vertices are reachable and the tree is connected_
* Goal: **Minimize** the sum of edge weights

## Steiner Tree Problems 

### Steiner Tree 
* Input: undirected graph, with edge weights, and list of terminal vertices 
* Task: **Select a susbet** of edges to form a steiner tree: _all terminal vertices are reachable and the tree is connected_
* Goal: **Minimize** the sum of edge weights

## Subsequence Problems

### Longest Increasing Subsequence 
* Input: List of numbers
* Task: **Select a subsequence** of numbers such that the subsequence is _increasing_
* Goal: **Maximize** the length of the subsequence

### Longest Alternating Subsequence
* Input: List of numbers
* Task: **Select a subsequence** of numbers such that the subsequence is _alterating (down-up)_
* Goal: **Maximize** the length of the subsequence

## Subset Sum Problems 

### Subset Sum
* Input:
    * Target number N 
    * List of numbers
* Task: **Select a subset** of the numbers such that their sum == N (satisfaction) or does not exceed N (optimization)
* Goal: If optimization, **minimize** the difference between N and the sum of the selected numbers

## Traveling Purchaser Problems

### Traveling Purchaser 
* Input:
  * list of items 
  * list of markets (vertices), with distance matrix
  * distance matrix of markets to / from origin
  * cost matrix of buying item at market
* Task: **Create a sequence** for buying all items at available markets
* Goal: **Minimize** the total cost: travel cost to/from origin, traveling between markets, and the prices of items at market it was purchased

## Traveling Salesman Problems

### Traveling Salesman
* Input: graph with a distance matrix
* Task: **Create a sequence** for the vertices: a _path that visits all vertices and returns to where it started_
* Goal: **Minimize** the total path weight

### Bottleneck Traveling Salesman 
* Input: graph with a distance matrix
* Task: **Create a sequence** for the vertices: a _path that visits all vertices and returns to where it started_
* Goal: **Minimize** the maximum edge weight in the path 

## Vertex Coloring Problems

### Vertex Coloring 
* Input: undirected graph and list of colors
* Task: **Assign colors to vertices**, such that _all adjacent vertices in the graph have different colors_
* Goal: **Minimize** the number of colors used

### Complete Coloring
* Input: undirected graph and list of colors
* Task: **Assign colors to vertices**, such that it is a _proper vertex coloring_, and _all color pairs appear at least once_
* Goal: **Maximize** the number of colors used

### Harmonious Coloring
* Input: undirected graph and list of colors
* Task: **Assign colors to vertices**, such that it is a _proper vertex coloring_, and _all color pairs appear at most once_
* Goal: **Minimize** the number of colors used

## Vertex Cover Problems

### Vertex Cover
* Input: undirected graph
* Task: **Select a subset** of vertices, such that _for all edges, at least one endpoint is covered by the subset_
* Goal: **Minimize** the number of selected vertices

## Warehouse Location Problems 

## Warehouse Location 
* Input: 
  * List of warehouses, each with capacity and cost 
  * List of stores, with cost matrix for store x warehouse
* Task: **Assign stores to warehouses**, such that _each warehouse's capacity is not exceeded_
* Goal: **Minimize** the total cost of the assignment: warehouse cost for each time the warehouse is used and the cost of assigning store to warehouse

