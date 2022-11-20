package main

import (
	"fmt"
	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type  Student struct {
	ID uint64
	Height uint8 // 150cm - 200cm
	Gene uint64
	Name string
}

var minHeight uint8 = 150
var maxHeight uint8 = 200

var dsn = "user:passwd@tcp(127.0.0.1:3306)/vsearchlogDB?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Errorf("open db error: %v", err))
		return
	}
	
	batchCnt := 0

	batchSize := 100
	var students []*Student
	totalCnt := 5000000
	for i := 0; i < totalCnt; i++ {
		students = append(students, &Student{
			Height: uint8(rand.Intn(int(maxHeight-minHeight)) + int(minHeight)),
			Gene: uint64(rand.Int63n(3000000)),
			Name: fmt.Sprintf("name|%v", rand.Int31n(1000000)),
		})
		batchCnt++
		if batchCnt == batchSize {
			result := db.Create(students)
			if result.Error != nil {
				fmt.Println(fmt.Errorf("db create error: %v", err))
				return
			}
			batchCnt = 0
			students = nil
			// fmt.Printf("inserted %v\n", i)
		}
	}
}
