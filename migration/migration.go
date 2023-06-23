package migration

import (
	"log"

	"github.com/ayowilfred95/interfaces"
	"github.com/ayowilfred95/helpers"
)




func createAccounts() {
	db := helpers.ConnectDB()

	users := &[4]interfaces.User{
		{Username: "Martins", Email: "martins@gmail.com", Password: "welcomehome"},
		{Username: "mojisola", Email: "mojisola@gmail.com", Password: "mojisola123"},
		{Username: "wilfred", Email: "wilfred@gmail.com", Password: "wilfred"},
	}
	//db.Create(&user)
	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatePassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Crypto Account", Name: string(users[i].Username + "s" + "account"), Balance: uint(100000 * int(i+1)), UserID: user.ID}

		db.Create(&account)

	}

	defer db.Close()
}

func Migrate() {
	//User := &interfaces.User{}
	//Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&interfaces.User{}, &interfaces.Account{})
	defer db.Close()

	createAccounts()

	log.Println("Created new account successfully")

}
