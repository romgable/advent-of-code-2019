package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "math"
)

func main() {
    dir, err := os.Getwd()
    file, err := os.Open(dir + "\\3-2\\input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    pathList := []string {}
    intersectionList := [] string {}
    stepMap1 := make(map[string]int)
    stepMap2 := make(map[string]int)
    lineNr := 0
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        findIntersections(scanner.Text(), lineNr, &pathList, &intersectionList, stepMap1, stepMap2)
        lineNr += 1
    }
    
    leastSteps := math.MaxInt64
    for _, a := range intersectionList {
        csum := getSteps(a, stepMap1, stepMap2)
        if csum < leastSteps {
            leastSteps = csum
        }
    }

    fmt.Println(leastSteps)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func findIntersections(line string, lineNr int, pathList *[]string, intersectionList *[]string, stepMap1 map[string]int, stepMap2 map[string]int) {
    x := 0
    y := 0
    steps := 0

    directions := strings.Split(line, ",")
    for _, dir := range directions {
        cmd := strings.SplitN(dir, "", 2)
        nr, _ := strconv.Atoi(cmd[1])

        switch cmd[0] {
        case "R":
            for i := 0; nr > i; i++ {
                x += 1
                steps++
                appendToList(x, y, steps, lineNr, pathList, intersectionList, stepMap1, stepMap2)
            }
        case "U":
            for i := 0; nr > i; i++ {
                y += 1
                steps++
                appendToList(x, y, steps, lineNr, pathList, intersectionList, stepMap1, stepMap2)
            }
        case "L":
            for i := 0; nr > i; i++ {
                x -= 1
                steps++
                appendToList(x, y, steps, lineNr, pathList, intersectionList, stepMap1, stepMap2)
            }
        case "D":
            for i := 0; nr > i; i++ {
                y -= 1
                steps++
                appendToList(x, y, steps, lineNr, pathList, intersectionList, stepMap1, stepMap2)
            }
        }
	}
}

func contains(s *[]string, e string) bool {
    for _, a := range *s {
        if a == e {
            return true
        }
    }
    return false
}

func coord(x int, y int) string {
    return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func appendToList(x int, y int, steps int, lineNr int, pathList *[]string, intersectionList *[]string, stepMap1 map[string]int, stepMap2 map[string]int) {
    crd := coord(x,y)
    if (lineNr == 0) {
        *pathList = append(*pathList, crd)
        stepMap1[crd] = steps
    } else if contains(pathList, crd) {
        *intersectionList = append(*intersectionList, crd)
        stepMap2[crd] = steps
    }
}

func getSteps(crd string, stepMap1 map[string]int, stepMap2 map[string]int) int {
    steps1 := stepMap1[crd]
    steps2 := stepMap2[crd]
    return steps1 + steps2
}
