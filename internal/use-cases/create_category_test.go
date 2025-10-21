package use_cases

import (
	"fmt"
	"strings"
	"testing"

	"gin-quickstart/internal/entities"
	"gin-quickstart/internal/repositories"
)

// Mock repository para testes unitários
type mockCategoryRepository struct {
	savedCategories []*entities.Category
	saveError      error
	listError      error
}

func (m *mockCategoryRepository) Save(category *entities.Category) error {
	if m.saveError != nil {
		return m.saveError
	}
	m.savedCategories = append(m.savedCategories, category)
	return nil
}

func (m *mockCategoryRepository) List() ([]*entities.Category, error) {
	if m.listError != nil {
		return nil, m.listError
	}
	return m.savedCategories, nil
}

func TestCreateCategoryUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRepo := &mockCategoryRepository{}
	useCase := NewCreateCategoryUseCase(mockRepo)
	categoryName := "Test Category"

	// Act
	err := useCase.Execute(categoryName)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(mockRepo.savedCategories) != 1 {
		t.Fatalf("Expected 1 category to be saved, got %d", len(mockRepo.savedCategories))
	}

	savedCategory := mockRepo.savedCategories[0]
	if savedCategory.Name != categoryName {
		t.Errorf("Expected category name %s, got %s", categoryName, savedCategory.Name)
	}

	if savedCategory.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}

	if savedCategory.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestCreateCategoryUseCase_Execute_InvalidName(t *testing.T) {
	// Arrange
	mockRepo := &mockCategoryRepository{}
	useCase := NewCreateCategoryUseCase(mockRepo)
	invalidName := "abc" // Nome muito curto (< 5 caracteres)

	// Act
	err := useCase.Execute(invalidName)

	// Assert
	if err == nil {
		t.Fatal("Expected error for invalid category name, got nil")
	}

	expectedErrorMsg := "name must be greater than 5, got 3"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, err.Error())
	}

	if len(mockRepo.savedCategories) != 0 {
		t.Errorf("Expected no categories to be saved, got %d", len(mockRepo.savedCategories))
	}
}

func TestCreateCategoryUseCase_Execute_RepositoryError(t *testing.T) {
	// Arrange
	expectedError := fmt.Errorf("repository error") // Erro simulado
	mockRepo := &mockCategoryRepository{
		saveError: expectedError,
	}
	useCase := NewCreateCategoryUseCase(mockRepo)
	categoryName := "Valid Category Name"

	// Act
	err := useCase.Execute(categoryName)

	// Assert
	if err == nil {
		t.Fatal("Expected repository error, got nil")
	}

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// Teste de integração usando o repositório in-memory real
func TestCreateCategoryUseCase_Integration(t *testing.T) {
	// Arrange
	repo := repositories.NewInMemotyCategoryRepository()
	useCase := NewCreateCategoryUseCase(repo)
	categoryName := "Integration Test Category"

	// Act
	err := useCase.Execute(categoryName)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verificar se a categoria foi realmente salva
	categories, err := repo.List()
	if err != nil {
		t.Fatalf("Error listing categories: %v", err)
	}

	if len(categories) != 1 {
		t.Fatalf("Expected 1 category, got %d", len(categories))
	}

	if categories[0].Name != categoryName {
		t.Errorf("Expected category name %s, got %s", categoryName, categories[0].Name)
	}
}

// Benchmark para medir performance
func BenchmarkCreateCategoryUseCase_Execute(b *testing.B) {
	mockRepo := &mockCategoryRepository{}
	useCase := NewCreateCategoryUseCase(mockRepo)
	categoryName := "Benchmark Category"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = useCase.Execute(categoryName)
	}
}

// Teste de table-driven para múltiplos cenários
func TestCreateCategoryUseCase_TableDriven(t *testing.T) {
	tests := []struct {
		name          string
		categoryName  string
		expectError   bool
		errorContains string
	}{
		{
			name:         "Valid category name",
			categoryName: "Valid Category",
			expectError:  false,
		},
		{
			name:          "Empty name",
			categoryName:  "",
			expectError:   true,
			errorContains: "name must be greater than 5",
		},
		{
			name:          "Short name",
			categoryName:  "abc",
			expectError:   true,
			errorContains: "name must be greater than 5",
		},
		{
			name:         "Long valid name",
			categoryName: "This is a very long category name that should be valid",
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockCategoryRepository{}
			useCase := NewCreateCategoryUseCase(mockRepo)

			err := useCase.Execute(tt.categoryName)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				if tt.errorContains != "" && err != nil {
					if !strings.Contains(err.Error(), tt.errorContains) {
						t.Errorf("Expected error to contain '%s', got '%s'", tt.errorContains, err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			}
		})
	}
}