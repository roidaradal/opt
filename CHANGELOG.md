## v0.3.37 - Spanning Tree Problems 
  * **Commit**: 2026-02-09 21:08
  * problem.NewSpanningTree
  * fn.ConstraintSpanningTree
  * fn.SpannedVertices
## v0.3.36 - Graph Path Problems 
  * **Commit**: 2026-02-08 20:55
  * problem.NewGraphPath
  * data.GraphPath, NewGraphPath
  * fn.AsGraphPath, StringGraphPath
  * fn.ConstraintSimplePath
  * fn.PathDistances
## v0.3.35 - K-Center Problems 
  * **Commit**: 2026-02-08 16:04
  * problem.NewKCenter
## v0.3.34 - Induced Path Problems
  * **Commit**: 2026-02-07 19:39
  * problem.NewInducedPath
  * fn.AsPathOrder
## v0.3.33 - K-Cut Problems 
  * **Commit**: 2026-02-06 18:25
  * problem.NewKCut
  * fn.ConnectedComponents
  * data.NewDirectedGraph
  * Add `topological_sort` to Satisfaction problems
## v0.3.32 - Clique Problems 
  * **Commit**: 2026-02-05 21:39
  * problem.NewClique
## v0.3.31 - Vertex Coloring Problems
  * **Commit**: 2026-02-04 23:03
  * problem.NewVertexColoring
## v0.3.30 - Graph Coloring Problems
  * **Commit**: 2026-02-03 22:13
  * data.GraphColoring
  * problem.NewEdgeColoring
  * problem.NewNumberColoring
  * fn.ConstraintProperVertexColoring
  * fn.CoreLookupValueOrder
## v0.3.29 - Graph Tour Problems 
  * **Commit**: 2026-02-02 22:52
  * problem.NewGraphTour
  * fn.IsEulerianPath
  * fn.IsHamiltonianPath
  * fn.StringEulerianPath
  * fn.CoreSortedCycle
  * fn.MirroredSequence, SortedCycle
## v0.3.28 - Matching Problems
  * **Commit**: 2026-02-01 12:59
  * problem.NewGraphMatching
  * Convert Clique Cover into a Graph Partition problem
  * Add discrete.Problem.UniformDomain()
## v0.3.27 - Sequence Problems 
  * **Commit**: 2026-01-31 22:00
  * problem.NewSubsequence 
  * problem.NewSubsetSum
## v0.3.26 - Set Problems
  * **Commit**: 2026-01-31 21:33
  * problem.NewSetCover
  * problem.NewSetPacking
  * problem.NewSetSplitting
## v0.3.25 - Cover Problems
  * **Commit**: 2026-01-31 21:17
  * problem.NewMaxCoverage
  * Add `exact_cover` to Satisfaction
## v0.3.24 - Satisfaction Problems 
  * **Commit**: 2026-01-31 20:55
  * problem.NewSatisfaction 
## v0.3.23 - Knapsack Problems
  * **Commit**: 2026-01-31 20:20
  * problem.NewKnapsack
## v0.3.22 - Interval Problems 
  * **Commit**: 2026-01-31 19:55
  * problem.NewInterval
## v0.3.21 - Independent Set Problems
  * **Commit**: 2026-01-31 19:39
  * problem.NewIndependentSet
## v0.3.20 - Partition Problems 
  * **Commit**: 2026-01-31 19:05
  * problem.NewGraphPartition
  * problem.NewNumberPartition
## v0.3.19 - Dominating Set Problems 
  * **Commit**: 2026-01-31 18:44
  * problem.NewDominatingSet 
## v0.3.18 - Graph Cover Problems
  * **Commit**: 2026-01-31 18:21
  * problem.NewVertexCover 
  * problem.NewEdgeCover
  * problem.NewCliqueCover
  * fn.IsClique
## v0.3.17 - Bin Problems
  * **Commit**: 2026-01-31 17:45
  * problem.NewBinCover
  * problem.NewBinPacking
## v0.3.16 - Problem Data Loaders 
  * **Commit**: 2026-01-31 17:00
  * data.Bins, NewBins 
  * data.Graph, NewUndirectedGraph
  * data.GraphPartition, NewGraphPartition
  * data.GraphVertices, GraphEdges
  * data.Intervals, NewIntervals
  * data.Knapsack, NewKnapsack
  * data.Sequence, NewSequence
  * data.Subsets, NewSubsets
