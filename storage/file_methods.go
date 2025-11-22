package storage

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"tasktracker/types"
)

func FileAddTask(file_path string, task types.Task) error {
	//проверка на айди уже сделана, задачи с таким id нет
	all_tasks, id, err := FileLoadTasks(file_path)
	if err != nil {
		return err
	}

	if err != nil {
		return errors.New("failed to convert task to json")
	}

	file, err := os.OpenFile(file_path, os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return errors.New("failed to open json file")
	}
	defer file.Close()

	all_tasks[id] = task

	if err = FileSaveTasks(file_path, all_tasks, id); err != nil {
		return err
	}

	return nil
}

func FileLoadTasks(filepath string) (map[int]types.Task, int, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, 0, errors.New("cannot read file")
	}

	// Если файл пустой или содержит только пробелы
	if len(strings.TrimSpace(string(data))) == 0 {
		return make(map[int]types.Task), 1, nil
	}

	var storageData types.StorageData

	if err = json.Unmarshal(data, &storageData); err != nil {
		return nil, 0, errors.New("cannot transform json file to storageData struct")
	}
	return storageData.Tasks, storageData.NextID, nil
}

func FileSaveTasks(file_path string, tasks map[int]types.Task, id int) error {
	var storageData types.StorageData

	storageData.NextID = id + 1
	storageData.Tasks = tasks

	json_format, err := json.MarshalIndent(storageData, "", "	")
	if err != nil {
		return err
	}

	if err = os.WriteFile(file_path, json_format, 0644); err != nil {
		return err
	}
	return nil
}

func FileDeleteTask(file_path string, id int) error {
	all_tasks, storage_next_id, err := FileLoadTasks(file_path)

	if err != nil {
		return err
	}

	_, ok := all_tasks[id]
	if !ok {
		return errors.New("cannot find this task in json file")
	}

	delete(all_tasks, id)

	FileSaveTasks(file_path, all_tasks, storage_next_id)
	return nil
}
