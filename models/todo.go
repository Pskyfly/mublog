package models

import (
	"fmt"
	"myblog/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//Todo在数据库中的增删改查
func AddTodo(todo *Todo) (err error) {
	//大坑，不要new一个结点
	err = dao.DB.Create(&todo).Error
	return
}
func FindTodoList(todolist *[]Todo) (err error) {
	err = dao.DB.Find(todolist).Error
	return
}

func GetTodoByID(id string) (Todo, error) {
	var err error
	var todo Todo
	err = dao.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, err
}

func SaveTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteTodo(todo *Todo) (err error) {
	err = dao.DB.Delete(todo).Error
	fmt.Println("------------------")
	return
}
