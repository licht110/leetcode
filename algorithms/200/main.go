type Island struct {
    lands [][]int 
}

func (i Island) search(c []int) bool {
    for _, l := range i.lands {
        if l[0] == c[0] && l[1] == c[1] {
            return true
        }
    }
    return false
}

func (i *Island) build(c []int, grid *[][]byte) {
    y, x := c[0], c[1]
    if (*grid)[y][x] == 48 {
        return
    }
    fmt.Println(y, x)
    i.lands = append(i.lands, []int{y, x})
    (*grid)[y][x] = 48
    if x+1 < len((*grid)[0]) {
        i.build([]int{y, x+1}, grid)
    }
    if y+1 < len((*grid)) {
        i.build([]int{y+1, x}, grid)
    }
    if y-1 >= 0 {
        i.build([]int{y-1, x}, grid)
    }
    if x-1 >= 0 {
        i.build([]int{y, x-1}, grid)
    }
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    islands := []Island{}
    for y, line := range grid {
        for x, l := range line {
            if l == 49 {
                if len(islands) == 0 {
                    newIsland := Island{lands: [][]int{}}
                    newIsland.build([]int{y, x}, &grid)
                    islands = append(islands, newIsland)
                }
                exist := false
                for _, island := range islands {
                    if island.search([]int{y, x}) {
                        exist = true
                        break
                    }
                }
                if !exist {
                    newIsland := Island{lands: [][]int{}}
                    newIsland.build([]int{y, x}, &grid)
                    islands = append(islands, newIsland)
                    //fmt.Println("processing", y, x, islands)
                }
            }
        }
    }
    return len(islands)
}
