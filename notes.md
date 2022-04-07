# routes 
- getBooks
  - /book/ => GET
- createBook
  - /book/ => POST
- getBookById
  - /book/{bookId} => GET
- updateBook
  - /book/{bookId} => PUT
- deleteBook
  - /book/{bookId} => DELETE

# technologies
- go
- gorm
- mySQL 
- json marshall, unmarshall
- gorilla mux

# folder structure
- CMD => main.go
- pkg => 
  - config => app.go
  - controllers => book-controller
  - models => book.go
  - routes => bookstore-routes
  - utils => utils.go

# roadmap
- routes
- 
- models
- book-controller