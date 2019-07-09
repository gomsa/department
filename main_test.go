package main

import (
	"context"
	"fmt"
	"testing"

	// 执行数据迁移
	_ "github.com/gomsa/department/database/migrations"

	"github.com/gomsa/department/hander"
	departmentPB "github.com/gomsa/department/proto/department"
	db "github.com/gomsa/department/providers/database"
	"github.com/gomsa/department/service"
)

func TestDepartmentCreate(t *testing.T) {
	repo := &service.DepartmentRepository{db.DB}
	h := hander.Department{repo}
	req := &departmentPB.Department{
		Id:   103,
		Name: "三楼",
	}
	res := &departmentPB.Response{}
	err := h.Create(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestDepartmentAll(t *testing.T) {
	repo := &service.DepartmentRepository{db.DB}
	h := hander.Department{repo}
	req := &departmentPB.Department{
		Parent: 1,
	}
	res := &departmentPB.Response{}
	err := h.All(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}


func TestDepartmentDelete(t *testing.T) {
	repo := &service.DepartmentRepository{db.DB}
	h := hander.Department{repo}
	req := &departmentPB.Department{
		Id: 103,
	}
	res := &departmentPB.Response{}
	err := h.Delete(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
