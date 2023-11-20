package choruspro

import (
	"context"
	"net/http"
	"time"
)

// ListeCategoriesSollicitationOptions est la structure de données utlisée
// pour appeler la méthode RechercherCategoriesSollicitation.
type ListeCategoriesSollicitationOptions struct {
	Code       string             `json:"codeCategorie,omitempty"`
	EstActif   bool               `json:"estActif,omitempty"`
	Libelle    string             `json:"libelleCategorie,omitempty"`
	Pagination *PaginationOptions `json:"pagination,omitempty"`
}

// ListeCategoriesSollicitationResponse est la structure de données représentant
// la réponse de la méthode RechercherCategoriesSollicitation.
type ListeCategoriesSollicitationResponse struct {
	CodeRetour int32                    `json:"codeRetour"`
	Libelle    string                   `json:"libelle"`
	Categories []CategorieSollicitation `json:"listeCategories"`
	Pagination *PaginationResponse      `json:"parametresRetour"`
}

// La méthode RechercherCategoriesSollicitation permet de rechercher les valeurs
// du référentiel catégorie Sollicitation paramétrées pour le mode connecté
func (s *TransversesService) RechercherCategoriesSollicitation(ctx context.Context, opts ListeCategoriesSollicitationOptions) (*ListeCategoriesSollicitationResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/categorieSollicitation", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeCategoriesSollicitationResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListeSousCategoriesSollicitationOptions est la structure de données utlisée
// pour appeler la méthode RechercherSousCategoriesSollicitation.
type ListeSousCategoriesSollicitationOptions struct {
	Code                 string             `json:"code,omitempty"`
	EstActif             bool               `json:"estActif,omitempty"`
	IdTechniqueCategorie int64              `json:"idTechniqueCategorie,omitempty"`
	Libelle              string             `json:"libelle,omitempty"`
	PaginationOptions    *PaginationOptions `json:"pagination,omitempty"`
}

// ListeSousCategoriesSollicitationResponse est la structure de données représentant
// la réponse de la méthode RechercherSousCategoriesSollicitation.
type ListeSousCategoriesSollicitationResponse struct {
	CodeRetour     int32                        `json:"codeRetour"`
	Libelle        string                       `json:"libelle"`
	SousCategories []SousCategoriesSolliciation `json:"listeSousCategories"`
	Pagination     *PaginationResponse          `json:"parametresRetour"`
}

// SousCategoriesSolliciation est la structure de données représentant
// des liste de catégories et de sous-catégories de sollicitation.
type SousCategoriesSolliciation struct {
	Categories     []CategorieSollicitation     `json:"categorie"`
	SousCategories []SousCategorieSollicitation `json:"ssCategorie"`
}

// CategorieSollicitation est la structure de données représentant
// une catégorie de sollicitation.
type CategorieSollicitation struct {
	Code        string `json:"codeCategorie"`
	Libelle     string `json:"libelleCategorie"`
	IdTechnique int64  `json:"idTechniqueCategorie"`
}

// SousCategorieSollicitation est la structure de données représentant
// une sous-catégorie de sollicitation.
type SousCategorieSollicitation struct {
	Code                     string     `json:"code"`
	DateCreation             *time.Time `json:"dateCreation"`
	DateDerniereModification *time.Time `json:"dateDerniereModification"`
	EstActif                 bool       `json:"estActif"`
	IdTechnique              int64      `json:"idTechniqueCategorie"`
	Libelle                  string     `json:"libelle"`
}

// Le service RechercherSousCategoriesSollicitation permet de rechercher les
// valeurs du référentiel Sous-catégorie Sollicitation paramétrées pour le mode connecté.
func (s *TransversesService) RechercherSousCategoriesSollicitation(ctx context.Context, opts ListeSousCategoriesSollicitationOptions) (*ListeSousCategoriesSollicitationResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/sousCategorieSollicitation", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeSousCategoriesSollicitationResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