## v0.3.15 - Object Format Loader 
  * **Commit**: 2026-01-31 14:59
  * data.Load - loads the given data file and returns the StringMap of given test case
  * Includes normal string values, list values, and map values
## v0.3.14 - Reformat Dataset
  * **Commit**: 2026-01-31 10:07 
  * Changed existing datasets to object format
## v0.3.13 - Graph Cover Problems 
  * **Commit**: 2026-01-30 19:53
  * problem.GraphCover
    * Vertex Cover 
    * Edge Cover 
    * Clique Cover
  * Merged dominatingSetProblem into graphCoverProblem (duplicate code)
## v0.3.12 - Test Case Data Shape
  * **Commit**: 2026-01-29 21:54
  * Update fn.LoadLines to read 2nd line of test case as data shape
  * Add fn.groupLines to group the lines according to the data shape
  * Adjust all functions that load test case to use the new line grouping
## v0.3.11 - Dominating Set Problems 
  * **Commit**: 2026-01-28 21:44
  * problem.DominatingSet 
    * Dominating Set 
    * Edge Dominating Set 
    * Efficient Dominating Set
## v0.3.10 - Independent Set Problems
  * **Commit**: 2026-01-27 22:12
  * problem.IndependentSet 
    * Independent Set
    * Rainbow Independent Set
## v0.3.9 - Partition Problems
  * **Commit**: 2026-01-26 22:32
  * problem.Partition
    * Graph Partition
    * Number Partition 
  * Add PROBLEMS.md
## v0.3.8 - Cover Problems
  * **Commit**: 2026-01-25 21:23
  * problem.Cover 
    * Exact Cover 
    * Max Coverage
## v0.3.7 - Set Problems 
  * **Commit**: 2026-01-25 20:23 
  * problem.Set 
    * Set Cover
    * Set Packing 
    * Set Splitting 
## v0.3.6 - Bin Problems
  * **Commit**: 2026-01-24 19:00
  * problem.Bin
    * Bin Cover
    * Bin Packing
  * fn.AsPartition
  * fn.PartitionSums, PartitionStrings
  * fn.ScoreCountUniqueValues 
  * fn.CoreSortedPartition 
  * fn.StringPartition
## v0.3.5 - Knapsack Problems
  * **Commit**: 2026-01-24 11:44
  * problem.Knapsack
      * Basic
      * Quadratic
## v0.3.4 - Satisfaction Problems
  * **Commit**: 2026-01-23 23:34
  * problem.Satisfaction
      * Langford Pair
      * Magic Series
      * N-Queens
  * fn.AsSequence
  * fn.ConstraintAllUnique
  * fn.CoreMirroredSequence, CoreMirroredValues
  * fn.StringSequence, StringValues
## v0.3.3 - Split Subset Problems 
  * **Commit**: 2026-01-23 21:41
  * problem.Interval 
      * Basic
      * Weighted 
  * problem.Subsequence 
      * Longest Increasing 
      * Longest Alternating 
  * problem.SubsetSum 
      * Basic
  * fn.ScoreSumWeightedValues
## v0.3.2 - Subset Problems 
  * **Commit**: 2026-01-22 23:04
  * problem.Subset
      * Activity Selection 
      * Longest Increasing Subsequence 
      * Subset Sum 
  * fn.LoadLines
  * fn.StringList, IntList, FloatList
  * fn.AsSubset
  * fn.ScoreSubsetSize
  * fn.String_Subset
## v0.3.1 - Discrete Package
  * **Commit**: 2026-01-21 22:08
  * discrete.SolutionDisplayFn
  * discrete.Inf, NegInf
  * Rename: MapDomain => Domain
  * problem.AddVariableDomains
## v0.3.0 - Problem Grouping
  * **Commit**: 2026-01-20 22:16
  * Reset repo contents 
  * Organize problems into groups
## v0.2.62 - Rainbow Matching 
  * **Commit**: 2026-01-19 20:47
  * problem.RainbowMatching
  * newRainbowMatching
## v0.2.61 - Rainbow Independent Set 
  * **Commit**: 2026-01-18 21:17
  * problem.RainbowIndependentSet 
  * newRainbowIndependentSet
