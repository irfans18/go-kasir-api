package server

// In-memory storage (sementara, nanti ganti database)
var ProdukData = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

var CategoriesData = []Category{
	{ID: 1, Name: "Makanan Ringan", Description: "Snack dan camilan"},
	{ID: 2, Name: "Minuman", Description: "Berbagai jenis minuman"},
	{ID: 3, Name: "Bumbu Dapur", Description: "Bumbu dan rempah-rempah"},
}
