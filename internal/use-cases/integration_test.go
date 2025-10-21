package use_cases

import (
	"testing"
	"time"

	"gin-quickstart/internal/repositories"
)

// Testes de integração completos entre Create e List Use Cases
func TestIntegration_CreateAndListCategories(t *testing.T) {
	// Arrange
	repo := repositories.NewInMemotyCategoryRepository()
	createUseCase := NewCreateCategoryUseCase(repo)
	listUseCase := NewListCategoriesUseCase(repo)

	// Test cases
	testCategories := []struct {
		name        string
		expectError bool
	}{
		{"Electronics & Gadgets", false},
		{"Books & Literature", false},
		{"Home & Garden", false},
		{"", true}, // Nome inválido - muito curto
		{"abc", true}, // Nome inválido - muito curto
		{"Sports & Recreation", false},
	}

	var validCategories []string

	// Act & Assert - Criar categorias
	for _, tc := range testCategories {
		err := createUseCase.Execute(tc.name)
		
		if tc.expectError {
			if err == nil {
				t.Errorf("Expected error for category '%s', got nil", tc.name)
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error for category '%s', got %v", tc.name, err)
			} else {
				validCategories = append(validCategories, tc.name)
			}
		}
	}

	// Act - Listar todas as categorias
	categories, err := listUseCase.Execute()

	// Assert - Verificar lista
	if err != nil {
		t.Fatalf("Error listing categories: %v", err)
	}

	if len(categories) != len(validCategories) {
		t.Fatalf("Expected %d categories, got %d", len(validCategories), len(categories))
	}

	// Verificar se todas as categorias válidas estão na lista
	for i, category := range categories {
		if category.Name != validCategories[i] {
			t.Errorf("Expected category %s at position %d, got %s", 
				validCategories[i], i, category.Name)
		}

		// Verificar se os timestamps foram definidos
		if category.CreatedAt.IsZero() {
			t.Errorf("Category %s has zero CreatedAt", category.Name)
		}

		if category.UpdatedAt.IsZero() {
			t.Errorf("Category %s has zero UpdatedAt", category.Name)
		}

		// Verificar se CreatedAt e UpdatedAt são próximos (nova categoria)
		timeDiff := category.UpdatedAt.Sub(category.CreatedAt)
		if timeDiff < 0 {
			timeDiff = -timeDiff
		}
		if timeDiff > time.Millisecond {
			t.Errorf("Category %s: CreatedAt and UpdatedAt should be close for new categories (diff: %v)", 
				category.Name, timeDiff)
		}
	}
}

// Teste de workflow completo: Create -> List -> Verificar ordem
func TestIntegration_CategoryOrder(t *testing.T) {
	// Arrange
	repo := repositories.NewInMemotyCategoryRepository()
	createUseCase := NewCreateCategoryUseCase(repo)
	listUseCase := NewListCategoriesUseCase(repo)

	categoryNames := []string{
		"First Category",
		"Second Category", 
		"Third Category",
	}

	// Act - Criar categorias em ordem
	for _, name := range categoryNames {
		err := createUseCase.Execute(name)
		if err != nil {
			t.Fatalf("Error creating category %s: %v", name, err)
		}
	}

	// Act - Listar categorias
	categories, err := listUseCase.Execute()
	if err != nil {
		t.Fatalf("Error listing categories: %v", err)
	}

	// Assert - Verificar ordem
	if len(categories) != len(categoryNames) {
		t.Fatalf("Expected %d categories, got %d", len(categoryNames), len(categories))
	}

	for i, category := range categories {
		if category.Name != categoryNames[i] {
			t.Errorf("Category at position %d: expected %s, got %s", 
				i, categoryNames[i], category.Name)
		}
	}
}

// Teste com repositório vazio
func TestIntegration_EmptyRepository(t *testing.T) {
	// Arrange
	repo := repositories.NewInMemotyCategoryRepository()
	listUseCase := NewListCategoriesUseCase(repo)

	// Act
	categories, err := listUseCase.Execute()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if categories == nil {
		t.Fatal("Expected empty slice, got nil")
	}

	if len(categories) != 0 {
		t.Fatalf("Expected empty list, got %d categories", len(categories))
	}
}

// Teste de multiple creates seguidos de list
func TestIntegration_MultipleCreatesAndList(t *testing.T) {
	// Arrange
	repo := repositories.NewInMemotyCategoryRepository()
	createUseCase := NewCreateCategoryUseCase(repo)
	listUseCase := NewListCategoriesUseCase(repo)

	// Act - Criar múltiplas categorias
	for i := 1; i <= 5; i++ {
		err := createUseCase.Execute("Category " + string(rune('0'+i)))
		if err != nil {
			// Nome muito curto, vamos usar nome válido
			err = createUseCase.Execute("Valid Category " + string(rune('0'+i)))
			if err != nil {
				t.Fatalf("Error creating category %d: %v", i, err)
			}
		}
	}

	// Act - Listar múltiplas vezes para verificar consistência
	for i := 0; i < 3; i++ {
		categories, err := listUseCase.Execute()
		if err != nil {
			t.Fatalf("Error listing categories (iteration %d): %v", i+1, err)
		}

		if len(categories) != 5 {
			t.Fatalf("Iteration %d: expected 5 categories, got %d", i+1, len(categories))
		}
	}
}