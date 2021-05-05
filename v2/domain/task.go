package domain

type TaskState int

const (
	Backlog TaskState = iota
	Cancel
	Done
)

// Task contain task attribute and task state.
type Task struct {
	Name  string
	state TaskState
}

// NewTask starts from the Backlog state.
func NewTask(name string) *Task {
	return &Task{
		Name:  name,
		state: Backlog,
	}
}

func (t *Task) GetState() TaskState {
	return t.state
}

func (t *Task) NextState(next TaskState) {
	t.state = next
}
