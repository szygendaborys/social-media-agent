package todo

import (
	"testing"
)

func TestGetTodos(t *testing.T) {
	// Reset the todos_db for testing
	todos_db = []Task{}

	// Test empty todos
	todos, err := GetTodos()
	if err != nil {
		t.Errorf("GetTodos() returned error: %v", err)
	}
	if len(todos) != 0 {
		t.Errorf("GetTodos() returned %d todos, want 0", len(todos))
	}

	// Test with some todos
	testTodo := Task{
		ID:        1,
		Title:     "Test Todo",
		Completed: false,
	}
	todos_db = append(todos_db, testTodo)

	todos, err = GetTodos()
	if err != nil {
		t.Errorf("GetTodos() returned error: %v", err)
	}
	if len(todos) != 1 {
		t.Errorf("GetTodos() returned %d todos, want 1", len(todos))
	}
	if todos[0].ID != testTodo.ID {
		t.Errorf("GetTodos() returned todo with ID %d, want %d", todos[0].ID, testTodo.ID)
	}
}

func TestAddTodo(t *testing.T) {
	// Reset the todos_db for testing
	todos_db = []Task{}

	// Test adding a valid todo
	testTodo := Task{
		ID:        1,
		Title:     "Test Todo",
		Completed: false,
	}

	err := AddTodo(testTodo)
	if err != nil {
		t.Errorf("AddTodo() returned error: %v", err)
	}

	// Verify the todo was added
	if len(todos_db) != 1 {
		t.Errorf("AddTodo() did not add todo to database, len(todos_db) = %d", len(todos_db))
	}
	if todos_db[0].ID != testTodo.ID {
		t.Errorf("AddTodo() added todo with ID %d, want %d", todos_db[0].ID, testTodo.ID)
	}

	// Test adding another todo
	anotherTodo := Task{
		ID:        2,
		Title:     "Another Todo",
		Completed: true,
	}

	err = AddTodo(anotherTodo)
	if err != nil {
		t.Errorf("AddTodo() returned error: %v", err)
	}

	// Verify both todos are in the database
	if len(todos_db) != 2 {
		t.Errorf("AddTodo() did not add second todo to database, len(todos_db) = %d", len(todos_db))
	}
}
