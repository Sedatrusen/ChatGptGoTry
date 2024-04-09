package repository

import "github.com/your-username/product-category-service/internal/models"

// CategoryRepository provides methods to interact with the category data store
type CategoryRepository struct {
	// Add necessary dependencies here
}

// CreateCategory creates a new product category
func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	// Implement logic to create a product category
}

// GetCategoryByID retrieves a product category by ID
func (r *CategoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	// Implement logic to get a product category by ID
}

// UpdateCategory updates an existing product category
func (r *CategoryRepository) UpdateCategory(category *models.Category) error {
	// Implement logic to update a product category
}

// DeleteCategoryByID deletes a product category by ID
func (r *CategoryRepository) DeleteCategoryByID(id int) error {
	// Implement logic to delete a product category by ID
}