## v0.2.60 - Ma* Induced Path 
  * **Commit**: 2026-01-17 21:26
  * problem.MaxInducedPath 
  * fn.AsPathOrder
## v0.2.59 - Traveling Purchaser Problem 
  * **Commit**: 2026-01-16 23:11
  * problem.TravelingPurchaser 
  * newTravelingPurchaser, tppCfg 
## v0.2.58 - Minimum K-Cut
  * **Commit**: 2026-01-15 22:18
  * problem.MinimumKCut 
  * fn.NewKWeightedGraph
## v0.2.57 - K-Minimum Spanning Tree 
  * **Commit**: 2026-01-14 21:38
  * problem.KMinimumSpanningTree 
  * newKMST
## v0.2.56 - Sum Coloring 
  * **Commit**: 2026-01-13 20:20
  * problem.SumColoring 
  * newSumColoring
## v0.2.55 - Bottleneck TSP 
  * **Commit**: 2026-01-13 19:47
  * problem.BottleneckTSP
  * fn.NewTravelingSalesman 
  * a.TSPCfg
## v0.2.54 - Efficient Dominating Set 
  * **Commit**: 2026-01-12 21:17
  * problem.EfficientDominatingSet
## v0.2.53 - Edge Dominating Set 
  * **Commit**: 2026-01-12 20:58
  * problem.EdgeDominatingSet
## v0.2.52 - Steiner Tree 
  * **Commit**: 2026-01-11 19:10
  * problem.SteinerTree 
  * newSteinerTree 
  * Adjust constraint.AllVerticesSpanned to include list of vertices
## v0.2.51 - Clique Cover 
  * **Commit**: 2026-01-10 20:10
  * problem.CliqueCover
## v0.2.50 - Minima* Path 
  * **Commit**: 2026-01-09 06:18
  * problem.MinimaxPath
## v0.2.49 - Widest Path 
  * **Commit**: 2026-01-08 20:55
  * problem.WidestPath 
  * discrete.PathDomain
## v0.2.48 - Longest Path 
  * **Commit**: 2026-01-08 20:34
  * problem.LongestPath 
  * constraint.SimplePath
## v0.2.47 - Shortest Path 
  * **Commit**: 2026-01-07 21:11
  * problem.ShortestPath
  * a.PathCfg, fn.NewPathProblem
  * discrete.Path
  * fn.AsPath
  * fn.Score_PathCost 
  * fn.String_Path
  * Update brute.LinearSolver to use comb.AllPermutationPositions for Path
  * Change SolutionSpace computation of Path problems
## v0.2.46 - Harmonious Coloring 
  * **Commit**: 2026-01-06 21:25
  * problem.HarmoniousColoring 
  * fn.NewVertexColoring
  * Fix Complete Coloring's constraint 
  * Change Complete Coloring to Maximize: CountUniqueValues
## v0.2.45 - Complete Coloring 
  * **Commit**: 2026-01-05 21:05
  * problem.CompleteColoring 
  * newCompleteColoring
## v0.2.44 - Min-Degree Spanning Tree 
  * **Commit**: 2026-01-04 22:40
  * problem.MinDegreeSpanningTree
  * constraint.AllVerticesSpanned
  * constraint.SpanningTree
## v0.2.43 - Nurse Scheduling 
  * **Commit**: 2026-01-04 02:47
  * problem.NurseSchedule
  * nurseSchedCfg, newNurseSchedule
  * constraint.MaxConsecutive
## v0.2.42 - Topological Sort 
  * **Commit**: 2026-01-03 08:25
  * problem.TopologicalSort 
  * fn.NewDirectedGraph
## v0.2.41 - Ma* Weight Matching 
  * **Commit**: 2026-01-03 08:05
  * problem.MaxWeightMatching
## v0.2.40 - Ma* Cardinality Matching 
  * **Commit**: 2026-01-03 08:00
  * problem.MaxCardinalityMatching
  * constraint.GraphMatching
## v0.2.39 - Edge Cover 
  * **Commit**: 2026-01-02 20:02
  * problem.EdgeCover
## v0.2.38 - Bin Cover 
  * **Commit**: 2026-01-02 07:10
  * problem.BinCover
