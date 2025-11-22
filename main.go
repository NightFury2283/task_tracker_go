package main

// main.go

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tasktracker/task"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { //считываем команды пользователя

		stroke := scanner.Text()

		if len(strings.Fields(stroke)) < 2 { //если кроме task-cli ничего нет, то продолжаем
			continue
		}

		if strings.Fields(stroke)[0] != "task-cli" { //если первая команда не task-cli, то продолжаем
			continue
		}

		stroke_without_cli := strings.Fields(stroke)[1:] //обрезаем

		Parse_Command(stroke_without_cli)
	}
}

func Parse_Command(stroke []string) {
	switch stroke[0] {
	case "add":
		if len(stroke) < 2 {
			fmt.Println("no description for new task")
			return
		}

		description := stroke[1]

		id, err := task.AddTask(description)

		if err != nil {
			fmt.Println("got error: ", err)
			return
		}

		fmt.Printf("Task added successfully (ID: %d)\n", id)
		return
	case "update":
		if len(stroke) < 3 {
			fmt.Println("no id or description in command update")
			return
		}
		id, err := strconv.Atoi(stroke[1])
		if err != nil {
			fmt.Println("cannot convert string to int ID", err)
			return
		}
		if err := task.UpdateTask(id, stroke[2]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("succesfully update task %s.\n", stroke[2])
	case "delete":
		if len(stroke) < 2 {
			fmt.Println("no id to delete task")
			return
		}
		id, err := strconv.Atoi(stroke[1])
		if err != nil {
			fmt.Println("cannot convert string to int ID", err)
		}
		if err := task.DeleteTask(id); err != nil {
			fmt.Println(err)
			return
		}
	case "list":
		if len(stroke) > 1 {
			Parse_Command_List(stroke[1:])
			return
		}
		if err := task.ListTasks(); err != nil {
			fmt.Println("failde to show you all tasks", err)
			return
		}
		return
	case "mark-in-progress":
		if len(stroke) < 2 {
			fmt.Println("no id in command")
			return
		}
		id, err := strconv.Atoi(stroke[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		task.MarkTaskInProgress(id)
	case "mark-done":
		if len(stroke) < 2 {
			fmt.Println("no id in command")
			return
		}
		id, err := strconv.Atoi(stroke[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		task.MarkTaskDone(id)
	default:
		fmt.Println("don't understand your command, try again")
		return
	}
}

func Parse_Command_List(stroke []string) { //сюда попадаем если команда была List и что-то дальше
	if err := task.ListWithParametr(stroke[0]); err != nil {
		fmt.Println(err)
		return
	}
	return
}
