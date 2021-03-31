package shared

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ResponseDetail struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseList data structure
type ResponseList struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Meta    *Meta  `json:"meta,omitempty"`
}

// Meta data structure
type Meta struct {
	TotalRecords int    `json:"totalRecords,omitempty"`
	TotalPages   int    `json:"totalPages,omitempty"`
	OrderBy      string `json:"orderBy,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Page         int    `json:"page,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

func ResponseDetailOutput(success bool, code int, message string, data interface{}) ResponseDetail {
	return ResponseDetail{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func GetFinalErrorSql(err error, dataName string) error {
	if err == sql.ErrNoRows {
		err = fmt.Errorf(ErrorDataNotFound, dataName)
	}

	return err
}

func GetError(t *testing.T, iserror bool, err error) {
	if iserror {
		t.Logf(err.Error())
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
}
