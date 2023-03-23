# streams

Basic Stream API written in Go.(Only supports Go1.18+)

[![CircleCI](https://circleci.com/gh/go-the-way/streams/tree/main.svg?style=shield)](https://circleci.com/gh/go-the-way/streams/tree/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-the-way/streams)
[![codecov](https://codecov.io/gh/go-the-way/streams/branch/main/graph/badge.svg?token=8MAR3J959H)](https://codecov.io/gh/go-the-way/streams)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-the-way/streams)](https://goreportcard.com/report/github.com/go-the-way/streams)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-the-way/streams?status.svg)](https://pkg.go.dev/github.com/go-the-way/streams?tab=doc)

## Functions

### Condition Functions
- Ternary
- TernaryV
- If
- IfV
- IfElse
- SwitchCase
- SwitchCaseV

### Filter Functions
- Filter
- FilterThenCount
- FilterThenMap
- FilterThenToMap
- FilterThenGroupBy
- FilterThenReduce
- FilterThenToSet

### ForEach Functions
- ForEach
- ForEachPtr
- MapEach
- ForEachConcurrent
- ForEachPtrConcurrent
- MapEachConcurrent

### Generate Functions
- Generate

### GroupBy Functions
- GroupBy
- GroupByThenMap
- GroupByThenToMap
- GroupByCounting
- GroupByCountingThenMap
- GroupByCountingThenToMap

### Map Functions
- Map
- MapThenFilter
- MapThenReduce
- MapThenGroupBy
- MapMap
- MapMapThenFilter
- MapMapThenReduce

### Reduce Functions
- Reduce
- ReduceMap

### ToMap Functions
- ToMap
- ToMapThenMap
- ToMap2
- ToMap3
- MapToMap
- MapToMapThenFilter
- MapToMapThenReduce

### ToSet Functions
- ToSet
- ToSetThenMap
- ToSet2
- ToSet3