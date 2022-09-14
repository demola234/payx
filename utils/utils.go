package utils

import(

    "math/rand"
)


func GenerateRandomString(length int, condition int ) (string){

	var letters [] rune
	if(condition == 0){
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	}
	if(condition == 1){
		letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		}
	if(condition == 2){
		letters = []rune("0123456789")
		}
	if(condition == 3){
		letters = []rune("abcdefghijklmnopqrstuvwxyz")
		}
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
}