Mazegen
=======

Maze generator written in go.

### Build

```
go build
```

### Usage
```
mazegen [flags] <columns> <rows> <output_file>

Flags:
      -cell int
        	cell size (default 10)
      -wall int
        	wall thickness (default 4)
      -path string
        	path type: simple, uniform(slow), convoluted (default "simple")
      -solve
        	generate solution along with maze
```