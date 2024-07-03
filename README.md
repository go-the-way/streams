# streams

Basic Stream API written in Go.(Only supports Go1.18+)

[![CircleCI](https://circleci.com/gh/go-the-way/streams/tree/main.svg?style=shield)](https://circleci.com/gh/go-the-way/streams/tree/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-the-way/streams)
[![codecov](https://codecov.io/gh/go-the-way/streams/branch/main/graph/badge.svg?token=8MAR3J959H)](https://codecov.io/gh/go-the-way/streams)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-the-way/streams)](https://goreportcard.com/report/github.com/go-the-way/streams)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-the-way/streams?status.svg)](https://pkg.go.dev/github.com/go-the-way/streams?tab=doc)

## Types

### Entry
- NewEntry
- MapToEntry
- EntryToMap

#### Map
- NewMap
- NewMapValue
- NewMapValue2
- NewMapValue3
- Len
- Put
- PutAll
- Delete
- Clear
- ContainsKey
- ContainsValue
- Iterate
- Keys
- Values
- KeyReduce
- ValueReduce

#### Set
- NewSet
- NewSetValue
- Len
- Add
- AddAll
- Delete
- Clear
- Contains
- Iterate
- Filter
- Reduce
- List

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

### Filter var Functions
- filterfunc.Eq
- filterfunc.Gt
- filterfunc.GtEq
- filterfunc.Lt
- filterfunc.LtEq

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

### Map var Functions
- mapfunc.Int2BinaryString
- mapfunc.Int2OctalString
- mapfunc.Int2String
- mapfunc.Int2HexString
- mapfunc.Int2Any
- mapfunc.String2Bool
- mapfunc.String2Int
- mapfunc.String2Byte
- mapfunc.String2Int8
- mapfunc.String2Int16
- mapfunc.String2Int32
- mapfunc.String2Int64
- mapfunc.String2UInt
- mapfunc.String2UInt8
- mapfunc.String2UInt16
- mapfunc.String2UInt32
- mapfunc.String2UInt64
- mapfunc.String2Float32
- mapfunc.String2Float64
- mapfunc.String2Any

### Reduce Functions
- Reduce
- ReduceMap

### Reduce var Functions
- reducefunc.Number
- reducefunc.NumberFunc
- reducefunc.Slice
- reducefunc.SliceFunc

### Reduce val Functions
- ReduceString
- ReduceStringWithComma
- ReduceStringWithDot
- ReduceStringWithSpace
- ReduceStringWithLine
- ReduceInt8
- ReduceInt16
- ReduceInt32
- ReduceInt
- ReduceInt64
- ReduceUint8
- ReduceUint16
- ReduceUint32
- ReduceUint
- ReduceUint64
- ReduceFloat32
- ReduceFloat64

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