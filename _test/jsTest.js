/**
 * TaskManager - A class for managing tasks and their lifecycle
 * @module TaskManager
 */

// Import required modules
import { v4 as uuidv4 } from 'uuid';
import { EventEmitter } from 'events';
import { Storage } from './storage.js';

// Constants
const PRIORITY = {
  LOW: 'low',
  MEDIUM: 'medium',
  HIGH: 'high',
  URGENT: 'urgent'
};

const STATUS = {
  PENDING: 'pending',
  IN_PROGRESS: 'in-progress',
  COMPLETED: 'completed',
  ARCHIVED: 'archived'
};

/**
 * Class representing a Task Manager
 * @extends EventEmitter
 */
class TaskManager extends EventEmitter {
  /**
   * Create a task manager
   * @param {Object} options - Configuration options
   * @param {Storage} options.storage - Storage implementation
   */
  constructor(options = {}) {
    super();
    this.storage = options.storage || new Storage();
    this.tasks = [];
    
    // Initialize the task list
    this._init();
  }
  
  /**
   * Initialize the task manager
   * @private
   */
  async _init() {
    try {
      // Load tasks from storage
      this.tasks = await this.storage.getAll() || [];
      this.emit('ready', this.tasks.length);
    } catch (error) {
      this.emit('error', error);
      console.error('Failed to initialize task manager:', error);
    }
  }
  
  /**
   * Create a new task
   * @param {Object} taskData - The task data
   * @returns {Object} The created task
   */
  createTask(taskData) {
    if (!taskData.title) {
      throw new Error('Task title is required');
    }
    
    // Create a new task with defaults
    const task = {
      id: uuidv4(),
      title: taskData.title,
      description: taskData.description || '',
      status: STATUS.PENDING,
      priority: taskData.priority || PRIORITY.MEDIUM,
      tags: taskData.tags || [],
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
      dueDate: taskData.dueDate || null
    };
    
    // Add to the task list
    this.tasks.push(task);
    
    // Persist to storage
    this.storage.save(task)
      .then(() => {
        this.emit('task:created', task);
      })
      .catch(error => {
        // If storage fails, we should probably remove the task from the list
        this.tasks = this.tasks.filter(t => t.id !== task.id);
        this.emit('error', error);
      });
    
    return task;
  }
  
  /**
   * Get a task by ID
   * @param {string} id - The task ID
   * @returns {Object|null} The task or null if not found
   */
  getTask(id) {
    return this.tasks.find(task => task.id === id) || null;
  }
  
  /**
   * Update a task
   * @param {string} id - The task ID
   * @param {Object} updates - The updates to apply
   * @returns {Object|null} The updated task or null if not found
   */
  updateTask(id, updates) {
    const taskIndex = this.tasks.findIndex(task => task.id === id);
    
    if (taskIndex === -1) {
      return null;
    }
    
    // Get the current task
    const task = this.tasks[taskIndex];
    
    // Create the updated task
    const updatedTask = {
      ...task,
      ...updates,
      updatedAt: new Date().toISOString()
    };
    
    // Don't allow changing the ID or creation date
    updatedTask.id = task.id;
    updatedTask.createdAt = task.createdAt;
    
    // Update the task in the list
    this.tasks[taskIndex] = updatedTask;
    
    // Persist to storage
    this.storage.update(updatedTask)
      .then(() => {
        this.emit('task:updated', updatedTask);
      })
      .catch(error => {
        // Revert to the original task if storage fails
        this.tasks[taskIndex] = task;
        this.emit('error', error);
      });
    
    return updatedTask;
  }
  
  /**
   * Delete a task
   * @param {string} id - The task ID
   * @returns {boolean} True if the task was deleted, false otherwise
   */
  deleteTask(id) {
    const taskIndex = this.tasks.findIndex(task => task.id === id);
    
    if (taskIndex === -1) {
      return false;
    }
    
    // Remove the task from the list
    const [deletedTask] = this.tasks.splice(taskIndex, 1);
    
    // Remove from storage
    this.storage.delete(id)
      .then(() => {
        this.emit('task:deleted', deletedTask);
      })
      .catch(error => {
        // Restore the task if storage fails
        this.tasks.splice(taskIndex, 0, deletedTask);
        this.emit('error', error);
      });
    
    return true;
  }
  
  /**
   * Get tasks filtered by criteria
   * @param {Object} filters - The filter criteria
   * @returns {Array} The filtered tasks
   */
  getTasks(filters = {}) {
    let filteredTasks = [...this.tasks];
    
    // Apply status filter
    if (filters.status) {
      filteredTasks = filteredTasks.filter(task => task.status === filters.status);
    }
    
    // Apply priority filter
    if (filters.priority) {
      filteredTasks = filteredTasks.filter(task => task.priority === filters.priority);
    }
    
    // Apply tag filter
    if (filters.tag) {
      filteredTasks = filteredTasks.filter(task => task.tags.includes(filters.tag));
    }
    
    // Apply search filter
    if (filters.search) {
      const searchLower = filters.search.toLowerCase();
      filteredTasks = filteredTasks.filter(task => 
        task.title.toLowerCase().includes(searchLower) || 
        task.description.toLowerCase().includes(searchLower)
      );
    }
    
    /* 
      Sorting logic:
      - By default, sort by creation date (newest first)
      - Can be overridden by the sort parameter
    */
    const sortBy = filters.sortBy || 'createdAt';
    const sortDir = filters.sortDir === 'asc' ? 1 : -1;
    
    filteredTasks.sort((a, b) => {
      if (a[sortBy] < b[sortBy]) return -1 * sortDir;
      if (a[sortBy] > b[sortBy]) return 1 * sortDir;
      return 0;
    });
    
    return filteredTasks;
  }
}

// Export the TaskManager class and constants
export { TaskManager, PRIORITY, STATUS };

// Example usage:
// const taskManager = new TaskManager();
// taskManager.on('ready', count => console.log(`Loaded ${count} tasks`));
// taskManager.createTask({ title: 'Learn JavaScript', priority: PRIORITY.HIGH });