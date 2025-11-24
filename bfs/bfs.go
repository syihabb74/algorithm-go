package main

import (
	"fmt"
)

// import "fmt";


func main () {

	grid := [][]byte{
    {'1','1','1','1','0',},
    {'1','1','0','0','0',},
    {'1','1','0','1','1',},
    {'0','0','0','1','1',},
    {'0','1','0','0','0',},
    {'0','1','0','1','1',},
	};
    islands := 0;

	directions := [][]int{ {1 ,0 ,}, {-1 ,0 ,}, {0 ,1 ,}, {0 ,-1 ,}, };
    // for direction each byte's '1' // bottom , top, right, left 
	rows, cols := len(grid), len(grid[0]);

    var bfs = func (r, c int) {
        q := [][]int{{r,c}};
        grid[r][c] = '0'

        for len(q) > 0 {
            front := q[0];
            fmt.Println(front, "Frontt")
            q = q[1:];
            fmt.Println(q, "Q AFT")
            row,col := front[0], front[1]

            for _, dir := range directions {
                fmt.Println(q)
                neighborRow, neighborCol := row + dir[0], col + dir[1];
                fmt.Println(neighborRow, neighborCol)

                if neighborRow < 0 ||
                   neighborCol < 0 ||
                   neighborRow >= rows || 
                   neighborCol >= cols ||
                   grid[neighborRow][neighborCol] == '0' {
                    continue
                   }

                   q = append(q, []int{neighborRow,neighborCol});
                   grid[neighborRow][neighborCol] = '0'

                   fmt.Println(grid)

            }

        }
        // fmt.Println(row, col)

    }



    for r := 0; r < rows ; r++ {
        for c := 0; c < cols ; c++ {
            if grid[r][c] != '1' {
                continue
            } // if not '1' looping at that element will be pass
            // fmt.Println(grid[r][c])
            // if not bfs run 
            bfs(r,c)
            islands++
        }
    }

    fmt.Println(grid)






}