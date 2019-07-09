package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/department/proto/department"
	"github.com/gomsa/department/service"
)

// Department 部门结构
type Department struct {
	Repo service.URepository
}

// List 获取所有部门
func (srv *Department) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	departments, err := srv.Repo.List(req)
	total, err := srv.Repo.Total(req)
	if err != nil {
		return err
	}
	res.Total = total
	res.Departments = departments
	return err
}

// Get 获取部门
func (srv *Department) Get(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	department, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Department = department
	return err
}

// Gets 根据条件查询 返回多条数据
func (srv *Department) Gets(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	departments, err := srv.Repo.Gets(req)
	if err != nil {
		return err
	}
	res.Departments = departments
	return err
}

// Create 创建部门
func (srv *Department) Create(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建部门失败")
	}
	res.Valid = true
	return err
}

// Update 更新部门
func (srv *Department) Update(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新部门失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除部门部门
func (srv *Department) Delete(ctx context.Context, req *pb.Department, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除部门失败")
	}
	res.Valid = valid
	return err
}
