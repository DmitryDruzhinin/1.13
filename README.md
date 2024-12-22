# This is a web service that allows users to send arithmetic expressions and receive the results of their calculation. 

# Launching the service.

```
git clone https://github.com/DmitryDruzhinin/1.13.git
cd 1.13
```
```
go run ./cmd/main.go
```

<br>

# Example

Request: 

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

Answer:

```json
{
    "result": "6.000000"
}
```

<br>


# Possible results:

1. Successful request

   Status: `200 OK`  
   Answer:

   ```json
   {
     "result": "6"
   }
   ```

2. Expression is not valid

   Status: `422`  
   Answer:

   ```json
   {
     "error": "Expression is not valid"
   }
   ```

3. Internal server error

   Status: `500`  
   Answer:

   ```json
   {
     "error": "Internal server error"
   }
   ```
