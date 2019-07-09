package service

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-log"

	pb "github.com/gomsa/department/proto/department"

	"github.com/jinzhu/gorm"
)

//URepository 仓库接口
type URepository interface {
	Create(department *pb.Department) (*pb.Department, error)
	Get(department *pb.Department) (*pb.Department, error)
	List(req *pb.ListQuery) ([]*pb.Department, error)
	Total(req *pb.ListQuery) (int64, error)
	Update(department *pb.Department) (bool, error)
	Delete(department *pb.Department) (bool, error)
}

// DepartmentRepository 部门仓库
type DepartmentRepository struct {
	DB *gorm.DB
}

// List 获取所有部门信息
func (repo *DepartmentRepository) List(req *pb.ListQuery) (departments []*pb.Department, err error) {
	db := repo.DB
	// 分页
	var limit, offset int64
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = -1
	}

	// 排序
	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "created_at desc"
	}
	// 查询条件
	if req.Id != "" {
		db = db.Where("id like ?", "%"+req.Id+"%")
	}
	if req.Parent != "" {
		db = db.Where("parent like ?", "%"+req.Parent+"%")
	}
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&departments).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return departments, nil
}

// Total 获取所有部门查询总量
func (repo *DepartmentRepository) Total(req *pb.ListQuery) (total int64, err error) {
	departments := []pb.Department{}
	db := repo.DB
	// 查询条件
	if req.Id != "" {
		db = db.Where("id like ?", "%"+req.Id+"%")
	}
	if req.Parent != "" {
		db = db.Where("parent like ?", "%"+req.Parent+"%")
	}
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if err := db.Find(&departments).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取部门信息
func (repo *DepartmentRepository) Get(department *pb.Department) (*pb.Department, error) {
	if department.Id != "" {
		if err := repo.DB.Model(&department).Where("id = ?", department.Id).Find(&department).Error; err != nil {
			return nil, err
		}
	}
	if department.Parent != "" {
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

// Gets 获取部门信息
func (repo *DepartmentRepository) Gets(department *pb.Department) (departments []*pb.Department, err error) {
	if department.Id != "" {
		if err = repo.DB.Model(&department).Where("id = ?", department.Id).Find(&departments).Error; err != nil {
			return nil, err
		}
	}
	if department.Parent != "" {
		if err = repo.DB.Model(&department).Where("parent = ?", department.Parent).Find(&departments).Error; err != nil {
			return nil, err
		}
	}
	if department.Name != "" {
		if err = repo.DB.Model(&department).Where("name = ?", department.Name).Find(&departments).Error; err != nil {
			return nil, err
		}
	}
	return departments, err
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
	if department.Id == "" {
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
