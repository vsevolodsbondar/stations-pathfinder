# Trains

A Go application that simulates train movement through a railway network while finding optimal independent routes between two stations.

The project includes two different path-finding algorithms and two train scheduling algorithms that can be compared in terms of correctness and performance.



## Features

- Parse railway maps from text files
- Validate map correctness
- Find all possible independent sets of routes (DFS)
- Find maximum number of independent routes (Edmonds-Karp / Max Flow)
- Schedule trains without collisions
- Benchmark algorithms
- Extensive unit and integration tests

## Project Structure

```
.
├── cli/                # Command-line argument handling
├── helper/
│   ├── distanceCalculations/
│   ├── routeUtils/
│   ├── schedulingUtils/
│   └── validation/
├── io/                 # Reading and parsing map files
├── models/             # Data models
├── pathfinder/         # Route finding algorithms
├── scheduler/          # Train scheduling algorithms
├── service/            # Algorithm runner
├── tests/              # Integration tests
│   └──testData/        # Test maps
└── *.map               # Example railway maps
```

## Algorithms

### Path Finding

The project contains two independent implementations.

### Algorithm 1

Uses Depth-First Search to enumerate all possible routes and then builds all sets of non-intersecting routes.

Characteristics:

- finds every possible path
- suitable for small and medium graphs
- useful for comparison and validation
- exponential complexity on dense graphs


### Algorithm 2

Uses Edmonds-Karp (Maximum Flow).

Characteristics:

- splits every station into input/output vertices
- intermediate stations have capacity = 1
- start/end capacities equal the number of trains
- extracts independent routes directly from the residual graph

Advantages:

- significantly faster
- scalable to large maps
- guarantees maximum number of independent routes


## Scheduling

Two scheduling algorithms are included.

The scheduler:

- distributes trains between available routes
- launches trains when possible
- guarantees no collisions
- simulates movement turn by turn

Example output:

```
T1-grasslands T2-farms T3-green_belt
T1-suburbs T2-downtown T3-village
T1-clouds T2-metropolis
...
```


## Map Format

Stations:

```
stations:
start,0,0
a,1,0
b,2,0
end,3,0
```

Connections:

```
connections:
start-a
a-b
b-end
```

## Usage

```
go run . [--algorithm=<seva|anatolii>] <network file> <start station> <end station> <train amount>
```

### Arguments

| Argument | Description |
|----------|-------------|
| `network file` | Path to the railway map file |
| `start station` | Name of the starting station |
| `end station` | Name of the destination station |
| `train amount` | Number of trains to schedule |

### Optional Flags

| Flag | Description |
|------|-------------|
| `--algorithm=seva` | Uses the DFS-based path-finding and scheduling implementation |
| `--algorithm=anatolii` | Uses the Max Flow (Edmonds-Karp) path-finding and scheduling implementation |

If no algorithm is specified, the default implementation is used.

### Examples

Run using the DFS implementation:

```bash
go run . --algorithm=seva jungle-desert.map jungle desert 10
```

Run using the Max Flow implementation:

```bash
go run . --algorithm=anatolii jungle-desert.map jungle desert 10
```

Using the default algorithm:

```bash
go run . jungle-desert.map jungle desert 10
```


## Testing

Run all tests:

```
go test ./tests
```

Run scheduler benchmarks:

```
go test -bench=. ./scheduler
```

Run pathfinder benchmarks:

```
go test -bench=. ./pathfinder
```


## Performance

The project contains benchmark suites for both path-finding and scheduling algorithms on:

- long linear graphs
- binary trees
- dense graphs
- large generated maps
- difficult routing scenarios

## Test Data

The repository includes many predefined railway maps covering:

- simple routes
- disconnected graphs
- invalid maps
- duplicate stations
- coordinate validation
- unreachable destinations
- dense graphs
- difficult rerouting cases
- stress tests
- large maps
