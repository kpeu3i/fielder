package main

import (
	"fmt"
	"reflect"

	"github.com/kpeu3i/fielder/examples/simple/models"
)

func main() {
	fmt.Println("enum string:", models.UserAccountColumnID.String())                 // id
	fmt.Println("enum validation:", models.UserAccountColumn("_INVALID_").IsValid()) // false

	columns1 := models.NewUserAccountColumnList()
	columns2 := models.NewUserAccountColumnList()

	fmt.Println("collection equality (1):", columns1.Equals(columns2))    // true
	fmt.Println("collection similarity (1):", columns1.Similar(columns2)) // true

	// Move ID column to the end of the list.
	columns1.
		Remove(models.UserAccountColumnID).
		Add(models.UserAccountColumnID)

	fmt.Println("collection equality (2):", columns1.Equals(columns2))    // false
	fmt.Println("collection similarity (2):", columns1.Similar(columns2)) // true

	columns1.Clear()
	fmt.Println("collection length (1):", columns1.Len())                                        // 0
	fmt.Println("collection contains check (1):", columns1.Contains(models.UserAccountColumnID)) // false

	// Skip adding duplicates.
	columns1.
		AddIfNotContains(models.UserAccountColumnID).
		AddIfNotContains(models.UserAccountColumnID)

	fmt.Println("collection length (2):", columns1.Len())                                        // 1
	fmt.Println("collection contains check (2):", columns1.Contains(models.UserAccountColumnID)) // true

	fmt.Println("collection_1:", columns1.Strings()) // [id]
	fmt.Println("collection_2:", columns2.Strings()) // [id created_at updated_at name surname email password]

	r := reflect.ValueOf(&models.UserAccount{})
	f := reflect.Indirect(r).FieldByName("id")
	fmt.Println(f.String())
}
