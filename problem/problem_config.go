package problem

import "github.com/roidaradal/fn/ds"

type binCfg struct {
	numBins  int
	capacity float64
	weight   []float64
}

type intervalCfg struct {
	activities []string
	start      []float64
	end        []float64
	weight     []float64
}

type graphCfg struct {
	*ds.Graph
	extra [][]string
}

type graphPartitionCfg struct {
	*ds.Graph
	numPartitions    int
	minPartitionSize int
	edgeWeight       []float64
}

type knapsackCfg struct {
	capacity  float64
	items     []string
	weight    []float64
	value     []float64
	pairBonus map[[2]int]float64
}

type subsetsCfg struct {
	universal []string
	names     []string
	subsets   [][]string
	extra     [][]string
}
