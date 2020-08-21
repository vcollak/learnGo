

### Get all books
curl -L -X GET 'http://localhost:8000/api/books' --data-raw ''

### Get book by ID
curl -L -X GET 'http://localhost:8000/api/books/1' --data-raw ''

### Create a book
curl -L -X POST 'http://localhost:8000/api/books' -H 'Content-Type: application/json' --data-raw '{
    "id": "1",
    "isbn": "11221122",
    "title": "Measure what Matters",
    "author": {
        "firstname": "John",
        "lastname": "Doer"
    }
}'


### Update a book
curl -L -X PUT 'http://localhost:8000/api/books/1' -H 'Content-Type: application/json' --data-raw '{
    "id": "1",
    "isbn": "11221122",
    "title": "Measure what Matters",
    "author": {
        "firstname": "John",
        "lastname": "Doerr"
    }
}'


### Delete a book
curl -L -X DELETE 'http://localhost:8000/api/books/1' --data-raw ''
