package main

import (
	"fmt"
	"sort"
)

type IdxSourceReplacement struct {
	Idx int
	Source string
	Target string
}

func filterIdxs(S string, indexes []int, sources []string, targets []string) []IdxSourceReplacement {
	idxSrcRepls := []IdxSourceReplacement{}

	for idx, _ := range indexes {
		currSIdx := indexes[idx]
		currSource := sources[idx]
		currTarget := targets[idx]

		srcIterator := 0
		srcMatch := false
		for sIterator := currSIdx; sIterator < currSIdx + len(currSource); sIterator++ {
			if S[sIterator] == currSource[srcIterator] {
				// create and add IdxSourceReplacement
				srcMatch = true
				srcIterator++
			} else {
				srcMatch = false
				break
			}
		}

		if srcMatch {
			idxSrcRepls = append(
				idxSrcRepls,
				IdxSourceReplacement {
					Idx:   currSIdx,
					Source: currSource,
					Target: currTarget,
				},
			)
		}
	}

	return idxSrcRepls
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	idxSrcRepls := filterIdxs(S, indexes, sources, targets)

	result := ""

	sort.SliceStable(idxSrcRepls, func (i int, j int) bool {
		return idxSrcRepls[i].Idx < idxSrcRepls[j].Idx
	})

	idxSrcReplsIterator := 0
	for idx := 0; idx < len(S); {
		if (idxSrcReplsIterator < len(idxSrcRepls)) && (idx == idxSrcRepls[idxSrcReplsIterator].Idx) {
			result = fmt.Sprintf("%s%s", result, idxSrcRepls[idxSrcReplsIterator].Target)
			idx += len(idxSrcRepls[idxSrcReplsIterator].Source)
			idxSrcReplsIterator++
		} else {
			result = fmt.Sprintf("%s%c", result, S[idx])
			idx++
		}
	}

	return result
}

func main() {
	fmt.Println(findReplaceString("abcd", []int{0, 2},[]string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceString("abcd", []int{0, 3},[]string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceString("abcd", []int{1, 2},[]string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceString("vmokgggqzp", []int{3,5,1}, []string{"kg","ggq","mo"}, []string{"s","so","bfr"}))
	fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"ab","ec"}, []string{"eee","ffff"}))
}