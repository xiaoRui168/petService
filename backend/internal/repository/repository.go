package repository

import (
	"petService/backend/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// --- Styling ---

func (r *Repository) GetAllStylings() ([]model.Styling, error) {
	var items []model.Styling
	err := r.db.Order("id asc").Find(&items).Error
	return items, err
}

func (r *Repository) GetStylingByID(id uint) (*model.Styling, error) {
	var item model.Styling
	err := r.db.First(&item, id).Error
	return &item, err
}

// --- Package ---

func (r *Repository) GetAllPackages() ([]model.Package, error) {
	var items []model.Package
	err := r.db.Order("id asc").Find(&items).Error
	return items, err
}

func (r *Repository) GetPackageByID(id uint) (*model.Package, error) {
	var item model.Package
	err := r.db.First(&item, id).Error
	return &item, err
}

// --- Photo ---

func (r *Repository) GetAllPhotos() ([]model.Photo, error) {
	var items []model.Photo
	err := r.db.Order("created_at desc").Find(&items).Error
	return items, err
}

// --- Review ---

func (r *Repository) GetAllReviews() ([]model.Review, error) {
	var items []model.Review
	err := r.db.Order("created_at desc").Find(&items).Error
	return items, err
}

// --- Appointment ---

func (r *Repository) CreateAppointment(appt *model.Appointment) error {
	return r.db.Create(appt).Error
}

func (r *Repository) GetAllAppointments() ([]model.Appointment, error) {
	var items []model.Appointment
	err := r.db.Order("created_at desc").Find(&items).Error
	return items, err
}
