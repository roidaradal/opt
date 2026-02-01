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

## Edge Cover Problems 

### Edge Cover
* Input: undirected graph
* Task: **Select a subset** of edges, such that _all vertices are covered by at least one selected edge endpoint_
* Goal: **Minimize** the number of selected edges

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

## Independent Set Problems

### Independent Set 
* Input: undirected graph 
* Task: **Select a subset** of vertices, such that _none are connected to each other_
* Goal: **Maximize** the number of selected vertices

### Rainbow Independent Set
* Input: undirected graph, where each vertex has an associated color
* Task: **Select a subset** of vertices, such that _none are connected to each other_, and _each selected vertex has a different color_
* Goal: **Maximize** the number of selected vertices

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

## Number Partition Problems

### Number Partition
* Input: list of numbers
* Task: **Split** the numbers into 2 groups such that the _sum of the 2 groups are equal_
* Goal: If optimization, **minimize** the difference between the 2 partition sums

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

## Vertex Cover Problems

### Vertex Cover
* Input: undirected graph
* Task: **Select a subset** of vertices, such that _for all edges, at least one endpoint is covered by the subset_
* Goal: **Minimize** the number of selected vertices