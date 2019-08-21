# CLI task manager

> [exercice 7 from Gophercises](https://gophercises.com/exercises/task)

```bash
# Build
make compile
# Run binary
./bin/task -h
# Or directly using go
go run . -h
# By default, store tasks in BoltDB
go run . add Do gophercises
# You can also store tasks in a YAML file
go run . -s yaml add Some new task
```
