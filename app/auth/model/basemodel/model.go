package basemodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                     // 主键
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at" form:"updated_at"` // 更新时间
	CreatedBy uint64    `gorm:"column:created_by;default:0;not null;" json:"created_by" form:"created_by"`     // 创建人
	UpdatedBy uint64    `gorm:"column:updated_by;default:0;not null;" json:"updated_by" form:"updated_by"`     // 更新人
}

// 分页条件
type PageWhereOrder struct {
	Order string
	Where string
	Value []interface{}
}

// Create
func Create(cdb *gorm.DB, value interface{}) error {
	return cdb.Create(value).Error
}

// Save
func Save(cdb *gorm.DB, value interface{}) error {
	return cdb.Save(value).Error
}

// Updates
func Updates(cdb *gorm.DB, where interface{}, value interface{}) error {
	return cdb.Model(where).Updates(value).Error
}

// Delete
func DeleteByModel(cdb *gorm.DB, model interface{}) (count int64, err error) {
	db := cdb.Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByWhere(cdb *gorm.DB, model, where interface{}) (count int64, err error) {
	db := cdb.Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByID(cdb *gorm.DB, model interface{}, id uint64) (count int64, err error) {
	db := cdb.Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByIDS(cdb *gorm.DB, model interface{}, ids []uint64) (count int64, err error) {
	db := cdb.Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// First
func FirstByID(cdb *gorm.DB, out interface{}, id int) (notFound bool, err error) {
	err = cdb.First(out, id).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// First
func First(cdb *gorm.DB, where interface{}, out interface{}) (notFound bool, err error) {
	err = cdb.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find
func Find(cdb *gorm.DB, where interface{}, out interface{}, orders ...string) error {
	db := cdb.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan
func Scan(cdb *gorm.DB, model, where interface{}, out interface{}) (notFound bool, err error) {
	err = cdb.Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// ScanList
func ScanList(cdb *gorm.DB, model, where interface{}, out interface{}, orders ...string) error {
	db := cdb.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// GetPage
func GetPage(cdb *gorm.DB, model, where interface{}, out interface{}, pageIndex, pageSize uint64, totalCount *uint64, whereOrder ...PageWhereOrder) error {
	db := cdb.Model(model).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList
func PluckList(cdb *gorm.DB, model, where interface{}, out interface{}, fieldName string) error {
	return cdb.Model(model).Where(where).Pluck(fieldName, out).Error
}
