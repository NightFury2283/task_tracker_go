package task

import (
	"errors"
	"fmt"
	"tasktracker/storage"
	"tasktracker/types"
	"time"
)

var file_path = "storage/storage.json"

func AddTask(description string) (int, error) {
	tasks, id, err := storage.FileLoadTasks(file_path)

	if err != nil {
		return 0, err
	}

	if _, ok := tasks[id]; ok {
		return 0, errors.New("this task.id already exists. error on server side")
	}

	task := types.Task{
		ID:          id,
		Description: description,
		Status:      types.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := storage.FileAddTask(file_path, task); err != nil {
		return 0, err
	}

	//TODO: update next id

	return id, nil
}

func UpdateTask(id int, new_description string) error {
	all_tasks, storage_next_id, err := storage.FileLoadTasks(file_path)
	if err != nil {
		return err
	}

	if _, ok := all_tasks[id]; !ok {
		return errors.New("no such task in storage json file")
	}
	//такая задача есть, можно менять описание
	task_to_update := all_tasks[id]

	//новое описание и обновление
	task_to_update.Description = new_description
	task_to_update.UpdatedAt = time.Now()

	all_tasks[id] = task_to_update

	if err = storage.FileSaveTasks(file_path, all_tasks, storage_next_id); err != nil {
		return err
	}

	return nil
}

func DeleteTask(id int) error {
	if err := storage.FileDeleteTask(file_path, id); err != nil {
		return err
	}
	return nil
}

func ListTasks() error {
	all_tasks, _, err := storage.FileLoadTasks(file_path)

	if err != nil {
		return err
	}
	for _, value := range all_tasks {
		fmt.Println("-----------------")
		fmt.Println(value.ID)
		fmt.Println(value.Description)
		fmt.Println(value.Status)
		fmt.Println(value.CreatedAt)
		fmt.Println(value.UpdatedAt)
		fmt.Println("-----------------")
	}
	return nil
}
