/*
Copyright <holder> All Rights Reserved.

SPDX-License-Identifier: Apache-2.0

History:
   Date     Who ID    Description
   -------- --- ---   -----------
   01/13/19 nanjj  Initial code

*/

package model

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UUID      string     `gorm:"type:varchar(64)"`
	Owner     int64
}

func (m *Model) BeforeCreate() (err error) {
	if m.UUID == "" {
		m.UUID = uuid.New().String()
	}
	return
}
