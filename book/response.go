package book

type BookResponse struct {
	ID          int    `josn:"id"`
	Title       string `josn:"title"`
	Price       int    `josn:"price"`
	Description string `josn:"description"`
	Rating      int    `josn:"rating"`
	Discount    int    `josn:"discount"`
}
