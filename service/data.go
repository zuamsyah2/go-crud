package service

import (
	"go-crud/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetData(db *gorm.DB, ctx *gin.Context) ([]model.Data, error) {
	var data []model.Data
	err := db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CreateData(db *gorm.DB, ctx *gin.Context, bodyRequest model.Data) error {
	err := db.Debug().Create(&bodyRequest).Error
	if err != nil {
		return nil
	}

	return nil
}

func UpdateData(db *gorm.DB, ctx *gin.Context, bodyRequest model.Data) error {
	data := model.Data{}
	err := db.Model(model.Data{}).Where("id = ?", bodyRequest.ID).First(&data).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return err
		}

		return err
	}

	err = db.Updates(bodyRequest).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteData(db *gorm.DB, ctx *gin.Context, id string) {
	// Misalkan Anda ingin menghapus user dengan ID 1
	var dataToDelete model.Data
	// Opsional: Cek apakah user ada sebelum menghapus
	if err := db.First(&dataToDelete, id).Error; err != nil {
		// Handle error, misalnya user tidak ditemukan
		if err == gorm.ErrRecordNotFound {
			println("Data dengan ID " + id + " tidak ditemukan.")
		} else {
			println("Error mencari user:", err.Error())
		}
	} else {
		// Jika user ditemukan, lanjutkan penghapusan
		if result := db.Delete(&dataToDelete); result.Error != nil {
			println("Gagal menghapus user:", result.Error.Error())
		} else {
			println("User dengan ID 1 berhasil dihapus secara janglanpermanen. Jumlah baris terpengaruh:", result.RowsAffected)
		}
	}
}
