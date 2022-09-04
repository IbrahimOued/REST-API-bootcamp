package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

// CRUD Functions

// Create product
func CreateProduct(products *[]Product, newProduct *Product) (err error) {
	*products = append(*products, *newProduct)
	return nil
}

// View product
func ReadProductById(products *[]Product, id int) (product Product) {
	for _, item := range *products {
		if item.Id == id {
			return item
		}
	}
	// otherwise we return the product with empty data
	return Product{}
}

func UpdateProductById(products *[]Product, updatedProduct *Product) (product Product) {
	for index, item := range *products {
		if item.Id == (*updatedProduct).Id {
			(*products)[index].Name = (*updatedProduct).Name
			(*products)[index].Stock = (*updatedProduct).Stock
			(*products)[index].Price = (*updatedProduct).Price

			return (*products)[index]
		}

	}
	return Product{}
}

func DeleteProductById(products *[]Product, id int) (product Product) {
	var deletedProduct Product
	for index, item := range *products {
		if item.Id == id {
			n := len(*products)
			deletedProduct = item
			(*products)[index] = (*products)[n-1]
			*products = (*products)[:n-1]

			return deletedProduct
		}
	}
	return Product{}
}
