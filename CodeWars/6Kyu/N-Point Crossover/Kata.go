package kata

import (
	"sort"
)

// https://www.codewars.com/kata/n-point-crossover/go

// Crossover ...
func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	if len(ns) == 0 {
		return xs, ys
	}

	sorted := removeDuplicatesAndSort(ns)
	trans := false

	indexSorted := 0
	for i := 0; i < len(xs); i++ {
		if indexSorted < len(sorted) && i == sorted[indexSorted] {
			trans = !trans
			indexSorted++
		}

		if trans {
			xs[i], ys[i] = ys[i], xs[i]
		}
	}
	return xs, ys
}

func removeDuplicatesAndSort(ns []int) []int {
	result := []int{}
	sort.Ints(ns)
	for i := 0; i < len(ns); i++ {
		if i == 0 || ns[i] != ns[i-1] {
			result = append(result, ns[i])
		}
	}
	return result
}

/* Clever
func Crossover(ns []int, xs []int,ys []int) ([]int, []int) {
  // Your code here
  nxs := make([]int, len(xs))
  nys := make([]int, len(ys))
  m := make(map[int]bool)
  for _, co := range ns {
    m[co] = false
  }
  t := []int{0,1}
  g := [][]int{xs,ys}

  for i := range xs {
    if _, ok := m[i]; ok {
      t[0], t[1] = t[1], t[0]
    }
    nxs[i], nys[i] = g[t[0]][i], g[t[1]][i]
  }
  return nxs, nys
}
*/

/* Clever
func Crossover(ns []int, xs []int,ys []int) ([]int, []int) {
  x := make(map[int]bool)
  for _, val := range ns {
    x[val] = true
  }
  for k := range x {
    for i := k; i < len(xs); i++ {
      xs[i], ys[i] = ys[i], xs[i]
    }
  }
  return xs, ys
}
*/
