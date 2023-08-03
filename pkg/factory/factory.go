package factory

import (
	"fmt"
	. "galileo/pkg/constant"
)

// NewTaskProgressKey
// /* 生成规则：taskProgress:taskName:executeId */
func NewTaskProgressKey(taskName string, executeId int64) string {
	return fmt.Sprintf("%s:%s:%d", TaskProgressKey, taskName, executeId)
}
