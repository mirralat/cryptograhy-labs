package main

import (
	"fmt"
	"strconv"
)


func expansion(keyPart []string) []string {
	var expendedKey []string = make([]string, 48)
	expansionTable := [48]int{31, 0, 1, 2, 3, 4, 3, 4, 5, 6, 7, 8, 7, 8, 9, 10, 11, 12, 11, 12, 13, 14, 15, 16, 15, 16, 17, 18, 19, 20, 19, 20, 21, 22, 23, 24, 23, 24, 25, 26, 27, 28, 27, 28, 29, 30, 31, 0,
    }

	for key, item := range expansionTable{
		expendedKey[key] = keyPart[item]
	}
	return expendedKey
}


func bitwise_xor(expendedKey []string) []string {
	var roundKey = []string{"1", "1", "1", "0", "0", "0", "0", "0", "1", "0", "1", "1", "1", "1", "1", "0", "0", "1", "1", "0", "0", "1", "1", "0", "1", "1", "1", "1", "0", "1", "1", "1", "0", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1", "1", "1", "0", "0", "0"}
	var xoredExpendedKey []string = make([]string, 48)

	for key, value := range roundKey{
		keyPart, err := strconv.Atoi(value)
		
		if err != nil {
		    panic(err)
		}
		
		expendedPart, err := strconv.Atoi(expendedKey[key])
		
		if err != nil {
		    panic(err)
		}
		
		res := keyPart ^ expendedPart
		result := strconv.Itoa(res)
		xoredExpendedKey[key] = result
	}
	return xoredExpendedKey
}


func s_box_generation(xorResultKey []string) []string {

	var sBoxTableRow0 = [][]int{{0, 0, 14}, {0, 1, 4}, {0, 2, 13}, {0, 3, 1}, {0, 4, 2}, {0, 5, 15}, {0, 6, 11}, {0, 7, 8}, {0, 8, 3}, {0, 9, 10}, {0, 10, 6}, {0, 11, 12}, {0, 12, 5}, {0, 13, 9}, {0, 14, 0}, {0, 15, 7}}
	var sBoxTableRow1 = [][]int{{1, 0, 0}, {1, 1, 15}, {1, 2, 7}, {1, 3, 4}, {1, 4, 14}, {1, 5, 2}, {1, 6, 13}, {1, 7, 1}, {1, 8, 10}, {1, 9, 6}, {1, 10, 12}, {1, 11, 11}, {1, 12, 9}, {1, 13, 5}, {1, 14, 3}, {1, 15, 8}}
	var sBoxTableRow2 = [][]int{{2, 0, 4}, {2, 1, 1}, {2, 2, 14}, {2, 3, 8}, {2, 4, 13}, {2, 5, 6}, {2, 6, 2}, {2, 7, 11}, {2, 8, 15}, {2, 9, 12}, {2, 10, 9}, {2, 11, 7}, {2, 12, 3}, {2, 13, 10}, {2, 14, 5}, {2, 15, 0}}
	var sBoxTableRow3 = [][]int{{3, 0, 15}, {3, 1, 12}, {3, 2, 8}, {3, 4, 4}, {3, 5, 9}, {3, 6, 1}, {3, 7, 7}, {3, 8, 5}, {3, 9, 11}, {3, 10, 3}, {3, 11, 14}, {3, 12, 10}, {3, 13, 0}, {3, 14, 6}, {3, 15, 13}}

	var rowsStable = map[string][][]int{
		"0": sBoxTableRow0,
		"1": sBoxTableRow1,
		"2": sBoxTableRow2,
		"3": sBoxTableRow3,
	}
    fmt.Println(rowsStable)
	var sBoxBase [][]string
	var slicesIndx []int
	var mock []string
	length := len(xorResultKey)
	foo := 0
	bar := 0

	for foo < length-1{
		foo += 6
		slicesIndx = append(slicesIndx, foo)
	}
	for _, value := range slicesIndx {
		sliceXor := xorResultKey[bar:value]
		sBoxBase = append(sBoxBase, sliceXor)
		bar = value
	}
	for _, value := range sBoxBase{
		string_position := value[0] + value[5]
		fmt.Println(string_position)
	}

	return mock

}

/*
func permutation() string{
	permutationTable := []int{}
	var resultOfRound 
}
*/

func main () {
	var key = []string{"0", "0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "1", "1", "1", "1", "1", "0", "0", "1", "0", "1", "0", "1", "0", "0", "1", "1", "1", "0", "0", "1", "1",
	}
	expended := expansion(key)
	fmt.Println(expended)
	xored := bitwise_xor(expended)
	fmt.Println(xored)
	data := s_box_generation(xored)
	fmt.Println(data)
}