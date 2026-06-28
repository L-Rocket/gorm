// User: I want to learn how to write GORM. I have a local postgres database and I want to connect to it using GORM. Can you provide me with a simple example of how to do that in Go?
// Assistant: Sure! Below is a simple example of how to connect to a local PostgreSQL database using GORM in Go. Make sure you have GORM and the PostgreSQL driver installed. You can install them using the following commands:
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres
// User: I've installed them now
// Assistant: Great! Now you can use the following code to connect to your local PostgreSQL database using GORM. Make sure to replace the connection parameters (username, password, dbname, host, and port) with your actual database credentials.

// wirte a demo for me below

package main

import (
	"fmt"

	"github.com/L-Rocket/gorm/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=lea password=xxx dbname=gorm_demo port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("I've successfully connected to local postgresql")
	}

	// You can now use the `db` object to interact with your database
	// For example, you can perform migrations, create tables, etc.
	db.AutoMigrate(&internal.User{})
	// this way, GORM will automatically create the "users" table in your database based on the User struct.
	// If User struct is changed, GORM will automatically update the table structure accordingly.
	// User: but where is the migration record? GORM will create a table called "gorm_migrations" in your database to keep track of the migrations that have been applied. This table will store information about the migration files that have been executed, allowing GORM to determine which migrations need to be run when you call AutoMigrate or other migration functions.
	// User: but I did not see the "gorm_migrations" table in my database. Why is that?
	// Assistant: The "gorm_migrations" table is not created by default when you use `AutoMigrate`. The `AutoMigrate` function is designed to automatically create or update tables based on your model definitions, but it does not track migrations in a separate table.
	// If you want to keep track of migrations, you would typically use a migration tool or library that integrates with GORM, such as `golang-migrate` or `gormigrate`. These tools allow you to define migration files and keep track of which migrations have been applied.

	// let's create a new user and save it to the database
	newUser := internal.User{Name: "John Doe", Age: 30}
	result := db.Create(&newUser)
	if result.Error != nil {
		panic("failed to create user")
	}
	fmt.Println("User created successfully, and the ID is:", newUser.ID)
	// User: does GORM automatically query the ID of the newly created user and assign it to the newUser.ID field?
	// Assistant: Yes, GORM automatically retrieves the ID of the newly created record and assigns it to the `ID` field of the struct you passed to the `Create` method. In this case, after calling `db.Create(&newUser)`, the `newUser.ID` field will contain the ID of the newly created user in the database.

	// let's query the user we just created
	var user internal.User
	db.First(&user, newUser.ID) // find user with ID
	fmt.Printf("User found: %+v\n", user)

	// 2. User: I want to learn how to write GORM. I have a local postgres database and I want to connect to it using GORM. Can you provide me with a simple example of how to do that in Go?

	db.Where("age > ?", 25).Find(&user) // find users with age greater than 25
	fmt.Printf("Users found: %+v\n", user)
	// why only one user is found? Because the `user` variable is a single instance of the `User` struct, so it can only hold one record. If you want to retrieve multiple users, you should use a slice of `User` structs instead. Here's how you can do that:
	var users []internal.User
	db.Where("age > ?", 25).Find(&users)
	fmt.Printf("Users found: %+v\n", users)

}
