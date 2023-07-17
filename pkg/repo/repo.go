package repo

type RepoTodo interface {
	Add(todo string)
	Remove(todo string)
	Get() []string
}

type MemoryRepo struct {
	todos []string
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{}
}

// Add or Create
func (r *MemoryRepo) Add(todo string) {
	r.todos = append(r.todos, todo)
}

// Remove
func (r *MemoryRepo) Remove(todo string) {
	for index, taskValue := range r.todos {
		if taskValue == todo {
			r.todos = append(r.todos[1:], r.todos[index+1:]...)
			break
		}
	}
}

func (r *MemoryRepo) Get() []string {
	return r.todos
}
