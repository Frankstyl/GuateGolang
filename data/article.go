//Package data helps you to create data structures to handle user and article information
package data

//Article is used to build a structure of type Article
type Article struct {
	Name    string //name of the article
	Price   string //Price of the article
	Picture string //path to picture
}

//CreateArticle Generate an Article
func CreateArticle(name, price, picturepath string) Article {
	a := Article{name, price, picturepath}
	return a
}