## v0.2.37 - K-Center 
  * **Commit**: 2026-01-02 07:02
  * problem.KCenter
  * kCenterCfg, newKCenter
  * a.BinProblemCfg 
  * fn.NewBinProblem
## v0.2.36 - Set Splitting 
  * **Commit**: 2026-01-01 15:55
  * problem.SetSplitting
## v0.2.35 - Set Packing   
  * **Commit**: 2026-01-01 15:12
  * problem.SetPacking
## v0.2.34 - Eulerian Cycle
  * **Commit**: 2025-12-31 15:42
  * problem.EulerCycle
  * fn.String_EulerianPath
  * fn.InvalidSolution
## v0.2.33 - Eulerian Path 
  * **Commit**: 2025-12-31 14:53
  * problem.EulerPath
## v0.2.32 - Hamiltonian Cycle 
  * **Commit**: 2025-12-31 13:31
  * problem.HamiltonCycle
  * fn.Core_SortedCycle
## v0.2.31 - Hamiltonian Path 
  * **Commit**: 2025-12-31 12:46
  * problem.HamiltonPath
## v0.2.30 - Dominating Set 
  * **Commit**: 2025-12-31 11:53
  * problem.DominatingSet 
## v0.2.29 - Weapon Target Assignment 
  * **Commit**: 2025-12-31 10:06
  * problem.WeaponTarget 
  * weaponCfg, newWeaponTarget
## v0.2.28 - Quadratic Knapsack 
  * **Commit**: 2025-12-31 08:08
  * problem.QuadraticKnapsack 
  * quadraticKnapsackCfg, newQuadraticKnapsack
  * constraint.Knapsack
## v0.2.27 - Bottleneck Assignment 
  * **Commit**: 2025-12-30 22:09
  * problem.LinearBottleneckAssignment 
  * problem.QuadraticBottleneckAssignment
  * newLBAP
  * fn.String_Assignment
## v0.2.26 - Quadratic Assignment 
  * **Commit**: 2025-12-30 00:12
  * problem.QuadraticAssignment 
  * quadraticAssignmentCfg, newQuadraticAssignment
## v0.2.25 - Generalized Assignment 
  * **Commit**: 2025-12-29 18:30
  * problem.GeneralizedAssignment 
  * generalizedAssignmentCfg, newGeneralizedAssignment
## v0.2.24 - Assignment with Teams 
  * **Commit**: 2025-12-28 22:21
  * Adjust Assignment to include teams
## v0.2.23 - Assignment Problem 
  * **Commit**: 2025-12-27 22:46
  * problem.Assignment 
  * assignmentCfg, newAssignment
  * discrete.IndexVariables
  * fn.ParseFloatInf
## v0.2.22 - Linear Brute Force Sequencer 
  * **Commit**: 2025-12-26 20:09
  * Update brute.LinearSolver for Sequencer 
  * Update problem.SolutionSpace for Sequence types
  * Fix CarSequence constraint
## v0.2.21 - Problem Types 
  * **Commit**: 2025-12-25 21:17
  * discrete.Assignment, Partition, Sequence, Subset
  * Add ProblemType to problems
## v0.2.20 - Traveling Salesman 
  * **Commit**: 2025-12-24 18:41
  * problem.TravelingSalesman 
  * newTravelingSalesman
## v0.2.19 - Flow Shop Scheduling 
  * **Commit**: 2025-12-23 18:55
  * problem.FlowShopSchedule
  * newFlowShop
## v0.2.18 - Open Shop Scheduling 
  * **Commit**: 2025-12-22 18:50
  * problem.OpenShopSchedule 
  * newOpenShop 
  * a.TaskString 
  * constraint.NoJobTaskOverlap
## v0.2.17 - Job Shop Scheduling 
  * **Commit**: 2025-12-21 20:50
  * problem.JobShopSchedule
  * newJobShop
  * a.ShopSchedCfg, a.Job, a.Task
  * a.TimeRange, a.SlotSched
  * a.NewJob, a.NewTask
  * constraint.NoMachineOverlap
  * fn.Score_ScheduleMakespan
  * fn.String_ShopSchedule
## v0.2.16 - Car Sequencing 
  * **Commit**: 2025-12-20 08:07
  * problem.CarSequencing
  * carSequenceCfg, newCarSequencing
