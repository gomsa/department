package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	// 执行数据迁移
	_ "github.com/gomsa/department/database/migrations"

	"github.com/gomsa/department/hander"
	departmentPB "github.com/gomsa/department/proto/department"
	db "github.com/gomsa/department/providers/database"
	"github.com/gomsa/department/service"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 权限服务实现
	repo := &service.DepartmentRepository{db.DB}
	departmentPB.RegisterDepartmentsHandler(srv.Server(), &hander.Department{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
