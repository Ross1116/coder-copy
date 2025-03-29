package test

// nolint :all
import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// TaskStatus represents the current state of a task
type TaskStatus string

// Constants for task statuses
const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusFailed    TaskStatus = "failed"
)

// Task represents a unit of work to be performed
type Task struct {
	ID          string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
	Result      string
	Error       error
}

// TaskManager handles the creation, tracking, and execution of tasks
type TaskManager struct {
	tasks  map[string]*Task
	mutex  sync.RWMutex
	logger *log.Logger
}

// NewTaskManager creates a new task manager instance
func NewTaskManager(logger *log.Logger) *TaskManager {
	return &TaskManager{
		tasks:  make(map[string]*Task),
		mutex:  sync.RWMutex{},
		logger: logger,
	}
}

// CreateTask adds a new task to the manager
func (tm *TaskManager) CreateTask(id, description string) (*Task, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// Validate inputs
	if id == "" {
		return nil, errors.New("task ID cannot be empty")
	}

	// Check if task with this ID already exists
	if _, exists := tm.tasks[id]; exists {
		return nil, fmt.Errorf("task with ID %s already exists", id)
	}

	// Create new task
	now := time.Now()
	task := &Task{
		ID:          id,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Add to tasks map
	tm.tasks[id] = task
	tm.logger.Printf("Created task %s: %s", id, description)

	return task, nil
}

// GetTask retrieves a task by ID
func (tm *TaskManager) GetTask(id string) (*Task, error) {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	task, exists := tm.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task with ID %s not found", id)
	}

	return task, nil
}

// ListTasks returns all tasks
func (tm *TaskManager) ListTasks() []*Task {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	tasks := make([]*Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

// ExecuteTask runs a task asynchronously
func (tm *TaskManager) ExecuteTask(id string, workFn func() (string, error)) error {
	// Get the task
	task, err := tm.GetTask(id)
	if err != nil {
		return err
	}

	// Update task status to running
	tm.mutex.Lock()
	if task.Status != StatusPending {
		tm.mutex.Unlock()
		return fmt.Errorf("cannot execute task %s with status %s", id, task.Status)
	}
	task.Status = StatusRunning
	task.UpdatedAt = time.Now()
	tm.mutex.Unlock()

	tm.logger.Printf("Starting execution of task %s", id)

	// Execute task asynchronously
	go func() {
		// Execute the work function
		result, err := workFn()

		// Update task with result
		tm.mutex.Lock()
		defer tm.mutex.Unlock()

		now := time.Now()
		task.UpdatedAt = now

		if err != nil {
			task.Status = StatusFailed
			task.Error = err
			tm.logger.Printf("Task %s failed: %v", id, err)
		} else {
			task.Status = StatusCompleted
			task.Result = result
			task.CompletedAt = &now
			tm.logger.Printf("Task %s completed with result: %s", id, result)
		}
	}()

	return nil
}

// DeleteTask removes a task from the manager
func (tm *TaskManager) DeleteTask(id string) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	if _, exists := tm.tasks[id]; !exists {
		return fmt.Errorf("task with ID %s not found", id)
	}

	delete(tm.tasks, id)
	tm.logger.Printf("Deleted task %s", id)

	return nil
}

// simulateWork is a helper function that simulates a long-running task
func simulateWork(duration time.Duration, shouldFail bool) (string, error) {
	/*
	   This function simulates work by sleeping for the specified duration.
	   It can be configured to fail or succeed based on the shouldFail parameter.
	*/
	time.Sleep(duration)

	if shouldFail {
		return "", errors.New("task simulation failed")
	}

	return "task completed successfully", nil
}

func main() {
	// Initialize logger
	logger := log.New(log.Writer(), "[TaskManager] ", log.LstdFlags)

	// Create task manager
	manager := NewTaskManager(logger)

	// Create some tasks
	task1, err := manager.CreateTask("task-1", "First example task")
	if err != nil {
		logger.Fatalf("Failed to create task: %v", err)
	}

	task2, err := manager.CreateTask("task-2", "Second example task")
	if err != nil {
		logger.Fatalf("Failed to create task: %v", err)
	}

	// Execute tasks
	err = manager.ExecuteTask(task1.ID, func() (string, error) {
		return simulateWork(2*time.Second, false)
	})
	if err != nil {
		logger.Printf("Failed to execute task %s: %v", task1.ID, err)
	}

	err = manager.ExecuteTask(task2.ID, func() (string, error) {
		return simulateWork(1*time.Second, true)
	})
	if err != nil {
		logger.Printf("Failed to execute task %s: %v", task2.ID, err)
	}

	// Wait for tasks to complete
	time.Sleep(3 * time.Second)

	// Print task results
	tasks := manager.ListTasks()
	fmt.Println("\nTask Results:")
	for _, task := range tasks {
		fmt.Printf("- Task %s: Status=%s", task.ID, task.Status)
		if task.Status == StatusCompleted {
			fmt.Printf(", Result=%s", task.Result)
		} else if task.Status == StatusFailed && task.Error != nil {
			fmt.Printf(", Error=%v", task.Error)
		}
		fmt.Println()
	}

	// Cleanup
	for _, task := range tasks {
		manager.DeleteTask(task.ID)
	}
}
