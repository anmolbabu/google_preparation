/*
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.



Example 1:

Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
Example 2:

Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
Output: [3.75000,0.40000,5.00000,0.20000]
Example 3:

Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
Output: [0.50000,2.00000,-1.00000,-1.00000]


Constraints:

1 <= equations.length <= 20
equations[i].length == 2
1 <= Ai.length, Bi.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= Cj.length, Dj.length <= 5
Ai, Bi, Cj, Dj consist of lower case English letters and digits.
 */

package main

type GIDWeight struct {
	Key string
	Val float64
}

func find(gidWeights map[string]GIDWeight, key string) GIDWeight {
	foundGIDWeight, ok := gidWeights[key]
	if !ok {
		gidWeights[key] = GIDWeight{
			Key: key,
			Val: 1.0,
		}

		foundGIDWeight = gidWeights[key]
	}

	if foundGIDWeight.Key != key {
		newEntry := find(gidWeights, foundGIDWeight.Key)
		gidWeights[key] = GIDWeight{
			Key: newEntry.Key,
			Val: foundGIDWeight.Val * newEntry.Val,
		}
	}

	return gidWeights[key]
}

func union(gidWeights map[string]GIDWeight, dividend, divisor string, value float64) {
	dividendGID := find(gidWeights, dividend)
	divisorGID := find(gidWeights, divisor)

	if divisorGID.Key != dividendGID.Key {
		gidWeights[dividendGID.Key] = GIDWeight {
			Key: divisorGID.Key,
			Val: (divisorGID.Val * value)/dividendGID.Val,
		}
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	gidWeights := make(map[string]GIDWeight)
	results := make([]float64, len(queries))

	for idx, currEq := range equations {
		union(gidWeights, currEq[0], currEq[1], values[idx])
	}

	for idx, currQuery := range queries {
		_, dividendFound := gidWeights[currQuery[0]]
		_, divisiorFound := gidWeights[currQuery[1]]

		if !dividendFound || !divisiorFound {
			results[idx] = -1.0
			continue
		}

		dividend := find(gidWeights, currQuery[0])
		divisor := find(gidWeights, currQuery[1])

		if dividend.Key != divisor.Key {
			results[idx] = -1.0
		} else {
			results[idx] = dividend.Val/divisor.Val
		}
	}

	return results
}

