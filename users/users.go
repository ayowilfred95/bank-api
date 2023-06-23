package users

import (
	"time"

	"github.com/ayowilfred95/helpers"
	"github.com/ayowilfred95/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


func prepareToken(user *interfaces.User) string {
	
	// sign up token

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)
	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount) map[string]interface{}{
	// setup response

	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	// prepare response
	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "username and password matches"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}

func Login(username string, pass string) map[string]interface{} {
	// add validation from helper

	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if valid {
		// connect to db
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}
	// verify pssword
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}
	// Find account

	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id,name,balance").Where("user_id=?", user.ID).Scan(&accounts)

	defer db.Close()

	var response = prepareResponse(user,accounts)
	return response
	}else {
return map[string]interface{}{"message":"not valid values"}
	}
	

}


func Register(username string, email string, pass string) map[string]interface{} {

	// add validation from helper

	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
		if valid {
db:= helpers.ConnectDB()
generatePassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatePassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Crypto Account", Name: string(username + "s" + "account"), Balance: 0, UserID: user.ID}

		db.Create(&account)

		defer db.Close()
		accounts:= []interfaces.ResponseAccount{}
			respAccount := interfaces.ResponseAccount{ID: account.ID, Name: account.Name, Balance: uint(account.Balance)}
			accounts = append(accounts, respAccount)
			var response =  prepareResponse(user, accounts)

			return response


		}else{
			return map[string]interface{}{"message":"not valid values,please enter match values"}
		}
}
