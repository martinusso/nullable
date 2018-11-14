# nullable

[![Build Status](https://travis-ci.org/martinusso/nullable.svg?branch=master)](https://travis-ci.org/martinusso/nullable)
[![Coverage Status](https://coveralls.io/repos/github/martinusso/nullable/badge.svg?branch=master)](https://coveralls.io/github/martinusso/nullable?branch=master)
[![GoDoc](https://godoc.org/github.com/martinusso/nullable?status.svg)](https://godoc.org/github.com/martinusso/nullable)
[![Go Report Card](https://goreportcard.com/badge/github.com/martinusso/nullable)](https://goreportcard.com/report/github.com/martinusso/nullable)

Package nullable provides a simple way to marshal/unmarshal Go structs from sql.Null* types.

## Installation

```
go get -u github.com/martinusso/nullable
```

## Why

```
var person struct {
	Name        sql.NullString
	Age         sql.NullInt64
	Married     sql.NullBool
	Height      sql.NullFloat64
}

got, _ := json.Marshal(person)
fmt.Println(string(got))
```

Output:
```JSON
{
   "Name":{
      "String":"",
      "Valid":false
   },
   "Age":{
      "Int64":0,
      "Valid":false
   },
   "Married":{
      "Bool":false,
      "Valid":false
   },
   "Height":{
      "Float64":0,
      "Valid":false
   }
}
```

## Usage

```go
type Person struct {
	Name        String
	Age         Int64
	Married     Bool
	Height      Float64
}
```

### json.Marshaler

```go
var person Person
got, _ := json.Marshal(person)
```

Output:
```JSON
{
	"Name": "John Doe",
	"Age": 42,
	"Married": true,
	"Height": 1.79
}
```

### json.Unmarshaler

```go
body := []byte(`{"Name":"John Doe","Age":42,"Married":true,"Height":1.79}`)
var person Person
json.Unmarshal(body, &person)
```

Output:
```go
(main.Person) {
 Name: (nullable.String) {
  NullString: (sql.NullString) {
   String: (string) (len=8) "John Doe",
   Valid: (bool) true
  }
 },
 Age: (nullable.Int64) {
  NullInt64: (sql.NullInt64) {
   Int64: (int64) 42,
   Valid: (bool) true
  }
 },
 Married: (nullable.Bool) {
  NullBool: (sql.NullBool) {
   Bool: (bool) true,
   Valid: (bool) true
  }
 },
 Height: (nullable.Float64) {
  NullFloat64: (sql.NullFloat64) {
   Float64: (float64) 1.79,
   Valid: (bool) true
  }
 }
}
```

## License

This software is open source, licensed under the The MIT License (MIT). See [LICENSE](https://github.com/martinusso/nullable/blob/master/LICENSE) for details.
