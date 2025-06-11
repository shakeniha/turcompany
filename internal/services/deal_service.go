package services
import(
	"errors"
	"strings"
	"turcompany/internal/models"
	"turcompany/internal/repositories"
)

type DealService struct {
	Repo *repositories.DealRepository
}
func NewDealService(repo *repositories.DealRepository) *DealService {
	return &DealService{Repo: repo}
}

func (s *DealService) Create(deal *models.Deals) error {
	if strings.HasPrefix(deal.Amount.Value, "-") {
		return errors.New("deal amount cannot be negative")
	}
	if deal.Status == ""{
		deal.Status ="new"
	}
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

