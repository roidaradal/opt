## Bin Problems

### Bin Cover
* Input: 
    * Number of bins 
    * Minimum capacity of each bin 
    * Weight of each item
* Task: **Distribute** the items to the bins such that each used bin has <u>at least minimum capacity</u>
* Goal: **Maximize** the number of bins used


### Bin Packing
* Input: 
    * Number of bins 
    * Maximum capacity of each bin 
    * Weight of each item
* Task: **Distribute** the items to the bins such that each bin <u>does not exceed maximum capacity</u>
* Goal: **Minimize** the number of bins used

## Cover Problems

### Exact Cover 
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Select a combination** of the subsets such that each item in the universal set is <u>covered exactly once</u> by one of the selected subsets

### Max Coverage
* Input: 
    * Universal set of items
    * List of subsets
    * Limit number N
* Task: **Select a combination** of up to N subsets
* Goal: **Maximize** the number of universal items covered by the selected subsets

## Interval Problems

### Activity Selection
* Input: 
    * List of activities
    * List of start times
    * List of end times
* Task: **Select a subset** of <u>non-overlapping</u> activities
* Goal: **Maximize** the number of selected activities

### Weighted Activity Selection
* Input: 
    * List of activities
    * List of start times
    * List of end times
    * Weight of each activity
* Task: **Select a subset** of <u>non-overlapping</u> activities
* Goal: **Maximize** the sum of weights of selected activities

## Knapsack Problems

### 0-1 Knapsack 
* Input: 
    * Knapsack capacity
    * List of items
    * List of item weights
    * List of item values
* Task: **Select a subset** of the items such that the <u>sum of the selected item weights do not exceed the knapsack capacity</u>
* Goal: **Maximize** the sum of the selected item's values

### Quadratic Knapsack
* Input: 
    * Knapsack capacity
    * List of items
    * List of item weights
    * List of item values
    * Bonus value for item pairs
* Task: **Select a subset** of the items such that the <u>sum of the selected item weights do not exceed the knapsack capacity</u>
* Goal: **Maximize** the total value of the selected items:
    * Sum of the selected item's values
    * SUm of the bonus values for item pairs

## Partition Problems

### Graph Partition
* Input: 
  * Number of partitions N 
  * Min partition size 
  * Weighted, undirected graph
* Task: **Partition** the vertices up to <u>N groups</u>, such that each <u>group size is at least the minimum partition size</u>
* Goal: **Minimize** the sum of the weights of <u>crossing edges</u> (v1 and v2 of edge belong in different groups)

### Number Partition
* Input: list of numbers
* Task: **Split** the numbers into 2 groups such that the <u>sum of the 2 groups are equal</u>
* Goal: If optimization, **minimize** the difference between the 2 partition sums

## Satisfaction Problems

### Langford Pair
* Input: N
* Task: **Arrange** the numbers 1,1,2,2,...,N,N such that for each number pair, their <u>gap in the sequence == the number</u> 

### Magic Series
* Input: N
* Task: **Assign** a value [0,N] to the slots [0,N] such that the <u>number assigned at slot X also appears X times in the sequence</u>

### N-Queens
* Input: N 
* Task: **Arrange** the N queens in the NxN board such that <u>no queens attack each other</u> (horizontal, vertical, diagonal) 

## Set Problems 

### Set Cover 
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Select a combination** of the subsets such that <u>each universal item is covered at least once</u> by the selected subsets
* Goal: **Minimize** the number of selected subsets

### Set Packing
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Select a combination** of the subsets such that <u>each covered item is only covered once</u> (no overlap)
* Goal: **Maximize** the number of selected subsets

### Set Splitting
* Input: 
    * Universal set of items
    * List of subsets
* Task: **Split** the items into <u>two groups</u>
* Goal: **Maximize** the number of subsets that are split by the partition (subset has items in both groups)

## Subsequence Problems

### Longest Increasing Subsequence 
* Input: List of numbers
* Task: **Select a subsequence** of numbers such that the subsequence is <u>increasing</u>
* Goal: **Maximize** the length of the subsequence

### Longest Alternating Subsequence
* Input: List of numbers
* Task: **Select a subsequence** of numbers such that the subsequence is <u>alterating (down-up)</u>
* Goal: **Maximize** the length of the subsequence

## Subset Sum Problems 

### Subset Sum
* Input:
    * Target number N 
    * List of numbers
* Task: **Select a subset** of the numbers such that their sum == N (satisfaction) or does not exceed N (optimization)
* Goal: If optimization, **minimize** the difference between N and the sum of the selected numbers
