package choruspro

import (
	"context"
	"net/http"
	"time"
)

type ListeCategoriesSollicitation struct {
	CodeRetour int32                    `json:"codeRetour"`
	Libelle    string                   `json:"libelle"`
	Categories []CategorieSollicitation `json:"listeCategories"`
	Pagination *PaginationResponse      `json:"parametresRetour"`
}

type ListeCategoriesSollicitationOptions struct {
	Code       string             `json:"codeCategorie,omitempty"`
	EstActif   bool               `json:"estActif,omitempty"`
	Libelle    string             `json:"libelleCategorie,omitempty"`
	Pagination *PaginationOptions `json:"pagination,omitempty"`
}

func (s *TransversesService) RechercherCategoriesSollicitation(ctx context.Context, opts ListeCategoriesSollicitationOptions) (*ListeCategoriesSollicitation, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/rechercher/categorieSollicitation", opts)
	if err != nil {
		return nil, err
	}

	categories := new(ListeCategoriesSollicitation)

	err = s.client.doRequest(ctx, req, categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

type ListeSousCategoriesSollicitation struct {
	CodeRetour     int32  `json:"codeRetour"`
	Libelle        string `json:"libelle"`
	SousCategories []struct {
		Categories     []CategorieSollicitation     `json:"categorie"`
		SousCategories []SousCategorieSollicitation `json:"ssCategorie"`
	} `json:"listeSousCategories"`
	Pagination *PaginationResponse `json:"parametresRetour"`
}

type CategorieSollicitation struct {
	Code        string `json:"codeCategorie"`
	Libelle     string `json:"libelleCategorie"`
	IdTechnique int64  `json:"idTechniqueCategorie"`
}

type SousCategorieSollicitation struct {
	Code                     string     `json:"code"`
	DateCreation             *time.Time `json:"dateCreation"`
	DateDerniereModification *time.Time `json:"dateDerniereModification"`
	EstActif                 bool       `json:"estActif"`
	IdTechnique              int64      `json:"idTechniqueCategorie"`
	Libelle                  string     `json:"libelle"`
}

type ListeSousCategoriesSollicitationOptions struct {
	Code                 string             `json:"code,omitempty"`
	EstActif             bool               `json:"estActif,omitempty"`
	IdTechniqueCategorie int64              `json:"idTechniqueCategorie,omitempty"`
	Libelle              string             `json:"libelle,omitempty"`
	PaginationOptions    *PaginationOptions `json:"pagination,omitempty"`
}

func (s *TransversesService) RechercherSousCategoriesSollicitation(ctx context.Context, opts ListeSousCategoriesSollicitationOptions) (*ListeSousCategoriesSollicitation, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/rechercher/sousCategorieSollicitation", opts)
	if err != nil {
		return nil, err
	}

	categories := new(ListeSousCategoriesSollicitation)

	err = s.client.doRequest(ctx, req, categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
