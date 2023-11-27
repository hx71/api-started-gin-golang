package helpers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/models"
	"gorm.io/gorm"
)

var (
	dbx *gorm.DB = config.SetupConnection()
)

func CreateLogInfo(userID, ipAddress, serviceName, methodName, metadata string) {
	// defer wg.Done()

	var logs models.AuditLog
	logs.ID = uuid.NewString()
	logs.UserID = userID
	logs.IPAddress = ipAddress
	logs.ServiceName = serviceName
	logs.MethodName = methodName
	logs.Metadata = metadata
	logs.Level = "Info"
	res := dbx.Save(&logs)
	fmt.Println(res)
}

func CreateLogError(userID, ipAddress, serviceName, methodName, metadata string) {
	var logs models.AuditLog
	logs.ID = uuid.NewString()
	logs.UserID = userID
	logs.IPAddress = ipAddress
	logs.ServiceName = serviceName
	logs.MethodName = methodName
	logs.Metadata = metadata
	logs.Level = "Error"
	res := dbx.Save(&logs)
	fmt.Println(res)
}
