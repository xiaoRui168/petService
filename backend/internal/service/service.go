package service

import (
	"petService/backend/internal/model"
	"petService/backend/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllStylings() ([]model.Styling, error) {
	return s.repo.GetAllStylings()
}

func (s *Service) GetStylingByID(id uint) (*model.Styling, error) {
	return s.repo.GetStylingByID(id)
}

func (s *Service) GetAllPackages() ([]model.Package, error) {
	return s.repo.GetAllPackages()
}

func (s *Service) GetPackageByID(id uint) (*model.Package, error) {
	return s.repo.GetPackageByID(id)
}

func (s *Service) GetAllPhotos() ([]model.Photo, error) {
	return s.repo.GetAllPhotos()
}

func (s *Service) GetAllReviews() ([]model.Review, error) {
	return s.repo.GetAllReviews()
}

func (s *Service) CreateAppointment(req *model.AppointmentRequest) error {
	appt := &model.Appointment{
		OwnerName:   req.OwnerName,
		Phone:       req.Phone,
		PetType:     req.PetType,
		PetName:     req.PetName,
		Date:        req.Date,
		Time:        req.Time,
		ServiceID:   req.ServiceID,
		ServiceType: req.ServiceType,
		StylingID:   req.StylingID,
		PackageID:   req.PackageID,
		Remark:      req.Remark,
	}
	return s.repo.CreateAppointment(appt)
}
