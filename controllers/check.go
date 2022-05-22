package controllers

import (
	"houseduino-be/models"
)

func (h *BaseHandlerSqlx) GetTestSqlx() {

	test := models.GetTestSqlx(h.db.DB)

	PrintLog("Test", test)
}
