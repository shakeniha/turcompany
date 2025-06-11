package services
import(
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type DealService struct {
	Repo *repositories.DealRepository
}
func NewDealService(repo *repositories.DealRepository)*DealService {
	return &DealService{Repo : repo}
}
func (s *DealService) Create(deal *models.Deals) error {
	return s.Repo.Create(deal)
}
func (s *DealService) Update(deal *models.Deals)error{
	return s.Repo.Update(deal)
}
func (s *DealService) GetByID(id string)(*models.Deals, error){
	return s.Repo.GetByID(id)
}
func (s *DealService) Delete(id string) error {
	return s.Repo.Delete(id)
}

