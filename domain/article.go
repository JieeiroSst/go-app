package domain

type Article struct {
	ID     int  `gorm:"primaryKey"`
	Title  string
	Body   string
	AuthorId int 
	Author Author
}

type Author struct {
	ID       int  `gorm:"primaryKey"`
	Name     string
}

type Repository interface {
	CreateArticle(Article) (err error)
	ArticleAll() (articles []Article, err error)
	ArticleById(int) (article Article, err error)
	UpdateArticle(int, Article) (err error)
	DeleteArticle(int) (err error)
	
	CreateAuthor(Author) (err error)
	AuthorAll() (authors []Author, err error)
	AuthorById(int) (author Author, err error)
	UpdateAuthor(int, Author) (err error)
	DeleteAuthor(int) (err error)
}
