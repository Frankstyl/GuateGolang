//Package data helps you to create data structures to handle user and article information
package data

//User struct keeps the user personal info. la estructura Usuario almacena la informacion personal de un usuario.
type User struct {
	Name  string //To save the name  para guardar el nombre
	Email string //To save the Email  para guardar el correo electronico
	Age   string //To save the Age para guardar la edad
	Phone string //to save the Phone number para guardar el numero de telefone
	Picture string //to save the url of the user's picture para guradar la url de la foto del usuario
}

//CreateUser makes a User
func CreateUser(name, email, age, phone, pictureurl string) User {
	a := User{name, email, age, phone, pictureurl}
	return a
}
