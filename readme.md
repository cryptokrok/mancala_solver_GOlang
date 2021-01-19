# mancala
#### Adversarial search on the traditional game of Mancala
[![Build Status](https://travis-ci.org/kNoAPP/mancala.svg?branch=master)](https://travis-ci.org/kNoAPP/mancala)

This project has two components. The first is a `mancala_solver` Go package designed for traversing Mancala's 
Min-Max decision tree. The second is a command-line driver `main` Go package to demo the solver.

Depth of search, state of the Mancala board, and legal action steps are provided in the solver package.


*An image of the Mancala driver finding optimal moves*
![](https://i.imgur.com/ydFsb5I.png "Mancala Driver")

*Winning the game after playing optimally*
![](https://i.imgur.com/FDhi0I6.png "Mancala Finish")

## Installation
This project uses Go to compile its source code. You can [get Go at their website](https://golang.org/dl/).

With Go installed, clone this repository to your [GODIR](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable) 
workspace. By default, this is usually your user's home folder.

```
git clone https://github.com/kNoAPP/mancala.git
cd mancala
```

Then build it with Go!
```
go build
```

A binary should be dropped in either the `Mancala` folder or your workspace's `bin/Mancala` folder. Run this binary
with `./mancala` if on Linux or `mancala` or `mancala.exe` if on Windows.

## As a Go Package
You can also directly import this project to your own Go projects! Run the following inside your project folder.
```
go get github.com/kNoAPP/mancala
```
Then, import the `mancala_solver` package with
```go
import "github.com/kNoAPP/mancala/pkg/mancala_solver"
```

You can now use `mancala_solver` with any of the following useful functionality:
1. `.MancalaState`
2. `.CalculateMove`
3. `.AdvanceState`
4. `#PrintBoard`
5. `#IsEndOfGame`

## Adversarial Search
This program uses [min-max search](https://www.cpp.edu/~ftang/courses/CS420/notes/adversarial%20search.pdf) on 
Mancala's decision tree to pick moves. Moves are made assuming the opponent is also playing optimally. 
But humanity isn't optimal, and we can't have nice things. 

So moves made by this program tend to air on the cautious side. An obvious, huge point gain may not be taken
if the search believes the opponent could regain advantage later. It prefers small risk victories over high risk 
rewards.

## Room for Improvement
Min-max tree searching grows exponentially with depth. If you want to consider the next 10 moves over the next 9, 
that becomes a *substantially bigger* problem. And bigger problems take longer to compute. 

In my own testing, I found this program can reasonably search up to **12 moves** in the future before becoming 
too slow. Here are some strategies to increase that number to maybe 21 or 22.

1. [Use Alpha-Beta Pruning](https://www.geeksforgeeks.org/minimax-algorithm-in-game-theory-set-4-alpha-beta-pruning/) - 
   This technique minimizes the scope of adversarial search. Basically, in min-max search, you maximize the minimum
   value or the exact opposite. If you're at a max node, you've got 37 already, and the min node you search returns a
   10, then there's no point to search the rest of that tree since at maximum, you'll get a 10 from that node. This
   technique can **massively** reduce search scope.
   
2. [Use Goroutines](https://tour.golang.org/concurrency/1) - This program currently runs synchronously. Since the 
   code doesn't run concurrently, it misses out on improved speed. If you broke up the search into many smaller tasks,
   you could run it on more than one logical core.
   
3. [Dynamic Programming](https://www.geeksforgeeks.org/dynamic-programming/) - Cache the results from the search. By
   keeping a list of optimal moves for each game state, you don't have to re-run the same calculations over and over.