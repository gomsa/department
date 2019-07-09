package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	departmentPB "github.com/gomsa/department/proto/department"
)

var (
	// Department 权限客户端
	Department departmentPB.DepartmentsClient
)

func init() {
	userSrvName := os.Getenv("DEPARTMENT_NAME")
	Department = departmentPB.NewDepartmentsClient(userSrvName, client.DefaultClient)
}
