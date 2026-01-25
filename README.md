# Kasir API

A simple REST API for managing products (produk) in a cashier system, built with Go. This API provides CRUD operations for product management with in-memory storage.

## Features

- ✅ List all products
- ✅ Get product by ID
- ✅ Create new product
- ✅ Update existing product
- ✅ Delete product
- ✅ List all categories
- ✅ Get category by ID
- ✅ Create new category
- ✅ Update existing category
- ✅ Delete category
- ✅ Health check endpoint

## Prerequisites

- Go 1.25.5 or higher
- A terminal/command line interface

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd go-kasir-api
```

2. Build the application:

```bash
go build -o kasir-api main.go
```

3. Run the application:

```bash
./kasir-api
```

Or run directly with Go:

```bash
go run main.go
```

The server will start on `localhost:8887`

## API Endpoints

### Health Check

**GET** `/health`

Check if the API is running.

**Response:**

```json
{
  "status": "OK",
  "message": "API Running"
}
```

### Get All Products

**GET** `/api/produk`

Retrieve all products.

**Response:**

```json
[
  {
    "id": 1,
    "nama": "Indomie Godog",
    "harga": 3500,
    "stok": 10
  },
  {
    "id": 2,
    "nama": "Vit 1000ml",
    "harga": 3000,
    "stok": 40
  }
]
```

### Get Product by ID

**GET** `/api/produk/{id}`

Retrieve a specific product by its ID.

**Response:**

```json
{
  "id": 1,
  "nama": "Indomie Godog",
  "harga": 3500,
  "stok": 10
}
```

**Error Response (404):**

```
Produk belum ada
```

### Create Product

**POST** `/api/produk`

Create a new product.

**Request Body:**

```json
{
  "nama": "Kecap Manis",
  "harga": 12000,
  "stok": 20
}
```

**Response (201):**

```json
{
  "id": 4,
  "nama": "Kecap Manis",
  "harga": 12000,
  "stok": 20
}
```

### Update Product

**PUT** `/api/produk/{id}`

Update an existing product.

**Request Body:**

```json
{
  "nama": "Indomie Goreng",
  "harga": 4000,
  "stok": 15
}
```

**Response:**

```json
{
  "id": 1,
  "nama": "Indomie Goreng",
  "harga": 4000,
  "stok": 15
}
```

**Error Response (404):**

```
Produk belum ada
```

### Delete Product

**DELETE** `/api/produk/{id}`

Delete a product by ID.

**Response:**

```json
{
  "message": "Berhasil hapus produk"
}
```

**Error Response (404):**

```
Produk belum ada
```

### Get All Categories

**GET** `/api/categories`

Retrieve all categories.

**Response:**

```json
[
  {
    "id": 1,
    "name": "Makanan Ringan",
    "description": "Snack dan camilan"
  },
  {
    "id": 2,
    "name": "Minuman",
    "description": "Berbagai jenis minuman"
  }
]
```

### Get Category by ID

**GET** `/api/categories/{id}`

Retrieve a specific category by its ID.

**Response:**

```json
{
  "id": 1,
  "name": "Makanan Ringan",
  "description": "Snack dan camilan"
}
```

**Error Response (404):**

```
Category belum ada
```

### Create Category

**POST** `/api/categories`

Create a new category.

**Request Body:**

```json
{
  "name": "Alat Tulis",
  "description": "Perlengkapan tulis kantor"
}
```

**Response (201):**

```json
{
  "id": 4,
  "name": "Alat Tulis",
  "description": "Perlengkapan tulis kantor"
}
```

### Update Category

**PUT** `/api/categories/{id}`

Update an existing category.

**Request Body:**

```json
{
  "name": "Makanan Berat",
  "description": "Nasi dan lauk pauk"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Makanan Berat",
  "description": "Nasi dan lauk pauk"
}
```

**Error Response (404):**

```
Category belum ada
```

### Delete Category

**DELETE** `/api/categories/{id}`

Delete a category by ID.

**Response:**

```json
{
  "message": "Berhasil hapus category"
}
```

**Error Response (404):**

```
Category belum ada
```

## Example Usage

### Using cURL

**Get all products:**

```bash
curl http://localhost:8887/api/produk
```

**Get product by ID:**

```bash
curl http://localhost:8887/api/produk/1
```

**Create a product:**

```bash
curl -X POST http://localhost:8887/api/produk \
  -H "Content-Type: application/json" \
  -d '{"nama": "Susu UHT", "harga": 5000, "stok": 30}'
```

**Update a product:**

```bash
curl -X PUT http://localhost:8887/api/produk/1 \
  -H "Content-Type: application/json" \
  -d '{"nama": "Indomie Goreng", "harga": 4000, "stok": 15}'
```

**Delete a product:**

```bash
curl -X DELETE http://localhost:8887/api/produk/1
```

**Health check:**

```bash
curl http://localhost:8887/health
```

### Using cURL for Categories

**Get all categories:**

```bash
curl http://localhost:8887/api/categories
```

**Get category by ID:**

```bash
curl http://localhost:8887/api/categories/1
```

**Create a category:**

```bash
curl -X POST http://localhost:8887/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Elektronik", "description": "Barang elektronik rumah tangga"}'
```

**Update a category:**

```bash
curl -X PUT http://localhost:8887/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Minuman Dingin", "description": "Minuman dingin segar"}'
```

**Delete a category:**

```bash
curl -X DELETE http://localhost:8887/api/categories/1
```

## Project Structure

```
go-kasir-api/
├── main.go          # Main application file with all handlers
├── go.mod           # Go module file
├── .gitignore       # Git ignore file
└── README.md        # This file
```

## Data Model

### Product (Produk)

```go
type Produk struct {
    ID    int    `json:"id"`    // Product ID (auto-generated)
    Nama  string `json:"nama"`  // Product name
    Harga int    `json:"harga"` // Product price
    Stok  int    `json:"stok"`  // Product stock
}
```

### Category

```go
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
```

## Default Data

The API comes with the following default data:

### Products

1. **Indomie Godog** - Harga: 3,500 - Stok: 10
2. **Vit 1000ml** - Harga: 3,000 - Stok: 40
3. **Kecap** - Harga: 12,000 - Stok: 20

### Categories

1. **Makanan Ringan** - Snack dan camilan
2. **Minuman** - Berbagai jenis minuman
3. **Bumbu Dapur** - Bumbu dan rempah-rempah

## Configuration

The server runs on:

- **Host:** `localhost`
- **Port:** `8887`
- **Full URL:** `http://localhost:8887`

To change the port, modify the `PORT` variable in `main.go`.

## Current Limitations

- ⚠️ Data is stored in-memory and will be lost when the server restarts
- ⚠️ No database persistence (planned for future updates)
- ⚠️ No authentication/authorization
- ⚠️ No input validation beyond basic JSON parsing

## Future Improvements

- [ ] Add database persistence (PostgreSQL/MySQL/MongoDB)
- [ ] Add input validation
- [ ] Add authentication and authorization
- [ ] Add request logging
- [ ] Add error handling middleware
- [ ] Add API documentation (Swagger/OpenAPI)
- [ ] Add unit tests
- [ ] Add Docker support
- [ ] Add environment variable configuration

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK` - Successful GET, PUT, DELETE operations
- `201 Created` - Successful POST operation
- `400 Bad Request` - Invalid request format or invalid ID
- `404 Not Found` - Product not found

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
