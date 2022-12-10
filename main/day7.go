package main

var lines []string
var myMap map[string]int

func main7() {

	lines = readLines("./resource/day7.txt")
}

// lines = readLines("test.txt")
// 	queue := queue.New()

// 	getDirQueue()

// 	fmt.Println("The map is: ", myMap)

// 	size := make([]int, 0, len(myMap))
// 	for _, val := range myMap {
// 		size = append(size, val)
// 	}

// 	sort.Ints(size)

// 	fmt.Println("The size is: ", size)

// 	result := 0
// 	for i := 0; i < len(size)-1; i++ {
// 		if size[i] > 100000 {
// 			fmt.Println("Result is: ", result)
// 		}
// 		tmp := size[i]
// 		for j := i + 1; j < len(size); j++ {
// 			if tmp+size[j] > 100000 {
// 				break
// 			}
// 			tmp += size[j]
// 		}
// 		if result < tmp {
// 			result = tmp
// 		}
// 	}

// 	fmt.Println("Result is: ", result)
// }

// func isInt(str string) int {
// 	if n, err := strconv.Atoi(str); err == nil {
// 		return n
// 	}
// 	return -1
// }

// func queueDir() {
// 	queue := queue.New()
// 	myMap := make(map[string]int, 0)
// 	for _, line := range lines {
// 		a := strings.Split(line, " ")
// 		if a[0] == "$" && a[1] == "cd" && a[2] != ".." {
// 			queue.Enqueue(a[2])
// 		} else if a[0] == "$" && a[1] == "ls" {

// 		}
// 	}
// 	fmt.Println("tree dir result 3: ", sum, dirs)
// 	return sum, dirs
// }

// func getDirQueue() {
// 	size, nestedDirs := treeDir(dirName)
// 	// no nested dir
// 	if len(nestedDirs) == 0 {
// 		myMap[dirName] = size
// 		fmt.Println("debug 1")
// 		return size
// 	} else { // if there is no nested dir
// 		subSize := 0
// 		for _, d := range nestedDirs {
// 			if val, ok := myMap[d]; ok { // map has the nested dir value
// 				subSize += val
// 				fmt.Println("debug 2: ", subSize)
// 			} else { // map doesn't have the nested dir value, need to keep looping
// 				tmp := getDirSize(d)
// 				subSize += tmp
// 				fmt.Println("debug 3")
// 			}
// 		}
// 		fmt.Println("debug 4")
// 		myMap[dirName] = size + subSize
// 		return size + subSize
// 	}
// }
