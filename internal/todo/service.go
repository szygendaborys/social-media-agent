package todo

var todos_db = []Task{}

func GetTodos() ([]Task, error) {
	return todos_db, nil
}

func AddTodo(todo Task) error {
	todos_db = append(todos_db, todo)
	return nil
}