## v0.2.15 - Car Painting 
  * **Commit**: 2025-12-19 06:50
  * problem.CarPainting
  * carPaintCfg, newCarPainting
  * a.StrInt, NewStrInt
## v0.2.14 - Binary PaintShop 
  * **Commit**: 2025-12-18 04:25
  * problem.BinaryPaintShop
  * binaryPaintCfg, newBinaryPaintShop
  * fn.CountColorChanges
## v0.2.13 - Scene Allocation
  * **Commit**: 2025-12-17 22:22
  * problem.SceneAllocation
  * sceneCfg, newSceneAllocation
## v0.2.12 - Pool Manager 
  * **Commit**: 2025-12-16 21:49 
  * worker.Pool 
  * worker.Pool.Run
## v0.2.11 - Manager Output 
  * **Commit**: 2025-12-15 21:21
  * Update SpaceSolver's output format: name | space 
  * Horizontal output:
      * SolverRunner 
      * SolutionReader
  * Add Columns() to Workers to return their column formats
## v0.2.10 - Space Solver 
  * **Commit**: 2025-12-14 15:40
  * worker.SpaceSolver
  * worker.SpaceSolver.Run
  * problem.SolutionSpaceEquation
## v0.2.9 - Solo Manager 
  * **Commit**: 2025-12-14 14:58
  * Update workers to return string instead of directly printing:
      * SolverRunner 
      * SolutionSaver
      * SolutionReader
  * worker.Config
  * Updated Worker APIs to use worker.Config to unify into Worker interface
  * Add test suite data: grouped into easy, medium, hard, ultra
  * worker.Manager interface 
  * worker.Solo
  * worker.Solo.Run
## v0.2.8 - SolutionReader 
  * **Commit**: 2025-12-13 19:03
  * worker.SolutionReader
  * worker.SolutionReader.Read
  * Renamed RunReporter => SolverRunner 
  * Renamed SolutionReporter => SolutionSaver 
  * Renamed Reporter interface => Runner
## v0.2.7 - Command-Line Helper 
  * **Commit**: 2025-12-13 13:33
  * Fix command-line parameter structure
  * Add color to help message
  * Display Usage:
      * List of task choices
      - List of problems and range of test cases
      * List of solver options 
      * List of logger options
  * Read command-line args from JSON file
## v0.2.6 - SolutionReporter 
  * **Commit**: 2025-12-13 08:42
  * worker.SolutionReporter 
  * worker.SolutionReporter.Run
## v0.2.5 - Command-Line Parameters 
  * **Commit**: 2025-12-12 21:53
  * RunReporter WithSolution
  * Main command-line Parameters
  * newProblem 
  * newWorker 
  * newSolverCreator
  * newLogger
## v0.2.4 - RunReporter 
  * **Commit**: 2025-12-11 21:34
  * worker.Reporter interface
  * worker.RunReporter 
  * worker.RunReporter.Run
## v0.2.3 - Linear Brute Force Steps 
  * **Commit**: 2025-12-10 21:48
  * worker.NoLogger, BatchLogger, StepsLogger
  * Update Logger interface to Output() and Steps()
  * brute.LinearSolver Steps
## v0.2.2 - Concurrent Brute Force Solver 
  * **Commit**: 2025-12-09 20:54
  * base.ConcurrentTaskFn
  * base.ConcurrentSolver 
  * base.ConcurrentSolver.Initialize
  * base.ConcurrentSolver.RunWorkers
  * base.Solver.DisplayProgress
  * base.Solver.Prelude
  * brute.ConcurrentSolver 
  * brute.NewConcurrentSolver 
  * brute.ConcurrentSolver.Solve
## v0.2.1 - Linear Brute Force Solver 
  * **Commit**: 2025-12-08 21:10
  * brute.LinearSolver 
  * brute.NewLinearSolver
  * brute.LinearSolver.Solve
  * base.Solver.AddSolution
  * worker.NewLogger
