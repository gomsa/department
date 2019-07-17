package service

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-log"

	pb "github.com/gomsa/department/proto/department"

	"github.com/jinzhu/gorm"
)

//Repository 仓库接口
type Repository interface {
	All(req *pb.Department) ([]*pb.Department, error)
	Create(department *pb.Department) (*pb.Department, error)
	Get(department *pb.Department) (*pb.Department, error)
	List(req *pb.ListQuery) ([]*pb.Department, error)
	Update(department *pb.Department) (bool, error)
	Delete(department *pb.Department) (bool, error)
}

// DepartmentRepository 部门仓库
type DepartmentRepository struct {
	DB *gorm.DB
}

// All 获取所有部门信息
func (repo *DepartmentRepository) All(req *pb.Department) (departments []*pb.Department, err error) {
	if err := repo.DB.Find(&departments).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return departments, nil
}

// List 获取所有部门信息
func (repo *DepartmentRepository) List(req *pb.ListQuery) (departments []*pb.Department, err error) {
	db := repo.DB
	// 查询条件
	if req.Id != 0 {
		db = db.Where("id = ?", req.Id)
	}
	db = db.Where("parent = ?", req.Parent)
	if err := db.Find(&departments).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return departments, nil
}

// Get 获取部门信息
func (repo *DepartmentRepository) Get(department *pb.Department) (*pb.Department, error) {
	if department.Id != 0 {
		if err := repo.DB.Model(&department).Where("id = ?", department.Id).Find(&department).Error; err != nil {
			return nil, err
		}
	}
	if department.Parent != 0 {
		if err := repo.DB.Model(&department).Where("parent = ?", department.Parent).Find(&department).Error; err != nil {
			return nil, err
		}
	}
	if department.Name != "" {
		if err := repo.DB.Model(&department).Where("name = ?", department.Name).Find(&department).Error; err != nil {
			return nil, err
		}
	}
	return department, nil
}

// Create 创建部门
func (repo *DepartmentRepository) Create(department *pb.Department) (*pb.Department, error) {
	err := repo.DB.Create(department).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return department, fmt.Errorf("注册部门失败")
	}
	return department, nil
}

// Update 更新部门
func (repo *DepartmentRepository) Update(department *pb.Department) (bool, error) {
	if department.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Department{
		Id: department.Id,
	}
	err := repo.DB.Model(id).Updates(department).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除部门
func (repo *DepartmentRepository) Delete(department *pb.Department) (bool, error) {
	id := &pb.Department{
		Id: department.Id,
	}
	err := repo.DB.Delete(id).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
