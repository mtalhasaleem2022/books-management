# Books API

REST API for managing books with Go, Gin, PostgreSQL, Redis, and Kafka.

## Features
- CRUD operations for books
- Redis caching
- Kafka event streaming
- Swagger documentation
- Docker support

## Setup
1. Start dependencies:
```bash
docker-compose up -d

**## Public API Endpoints**
Here are the API endpoints for the Books Management System:

**Method 	Endpoint	 Description**
1. GET  	  /api/v1/books	        Retrieve a list of all books (with pagination).
2. GET	    /api/v1/books/{id}	  Retrieve a specific book by its ID.
3. POST  	  /api/v1/books	        Create a new book.
4. PUT	    /api/v1/books/{id}  	Update an existing book by its ID.
5. DELETE	  /api/v1/books/{id}	  Delete a book by its ID.


**## cURL commands**

```
1. Get All Books

curl -X GET "http://localhost:8080/api/v1/books"

```

2. Create a Book
```bash

curl -X POST "http://localhost:8080/api/v1/books" \
-H "Content-Type: application/json" \
-d '{
  "title": "New Book",
  "author": "Author Name",
  "year": 2023
}'

3. Update a Book
```bash

curl -X PUT "http://localhost:8080/api/v1/books/1" \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Book",
  "author": "Updated Author",
  "year": 2024
}'

4. Delete a Book
```bash

curl -X DELETE "http://localhost:8080/api/v1/books/1"


**##** **Example Responses**

1. Get All Books
[
  {
    "id": 1,
    "title": "Test Book 1",
    "author": "Author 1",
    "year": 2023
  },
  {
    "id": 2,
    "title": "Test Book 2",
    "author": "Author 2",
    "year": 2022
  }
]

```
2. Get Book by ID

{
  "id": 1,
  "title": "Test Book 1",
  "author": "Author 1",
  "year": 2023
}

```

3. Create a Book

{
  "id": 3,
  "title": "New Book",
  "author": "Author Name",
  "year": 2023
}

```

4. Update a Book

{
  "id": 1,
  "title": "Updated Book",
  "author": "Updated Author",
  "year": 2024
}

```

5. Delete a Book
Status Code: 204 No Content

```


## Postman Collection
Below is the **Postman Collection JSON** that you can import into Postman for testing:

Postman Collection JSON

{
	"info": {
		"_postman_id": "your-postman-id",
		"name": "Books Management API",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Get All Books",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/v1/books",
					"protocol": "http",
					"host": ["localhost"],
					"port": "8080",
					"path": ["api", "v1", "books"]
				}
			}
		},
		{
			"name": "Get Book by ID",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/v1/books/1",
					"protocol": "http",
					"host": ["localhost"],
					"port": "8080",
					"path": ["api", "v1", "books", "1"]
				}
			}
		},
		{
			"name": "Create a Book",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"title\": \"New Book\",\n  \"author\": \"Author Name\",\n  \"year\": 2023\n}"
				},
				"url": {
					"raw": "http://localhost:8080/api/v1/books",
					"protocol": "http",
					"host": ["localhost"],
					"port": "8080",
					"path": ["api", "v1", "books"]
				}
			}
		},
		{
			"name": "Update a Book",
			"request": {
				"method": "PUT",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"title\": \"Updated Book\",\n  \"author\": \"Updated Author\",\n  \"year\": 2024\n}"
				},
				"url": {
					"raw": "http://localhost:8080/api/v1/books/1",
					"protocol": "http",
					"host": ["localhost"],
					"port": "8080",
					"path": ["api", "v1", "books", "1"]
				}
			}
		},
		{
			"name": "Delete a Book",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/api/v1/books/1",
					"protocol": "http",
					"host": ["localhost"],
					"port": "8080",
					"path": ["api", "v1", "books", "1"]
				}
			}
		}
	]
}

```
**How to Import into Postman**
1. Open Postman.

2. Click on Import in the top-left corner.

3. Select Raw Text and paste the JSON above.

4. Click Import.
```
