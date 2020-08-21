

### To add more items 
curl -L -X POST 'localhost:8080/todos' -H 'Content-Type: application/json' --data-raw '
    {      
        "name": "Task #123",
        "completed": true,
        "due": "2020-07-27T20:34:58.651387237Z"
    }
'