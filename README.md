# MySQL Query Builder for Golang.

## Build Query for MySQL

`MySQL Query Builder` is a lightweight and flexible query builder for MySQL in Go, designed to simplify the process of constructing SQL queries. This package enables developers to create complex queries with ease while ensuring clean and readable code.

## Use Cases

This package is suitable for you if you need to perform some queries on:

* Perform `INSERT`, `BULK INSERT` `UPDATE` and `DELETE` query
* Perform `SELECT` query
* Perform where query (`Where, OrWhere, WhereIn, WhereDate` e.t.c)
* Perform `Sorting`
* Perform `Limit, Offset`
* Perform Aggregate Query `(count, sum, avg, max, min)`
* Perform MySQL `transaction`

## Installation
```bash
go get github.com/ruhulfbr/go-mysql-qb
```

## Basic Usage

```go
package main

import (
	"fmt"
	"log"
	DB "github.com/ruhulfbr/go-mysql-qb"
)

func main() {
	// Connect to the database
	DB.ConnectDB("username", "password", "localhost:3306", "your_database_name")
	defer DB.CloseDB()

	// Create a new record
	qb := DB.Table("users")
	_, err := qb.Insert(map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   30,
	})
	if err != nil {
		log.Fatalf("Insert error: %v", err)
	}
	fmt.Println("Inserted new user.")

	// Update a record
	_, err = qb.Where("email", "=", "john@example.com").Update(map[string]interface{}{
		"name": "John Smith",
	})
	if err != nil {
		log.Fatalf("Update error: %v", err)
	}
	fmt.Println("Updated user name.")

	// Select records
	users, err := qb.Select("name", "email").Where("age", ">", 18).Get()
	if err != nil {
		log.Fatalf("Select error: %v", err)
	}
	fmt.Println("Users:", users)

	// Count records
	count, err := qb.Count()
	if err != nil {
		log.Fatalf("Count error: %v", err)
	}
	fmt.Printf("Total users: %d\n", count)

	// Delete a record
	_, err = qb.Where("email", "=", "john@example.com").Delete()
	if err != nil {
		log.Fatalf("Delete error: %v", err)
	}
	fmt.Println("Deleted user.")

	// Using aggregate functions
	sumAge, err := qb.Sum("age")
	if err != nil {
		log.Fatalf("Sum error: %v", err)
	}
	fmt.Printf("Total age of users: %.2f\n", sumAge)

	avgAge, err := qb.Avg("age")
	if err != nil {
		log.Fatalf("Avg error: %v", err)
	}
	fmt.Printf("Average age of users: %.2f\n", avgAge)

	maxAge, err := qb.Max("age")
	if err != nil {
		log.Fatalf("Max error: %v", err)
	}
	fmt.Printf("Maximum age of users: %.2f\n", maxAge)

	minAge, err := qb.Min("age")
	if err != nil {
		log.Fatalf("Min error: %v", err)
	}
	fmt.Printf("Minimum age of users: %.2f\n", minAge)

	// Bulk Insert
	_, err = qb.BulkInsert([]map[string]interface{}{
		{"name": "Alice", "email": "alice@example.com", "age": 28},
		{"name": "Bob", "email": "bob@example.com", "age": 35},
	})
	if err != nil {
		log.Fatalf("Bulk insert error: %v", err)
	}
	fmt.Println("Inserted multiple users.")
}


```

### Available where operators

* `=` (default operator, can be omitted)
* `>`
* `<`
* `<=`
* `>=`
* `!=`

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/ruhulfbr/go-mysql-qb/tree/main?tab=MIT-1-ov-file#readme) file for details.


## Support

If you found an issue or had an idea please refer [to this section](https://github.com/ruhulfbr/go-mysql-qb/issues).

## Authors

* **Md Ruhul Amin** - [Github](https://github.com/ruhulfbr)
