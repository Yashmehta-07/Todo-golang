package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_"strconv"
	"log"
	
)

type Task struct{
	Id int `json:"Id"`
	Desc string `json:"Desc"`
}

var tasks []Task
var taskId int


func Add(w http.ResponseWriter, r *http.Request){

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask) 

	if err != nil || newTask.Desc == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return	
	}

	taskId++

	newTask.Id= taskId
	tasks = append(tasks,newTask)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Task added successfully!",
		"task":    newTask,
	}
	json.NewEncoder(w).Encode(response)

}

func List(w http.ResponseWriter, r *http.Request){

	if len(tasks) == 0 {
		w.Header().Set("Content-Type", "application/json")

		message := map[string]interface{}{
			"message": "No Task Found",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func Update( w http.ResponseWriter, r *http.Request){

	if len(tasks) == 0 {
		w.Header().Set("Content-Type", "application/json")

		message := map[string]interface{}{
			"message": "Todo is empty",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	//extracting id  from url
	// idStr := r.URL.Query().Get("id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil || id <= 0 {
	// 	http.Error(w, "Invalid task ID", http.StatusBadRequest)
	// 	return
	// }


	//extracting id from body
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask) 
	id := newTask.Id 
	if err != nil || id <= 0 || newTask.Desc == "" {
		http.Error(w, "Invalid task ID or description", http.StatusBadRequest)
		return	
	}

	index:=-1

	for k,v := range tasks{
		if id==v.Id{
			index=k
			break
		}
	}

	if index == -1{
		http.Error(w, "Task not found,invaild id", http.StatusNotFound)
		return
	}

	tasks[index].Desc =  newTask.Desc

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Task updated successfully!",
		"task":    tasks[index],
	})

}

func Delete( w http.ResponseWriter, r *http.Request){

	if len(tasks) == 0 {
		w.Header().Set("Content-Type", "application/json")

		message := map[string]interface{}{
			"message": "Todo is empty, no task to delete",
		}
		json.NewEncoder(w).Encode(message)
		return
	}

	//extracting id from body
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask) 
	id := newTask.Id 

	if err != nil || id <= 0  {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return	
	}
	
	//removing the task

	index:=-1

	for k,v := range tasks{
		if id==v.Id{
			index=k
			break
		}
	}

	// it not found then return
	if (index == -1){
		http.Error(w, "Task not found,invaild id", http.StatusNotFound)
		return
	}

	//slicing
	tmp1:= tasks[:index]
	tmp2:= tasks[index+1:]
	tasks= append(tmp1, tmp2...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})

}


func main() {

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			List(w, r)
		} else if r.Method == http.MethodPost {
			Add(w, r)
		} else if r.Method == http.MethodPut {
			Update(w, r)
		}else if r.Method == http.MethodDelete {
			Delete( w, r)
		}else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})

	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
	
}

