package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// About handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us Page")
}

// POST request handler
func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle POST request
		fmt.Fprintf(w, "POST request received")
	}
}

// GET request handler
func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle GET request
		fmt.Fprintf(w, "GET request received")
	}
}

// PUT request handler
func handlePut(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// Handle PUT request
		fmt.Fprintf(w, "PUT request received")
	}
}

// Query handler
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the URL
	query := r.URL.Query().Get("q")   // Get the 'q' parameter (search query)
	sort := r.URL.Query().Get("sort") // Get the 'sort' parameter (optional)

	// Prepare the response message
	var response string
	if query == "" {
		response = "No search query provided."
	} else {
		response = fmt.Sprintf("Searching for: %s", query)
	}

	if sort != "" {
		response += fmt.Sprintf(" | Sorted by: %s", sort)
	}

	// Send the response
	fmt.Fprintf(w, response)
}

// Form handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Retrieve the form value
		name := r.FormValue("name")

		// Send a response with the submitted name
		fmt.Fprintf(w, "Form submitted with name: %s", name)
	} else {
		// Display a simple HTML form
		fmt.Fprintf(w, `
			<form action="/submit" method="post">
				<input type="text" name="name" placeholder="Enter your name" />
				<button type="submit">Submit</button>
			</form>
		`)
	}
}

// Handler for URL path parameters
func pathParamHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the item ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/items/")

	// Send a response with the extracted item ID
	fmt.Fprintf(w, "Item ID: %s", id)
}

// JSON handler
type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items = make(map[int]Item)
var idCounter = 1
var mu sync.Mutex

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listItems(w)
	case http.MethodPost:
		createItem(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getItem(w, id)
	case http.MethodPut:
		updateItem(w, r, id)
	case http.MethodDelete:
		deleteItem(w, id)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// List all items
func listItems(w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()

	var result []Item
	for _, item := range items {
		result = append(result, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	itemID := idCounter
	idCounter++
	items[itemID] = item
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": itemID})
}

// Get an item by ID
func getItem(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	item, exists := items[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Update an item by ID
func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	item, exists := items[id]
	if !exists {
		mu.Unlock()
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	updatedItem.Name = item.Name // Preserve the name if not provided in the request
	items[id] = updatedItem
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedItem)
}

// Delete an item by ID
func deleteItem(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	_, exists := items[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(items, id)
	w.WriteHeader(http.StatusNoContent)
}

// Error handler
func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Custom error message", http.StatusInternalServerError)
}

// Logging middleware
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/post", handlePost)
	mux.HandleFunc("/get", handleGet)
	mux.HandleFunc("/put", handlePut)
	http.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/error", errorHandler)
	http.HandleFunc("/items", itemsHandler)
    http.HandleFunc("/items/", itemHandler)

	loggedMux := loggingHandler(mux)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
