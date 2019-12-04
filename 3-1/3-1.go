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
    file, err := os.Open(dir + "\\3-1\\input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    pathList := []string {}
    intersectionList := [] string {}
    lineNr := 0
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        findIntersections(scanner.Text(), lineNr, &pathList, &intersectionList)
        lineNr += 1
    }

    closestIntersection := math.MaxInt64
    for _, a := range intersectionList {
        csum := sum(a)
        if csum < closestIntersection {
            closestIntersection = csum
        }
    }

    fmt.Println(closestIntersection)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func findIntersections(line string, lineNr int, pathList *[]string, intersectionList *[]string) {
    x := 0
    y := 0

    directions := strings.Split(line, ",")
    for _, dir := range directions {
        cmd := strings.SplitN(dir, "", 2);
        nr, _ := strconv.Atoi(cmd[1])

        switch cmd[0] {
        case "R":
            for i := 0; nr > i; i++ {
                x += 1
                appendToList(x, y, lineNr, pathList, intersectionList)
            }
        case "U":
            for i := 0; nr > i; i++ {
                y += 1
                appendToList(x, y, lineNr, pathList, intersectionList)
            }
        case "L":
            for i := 0; nr > i; i++ {
                x -= 1
                appendToList(x, y, lineNr, pathList, intersectionList)
            }
        case "D":
            for i := 0; nr > i; i++ {
                y -= 1
                appendToList(x, y, lineNr, pathList, intersectionList)
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

func appendToList(x int, y int, lineNr int, pathList *[]string, intersectionList *[]string) {
    crd := coord(x,y)
    if (lineNr == 0) {
        *pathList = append(*pathList, crd);
    } else if contains(pathList, crd) {
        *intersectionList = append(*intersectionList, crd);
    }
}

func sum(crd string) int {
    c := strings.Split(crd, "_")
    c1, _ := strconv.Atoi(c[0])
    c2, _ := strconv.Atoi(c[1])
    if c1 < 0 {
        c1 = -c1
    }
    if c2 < 0 {
        c2 = -c2
    }
    return c1 + c2
}
