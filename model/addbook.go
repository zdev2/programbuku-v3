package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (cr *DaftarBuku) CreateBuku(db *gorm.DB) error {
	err := db.Model(DaftarBuku{}).Create(&cr).Error
	if err != nil {
		return err
	}
	return nil
}
func (cr *DaftarBuku) GetByID(db *gorm.DB, id uint) (DaftarBuku, error) {
	respon := DaftarBuku{}
	err := db.Model(DaftarBuku{}).Where("id = ?", id).Take(&respon).Error
	if err != nil {
		return DaftarBuku{}, err // kalo error ksni
	}
	return respon, nil // kalo ga error ksni
}

func (cr *DaftarBuku) GetAll(db *gorm.DB) ([]DaftarBuku, error) {
	resp := []DaftarBuku{}
	err := db.Model(DaftarBuku{}).Find(&resp).Error
	if err != nil {
		return []DaftarBuku{}, err
	}

	return resp, nil
}

func (cr *DaftarBuku) UpdateOne(db *gorm.DB) error {
	err := db.Model(DaftarBuku{}).Where("id = ?", cr.ID).Updates(map[string]interface{}{
		"isbn":    cr.ISBN,
		"penulis": cr.Penulis,
		"tahun":   cr.Tahun,
		"judul":   cr.Judul,
		"gambar":  cr.Gambar,
		"stok":    cr.Stok,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *DaftarBuku) DeleteById(db *gorm.DB) error {
	err := db.Model(DaftarBuku{}).Where("id = ?", cr.ID).Delete(&cr).Error
	if err != nil {
		return err
	}
	return nil
}
func (book *DaftarBuku) UpsertBuku(db *gorm.DB) error {
	result := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface {
		}{
			"isbn": book.ISBN, "penulis": book.Penulis,
			"tahun": book.Tahun, "judul": book.Judul, "gambar": book.Gambar,
			"stok": book.Stok}),
	}).Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