## v0.2.0 - Linear Solver
  * **Commit**: 2025-12-07 22:43
  * worker.LogLevel 
  * worker.Solver Interface 
  * worker.Result
  * base.Solver, LinearSolver
  * base.Solver.Initialize, GetResult
  * base.Solver.IsComplete, IsScoreBetter
  * base.LinearSolver.IsSolutionLimitReached
  * base.LinearSolver.IsIterationLimitReached
  * worker.Logger Interface 
  * worker.CmdLogger, NoLogger
## v0.1.20 - Minimum Spanning Tree 
  * **Commit**: 2025-12-06 21:14
  * problem.MinimumSpanningTree
  * fn.NewWeightedGraph
## v0.1.19 - Independent Set 
  * **Commit**: 2025-12-06 20:59
  * problem.IndependentSet
## v0.1.18 - Clique 
  * **Commit**: 2025-12-06 20:53
  * problem.Clique
## v0.1.17 - Vertex Cover 
  * **Commit**: 2025-12-06 20:47
  * problem.VertexCover 
  * fn.NewUnweightedGraph
## v0.1.16 - Graph Partition 
  * **Commit**: 2025-12-06 20:34
  * problem.GraphPartition
  * graphPartitionCfg, newGraphPartition
## v0.1.15 - Graph Coloring 
  * **Commit**: 2025-12-06 17:45
  * problem.GraphColoring
  * newGraphColoring 
  * fn.String_Map
## v0.1.14 - Edge Coloring 
  * **Commit**: 2025-12-06 11:01
  * problem.EdgeColoring 
  * newEdgeColoring 
  * fn.Core_LookupValueOrder
## v0.1.13 - Warehouse Location 
  * **Commit**: 2025-12-06 10:31 
  * problem.WarehouseLocation 
  * warehouseCfg, newWarehouseLocation
## v0.1.12 - Set Cover 
  * **Commit**: 2025-12-06 10:12
  * problem.SetCover
## v0.1.11 - Exact Cover 
  * **Commit**: 2025-12-06 09:40
  * problem.ExactCover 
  * fn.NewSubsets
  * a.Subsets, NewSubsets
## v0.1.10 - Subset Sum 
  * **Commit**: 2025-12-05 23:10
  * problem.SubsetSum 
  * newSubsetSum
## v0.1.9 - Number Partition 
  * **Commit**: 2025-12-05 22:41
  * problem.NumberPartition
  * newNumberPartition
## v0.1.8 - Bin Packing 
  * **Commit**: 2025-12-05 20:54
  * problem.BinPacking 
  * binPackingCfg, newBinPacking 
  * fn.Score_CountUniqueValues 
  * fn.Core_SortedPartition 
  * fn.String_Partitions
  * fn.AsPartitions, PartitionSums, PartitionStrings
## v0.1.7 - N-Queens 
  * **Commit**: 2025-12-04 22:02
  * problem.NQueens 
  * hasDiagonalConflict
  * fn.Core_MirroredValues
## v0.1.6 - Langford Pair 
  * **Commit**: 2025-12-04 21:42
  * problem.LangfordPair 
  * constraint.AllUnique
  * fn.AsSequence
  * fn.CoreMirroredSequence
  * fn.String_Sequence
## v0.1.5 - Knapsack 
  * **Commit**: 2025-12-03 22:17
  * problem.Knapsack 
  * knapsackCfg, newKnapsack
## v0.1.4 - Longest Increasing Subsequence 
  * **Commit**: 2025-12-03 22:09
  * problem.LongestIncreasingSubsequence 
  * newLIS
## v0.1.3 - Magic Series 
  * **Commit**: 2025-12-03 21:56
  * problem.MagicSeries
## v0.1.2 - Resource Optimization 
  * **Commit**: 2025-12-03 21:13
  * problem.ResourceOptimization
  * resourceOptCfg, newResourceOptimziation
  * fn.ScoreSumWeightedValues
  * fn.String_Values
## v0.1.1 - Activity Selection
  * **Commit**: 2025-12-02 22:41
  * problem.ActivitySelection
  * activitySelectionCfg, newActivitySelection
  * fn.LoadProblem
  * fn.AsSubset 
  * fn.SubsetSize, ScoreSubsetSize
  * fn.String_Subset
## v0.1.0 - Discrete Optimization
  * **Commit**: 2025-12-01 21:08
  * discrete.Problem
  * discrete.Variable, Value 
  * discrete.Constraint 
  * discrete.Solution