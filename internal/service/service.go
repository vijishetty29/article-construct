package service

import (
	"encoding/json"
	drawConstruct "github.com/user/article-construct-demo/internal/draw"
	"github.com/user/article-construct-demo/internal/dto"
	"github.com/user/article-construct-demo/internal/repository"
	"log"
	"net/http"
)

type ConstructHandler interface {
	GetConstructForIan(w http.ResponseWriter, req *http.Request)
	GetItemForIan(w http.ResponseWriter, req *http.Request)
	GetCaseForIan(w http.ResponseWriter, req *http.Request)
	GetVariantForIan(w http.ResponseWriter, req *http.Request)
}
type constructHandler struct {
	repo repository.ConstructRepository
}

func (c constructHandler) GetItemForIan(w http.ResponseWriter, req *http.Request) {
	ian := req.URL.Query().Get("ian")
	var item *dto.ItemDto

	//request on repository
	if result, err := c.repo.GetItemForIan(ian, req.Context()); err != nil {
		log.Println(err)
	} else {
		item = result
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (c constructHandler) GetCaseForIan(w http.ResponseWriter, req *http.Request) {
	ian := req.URL.Query().Get("ian")
	var sk *dto.CaseDto

	//request on repository
	if result, err := c.repo.GetCaseForIan(ian, req.Context()); err != nil {
		log.Println(err)
	} else {
		sk = result
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sk)
}

func (c constructHandler) GetVariantForIan(w http.ResponseWriter, req *http.Request) {
	ian := req.URL.Query().Get("ian")
	var ea *dto.VariantDto

	//request on repository
	if result, err := c.repo.GetVariantForIan(ian, req.Context()); err != nil {
		log.Println(err)
	} else {
		ea = result
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ea)
}

func (c constructHandler) GetConstructForIan(w http.ResponseWriter, req *http.Request) {

	ian := req.URL.Query().Get("ian")
	country := req.URL.Query().Get("country")

	//request on repository
	if result, err := c.repo.GetConstructForIan(ian, country, req.Context()); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Construct not found for IAN " + ian + " and country " + country)
	} else {
		drawConstruct.DrawItemConstruct(result, country, w)
	}
}

func NewConstructHandler(repo repository.ConstructRepository) ConstructHandler {
	return &constructHandler{repo: repo}
}
