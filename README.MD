# gobodyparser

## Example

```go 
func main() {
	router := mux.NewRouter() // create router

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		// application/json request {{ "firstName" : "John", "lastName" : "Doe" }}
		bodyParser := gobodyparser.BodyParser(request.Body)

		for key, value := range bodyParser {
			fmt.Println(key, value)
		}
	}).Methods("POST")
}
```