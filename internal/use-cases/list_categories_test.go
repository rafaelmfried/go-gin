package use_cases

import (
	"fmt"
	"testing"
	"time"

	"gin-quickstart/internal/entities"
	"gin-quickstart/internal/repositories"
)

func TestListCategoriesUseCase_Execute_Success(t *testing.T) {
	// Arrange
	mockRepo := &mockCategoryRepository{
		savedCategories: createMockCategories(),
	}
	useCase := NewListCategoriesUseCase(mockRepo)

	// Act
	categories, err := useCase.Execute()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(categories) != 3 {
		t.Fatalf("Expected 3 categories, got %d", len(categories))
	}

	// Verificar se as categorias estão corretas
	expectedNames := []string{"Category 1", "Category 2", "Category 3"}
	for i, category := range categories {
		if category.Name != expectedNames[i] {
			t.Errorf("Expected category name %s, got %s", expectedNames[i], category.Name)
		}
	}
}

func TestListCategoriesUseCase_Execute_EmptyList(t *testing.T) {
	// Arrange
	mockRepo := &mockCategoryRepository{
		savedCategories: []*entities.Category{}, // Lista vazia
	}
	useCase := NewListCategoriesUseCase(mockRepo)

	// Act
	categories, err := useCase.Execute()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(categories) != 0 {
		t.Fatalf("Expected empty list, got %d categories", len(categories))
	}

	if categories == nil {
		t.Error("Expected empty slice, got nil")
	}
}

func TestListCategoriesUseCase_Execute_RepositoryError(t *testing.T) {
	// Arrange
	expectedError := fmt.Errorf("database connection error")
	mockRepo := &mockCategoryRepository{
		listError: expectedError,
	}
	useCase := NewListCategoriesUseCase(mockRepo)

	// Act
	categories, err := useCase.Execute()

	// Assert
	if err == nil {
		t.Fatal("Expected repository error, got nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if categories != nil {
		t.Error("Expected nil categories on error, got non-nil")
	}
}

// Teste de integração usando o repositório in-memory real
func TestListCategoriesUseCase_Integration(t *testing.T) {
	// Arrange
	repo := repositories.NewInMemotyCategoryRepository()
	listUseCase := NewListCategoriesUseCase(repo)
	createUseCase := NewCreateCategoryUseCase(repo)

	// Adicionar algumas categorias
	categoryNames := []string{"Electronics", "Books", "Clothing"}
	for _, name := range categoryNames {
		err := createUseCase.Execute(name)
		if err != nil {
			t.Fatalf("Error creating category %s: %v", name, err)
		}
	}

	// Act
	categories, err := listUseCase.Execute()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(categories) != len(categoryNames) {
		t.Fatalf("Expected %d categories, got %d", len(categoryNames), len(categories))
	}

	// Verificar se todas as categorias estão presentes
	for i, category := range categories {
		if category.Name != categoryNames[i] {
			t.Errorf("Expected category name %s, got %s", categoryNames[i], category.Name)
		}

		if category.CreatedAt.IsZero() {
			t.Errorf("Category %s has zero CreatedAt", category.Name)
		}

		if category.UpdatedAt.IsZero() {
			t.Errorf("Category %s has zero UpdatedAt", category.Name)
		}
	}
}

// Benchmark para medir performance
func BenchmarkListCategoriesUseCase_Execute(b *testing.B) {
	mockRepo := &mockCategoryRepository{
		savedCategories: createMockCategories(),
	}
	useCase := NewListCategoriesUseCase(mockRepo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = useCase.Execute()
	}
}

// Teste de performance com diferentes tamanhos de lista
func BenchmarkListCategoriesUseCase_LargeList(b *testing.B) {
	// Criar uma lista grande de categorias para teste de performance
	largeList := make([]*entities.Category, 1000)
	for i := 0; i < 1000; i++ {
		category, _ := entities.NewCategory(fmt.Sprintf("Category %d", i))
		largeList[i] = category
	}

	mockRepo := &mockCategoryRepository{
		savedCategories: largeList,
	}
	useCase := NewListCategoriesUseCase(mockRepo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = useCase.Execute()
	}
}

// Teste de stress - múltiplas execuções simultâneas
func TestListCategoriesUseCase_ConcurrentAccess(t *testing.T) {
	// Arrange
	mockRepo := &mockCategoryRepository{
		savedCategories: createMockCategories(),
	}
	useCase := NewListCategoriesUseCase(mockRepo)

	// Act - Execute múltiplas goroutines simultaneamente
	done := make(chan bool, 10)
	errors := make(chan error, 10)

	for i := 0; i < 10; i++ {
		go func() {
			categories, err := useCase.Execute()
			if err != nil {
				errors <- err
				return
			}
			if len(categories) != 3 {
				errors <- fmt.Errorf("expected 3 categories, got %d", len(categories))
				return
			}
			done <- true
		}()
	}

	// Assert
	for i := 0; i < 10; i++ {
		select {
		case err := <-errors:
			t.Fatalf("Concurrent execution failed: %v", err)
		case <-done:
			// Success
		case <-time.After(time.Second):
			t.Fatal("Test timed out")
		}
	}
}

// Helper function para criar categorias mock
func createMockCategories() []*entities.Category {
	now := time.Now()
	return []*entities.Category{
		{
			ID:        1,
			Name:      "Category 1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        2,
			Name:      "Category 2",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        3,
			Name:      "Category 3",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}