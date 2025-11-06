package manager

import (
	"fmt"
	"time"
)

type Id int

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
	Any        Status = "any"
)

func (st Status) IsValid() bool {
	return st == Todo || st == InProgress || st == Done || st == Any
}

type Task struct {
	Id          Id
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskManager struct {
	Tasks  map[Id]*Task
	NextId Id
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:  make(map[Id]*Task),
		NextId: 1,
	}
}

func (tm *TaskManager) Add(description string) Id {
	tm.Tasks[tm.NextId] = &Task{
		Id:          tm.NextId,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tm.NextId++
	return tm.NextId - 1
}

func (tm *TaskManager) Update(id Id, newDescription string) bool {
	task, ok := tm.Tasks[id]
	if !ok {
		return false
	}
	task.Description = newDescription
	task.UpdatedAt = time.Now()
	return true
}

func (tm *TaskManager) Delete(id Id) bool {
	if _, ok := tm.Tasks[id]; !ok {
		return false
	}
	delete(tm.Tasks, id)
	tm.normaliseId()
	return true
}

func (tm *TaskManager) ChangeStatus(id Id, status Status) bool {
	task, ok := tm.Tasks[id]
	if !ok {
		return false
	}
	task.Status = status
	task.UpdatedAt = time.Now()
	return true
}

func (tm *TaskManager) List(marker Status) int {
	if len(tm.Tasks) == 0 {
		return 0
	}

	count := 0
	for _, task := range tm.Tasks {
		if marker != Any && task.Status != marker {
			continue
		}

		fmt.Printf("[%d] %s | %s | Created: %s | Updated: %s\n",
			task.Id, task.Description, task.Status,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.UpdatedAt.Format("2006-01-02 15:04"))

		count++
	}
	return count
}

func (tm *TaskManager) normaliseId() {
	newTasks := make(map[Id]*Task)
	var newId Id = 1

	for _, task := range tm.Tasks {
		task.Id = newId
		newTasks[newId] = task
		newId++
	}

	tm.Tasks = newTasks
	tm.NextId = newId
}
