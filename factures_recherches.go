package choruspro

import (
	"context"
	"net/http"
	"time"
)

type RechercherDemandePaiementOptions struct {
	CadreFacturation CadreFac `json:"cadreFacturation,omitempty"`

	// 2 value possible : "DEMANDE_PAIEMENT" or "DP_ARCHIVE"
	DpOuDpArchive string `json:"dpOuDpArchive,omitempty"`

	IdServiceDestinataire         int64           `json:"idServiceDestinataire,omitempty"`
	IdServiceFournisseur          int64           `json:"idServiceFournisseur,omitempty"`
	IdServiceMoa                  int64           `json:"idServiceMoa,omitempty"`
	IdServiceMoe                  int64           `json:"idServiceMoe,omitempty"`
	MontantAPayerMax              float32         `json:"montantAPayerMax,omitempty"`
	MontantAPayerMin              float32         `json:"montantAPayerMin,omitempty"`
	MontantHTMax                  float32         `json:"montantHTMax,omitempty"`
	MontantHTMin                  float32         `json:"montantHTMin,omitempty"`
	MontantTTCMax                 float32         `json:"montantTTCMax,omitempty"`
	MontantTTCMin                 float32         `json:"montantTTCMin,omitempty"`
	NumeroBonCommande             string          `json:"numeroBonCommande,omitempty"`
	NumeroDossierFacturation      string          `json:"numeroDossierFacturation,omitempty"`
	NumeroDpMandat                string          `json:"numeroDpMandat,omitempty"`
	NumeroFacture                 string          `json:"numeroFacture,omitempty"`
	NumeroFactureMemoire          string          `json:"numeroFactureMemoire,omitempty"`
	NumeroFactureOrigine          string          `json:"numeroFactureOrigine,omitempty"`
	NumeroMarche                  string          `json:"numeroMarche,omitempty"`
	PeriodeDateDepotAu            *Date           `json:"periodeDateDepotAu,omitempty"`
	PeriodeDateDepotDu            *Date           `json:"periodeDateDepotDu,omitempty"`
	PeriodeDateFournisseurAu      *Date           `json:"periodeDateFournisseurAu,omitempty"`
	PeriodeDateFournisseurDu      *Date           `json:"periodeDateFournisseurDu,omitempty"`
	PeriodeDateHeureEtatCourantAu *time.Time      `json:"periodeDateHeureEtatCourantAu,omitempty"`
	PeriodeDateHeureEtatCourantDu *time.Time      `json:"periodeDateHeureEtatCourantDu,omitempty"`
	RechercheSirenDestinataire    bool            `json:"rechercheSirenDestinataire,omitempty"`
	RechercheSirenFournisseur     bool            `json:"rechercheSirenFournisseur,omitempty"`
	RoleUtilisateur               RoleUtilisateur `json:"roleUtilisateur,omitempty"`
	StatutCourant                 string          `json:"statutCourant,omitempty"`
	TypeDemandePaiement           string          `json:"typeDemandePaiement,omitempty"`
	TypeFacture                   TypeFacture     `json:"typeFacture,omitempty"`

	// Values available using Transverses.RecupererTypesFactureTravaux()
	TypeFactureTravaux string `json:"typeFactureTravaux,omitempty"`

	StructureDestinataire *StructureDestinataireParam `json:"structureDestinataire,omitempty"`
	StructureFournisseur  *StructureFournisseurParam  `json:"structureFournisseur,omitempty"`
	StructureMoa          *StructureMoaParam          `json:"structureMoa,omitempty"`
	StructureMoe          *StructureMoeParam          `json:"structureMoe,omitempty"`
	StructureValideur     *StructureValideurParam     `json:"structureValideur,omitempty"`
}

type StructureDestinataireParam struct {
	IdStructure              int64  `json:"idStructureDestinataire,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureDestinataire,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureDestinataire,omitempty"`
}

type StructureFournisseurParam struct {
	IdStructure              int64  `json:"idStructureFournisseur,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureFournisseur,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureFournisseur,omitempty"`
}

type StructureMoaParam struct {
	IdStructure              int64  `json:"idStructureMoa,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureMoa,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureMoa,omitempty"`
}

type StructureMoeParam struct {
	IdStructure              int64  `json:"idStructureMoe,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureMoe,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureMoe,omitempty"`
}

type StructureValideurParam struct {
	IdStructure              int64  `json:"idStructureValideur,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureValideur,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureValideur,omitempty"`
}

type RechercherDemandePaiementResponse struct {
	CodeRetour       int                `json:"codeRetour,omitempty"`
	Libelle          string             `json:"libelle,omitempty"`
	DemandesPaiement *[]DemandePaiement `json:"listeDemandePaiement,omitempty"`
}

type DemandePaiement struct {
	CadreFacturation                 CadreFac    `json:"cadreFacturation,omitempty"`
	CodeJuridiction                  string      `json:"codeJuridiction,omitempty"`
	CodeServiceDestinataire          string      `json:"codeServiceDestinataire,omitempty"`
	CodeServiceSollicitation         string      `json:"codeServiceSollicitation,omitempty"`
	CommentaireEtatCourant           string      `json:"commentaireEtatCourant,omitempty"`
	DateDepot                        *Date       `json:"dateDepot,omitempty"`
	DateFournisseur                  *Date       `json:"dateFournisseur,omitempty"`
	DateHeureEtatCourant             *time.Time  `json:"dateHeureEtatCourant,omitempty"`
	DpOuDpArchive                    string      `json:"dpOuDpArchive,omitempty"`
	EtatCourant                      string      `json:"etatCourant,omitempty"`
	IdDemandePaiement                int64       `json:"idDemandePaiement,omitempty"`
	IdDossierFacturation             int64       `json:"idDossierFacturation,omitempty"`
	IdServiceDestinataire            int64       `json:"idServiceDestinataire,omitempty"`
	IdServiceFournisseur             int64       `json:"idServiceFournisseur,omitempty"`
	IdServiceMoa                     int64       `json:"idServiceMoa,omitempty"`
	IdServiceMoe                     int64       `json:"idServiceMoe,omitempty"`
	IdStructureDestinataire          int64       `json:"idStructureDestinataire,omitempty"`
	IdStructureFournisseur           int64       `json:"idStructureFournisseur,omitempty"`
	IdStructureMoa                   int64       `json:"idStructureMoa,omitempty"`
	IdStructureMoe                   int64       `json:"idStructureMoe,omitempty"`
	IdStructureValideur1             int64       `json:"idStructureValideur1,omitempty"`
	IdStructureValideur2             int64       `json:"idStructureValideur2,omitempty"`
	IdentifiantStructureDestinataire string      `json:"identifiantStructureDestinataire,omitempty"`
	IdentifiantStructureFournisseur  string      `json:"identifiantStructureFournisseur,omitempty"`
	LibelleCodePostaleDeptDR         string      `json:"libelleCodePostaleDeptDR,omitempty"`
	MontantAPayer                    float32     `json:"montantAPayer,omitempty"`
	MontantHT                        float32     `json:"montantHT,omitempty"`
	MontantTTC                       float32     `json:"montantTTC,omitempty"`
	NomServiceDestinataire           string      `json:"nomServiceDestinataire,omitempty"`
	NomServiceSollicitation          string      `json:"nomServiceSollicitation,omitempty"`
	NombreResultatTotal              string      `json:"nombreResultatTotal,omitempty"`
	NumeroBonCommande                string      `json:"numeroBonCommande,omitempty"`
	NumeroDemandePaiement            string      `json:"numeroDemandePaiement,omitempty"`
	NumeroDossierFacturation         string      `json:"numeroDossierFacturation,omitempty"`
	NumeroDpMandat                   string      `json:"numeroDpMandat,omitempty"`
	NumeroFactureOrigine             string      `json:"numeroFactureOrigine,omitempty"`
	NumeroMarche                     string      `json:"numeroMarche,omitempty"`
	RaisonSocialeDeptDR              string      `json:"raisonSocialeDeptDR,omitempty"`
	RaisonSocialeJuridiction         string      `json:"raisonSocialeJuridiction,omitempty"`
	RaisonSocialeTGI                 string      `json:"raisonSocialeTGI,omitempty"`
	SiretDestinataireQualif          string      `json:"siretDestinataireQualif,omitempty"`
	SiretFournisseurQualif           string      `json:"siretFournisseurQualif,omitempty"`
	TypeDemandePaiement              string      `json:"typeDemandePaiement,omitempty"`
	TypeFacture                      TypeFacture `json:"typeFacture,omitempty"`
	TypeFactureTravaux               string      `json:"typeFactureTravaux,omitempty"`
}

func (s *FacturesService) RechercherDemandePaiement(ctx context.Context, opts RechercherDemandePaiementOptions) (*RechercherDemandePaiementResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/rechercher/demandePaiement", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherDemandePaiementResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type RechercherFactureRecipiendaireOptions struct {
	CadreFacturation                  CadreFac            `json:"cadreFacturation,omitempty"`
	FactureTelechargeeParDestinataire bool                `json:"factureTelechargeeParDestinataire,omitempty"`
	IdDestinataire                    int64               `json:"idDestinataire,omitempty"`
	IdServiceExecutant                int64               `json:"idServiceExecutant,omitempty"`
	IdStructureValideur               int64               `json:"idStructureValideur,omitempty"`
	IdUtilisateurCourant              int64               `json:"idUtilisateurCourant,omitempty"`
	ListeFournisseurs                 []*FournisseurParam `json:"listeFournisseurs,omitempty"`
	ListeTypesFacture                 []string            `json:"listeTypeFacture,omitempty"`
	ListeTypesFactureTravaux          []string            `json:"listeTypeFactureTravaux,omitempty"`
	MontantAPayerMax                  float32             `json:"montantApayerMax,omitempty"`
	MontantAPayerMin                  float32             `json:"montantApayerMin,omitempty"`
	MontantHTMax                      float32             `json:"montantHTMax,omitempty"`
	MontantHTMin                      float32             `json:"montantHTMin,omitempty"`
	MontantTTCMax                     float32             `json:"montantTTCMax,omitempty"`
	MontantTTCMin                     float32             `json:"montantTTCMin,omitempty"`
	NumeroBonCommande                 string              `json:"numeroBonCommande,omitempty"`
	NumeroFacture                     string              `json:"numeroFacture,omitempty"`
	NumeroMarche                      string              `json:"numeroMarche,omitempty"`
	Pagination                        *PaginationOptions  `json:"paramRecherche,omitempty"`
	PeriodeDateDepotAu                *Date               `json:"periodeDateDepotAu,omitempty"`
	PeriodeDateDepotDu                *Date               `json:"periodeDateDepotDu,omitempty"`
	PeriodeDateFactureAu              *Date               `json:"periodeDateFactureAu,omitempty"`
	PeriodeDateFactureDu              *Date               `json:"periodeDateFactureDu,omitempty"`
	PeriodeDateHeureEtatCourantAu     *time.Time          `json:"periodeDateHeureEtatCourantAu,omitempty"`
	PeriodeDateHeureEtatCourantDu     *time.Time          `json:"periodeDateHeureEtatCourantDu,omitempty"`
	RechercheSirenFournisseur         bool                `json:"rechercheSirenFournisseur,omitempty"`
	RecupererTaille                   bool                `json:"recupererTaille,omitempty"`
	StatutCourant                     []string            `json:"statutCourant,omitempty"`
	TypeDemandePaiement               string              `json:"typeDemandePaiement,omitempty"`
}

type FournisseurParam struct {
	IdFournisseur  int64   `json:"idFournisseur,omitempty"`
	ListeIdService []int64 `json:"listeIdService,omitempty"`
}

type RechercherFactureRecipiendaireResponse struct {
	CodeRetour         int                 `json:"codeRetour,omitempty"`
	Libelle            string              `json:"libelle,omitempty"`
	Factures           *[]FactureRecherche `json:"listeFactures,omitempty"`
	NbResultatsParPage int                 `json:"nbResultatsParPage,omitempty"`
	PageCourante       int                 `json:"pageCourante,omitempty"`
	Pages              int                 `json:"pages,omitempty"`
	Total              int                 `json:"total,omitempty"`
}

type FactureRecherche struct {
	CodeDestinataire                  string      `json:"codeDestinataire,omitempty"`
	CodeFournisseur                   string      `json:"codeFournisseur,omitempty"`
	CodeMOA                           string      `json:"codeMOA,omitempty"`
	CodeMOE                           string      `json:"codeMOE,omitempty"`
	CodeServiceExecutant              string      `json:"codeServiceExecutant,omitempty"`
	CodeServiceFournisseur            string      `json:"codeServiceFournisseur,omitempty"`
	CommentaireEtatCourant            string      `json:"commentaireEtatCourant,omitempty"`
	DateDepot                         *Date       `json:"dateDepot,omitempty"`
	DateFacture                       *Date       `json:"dateFacture,omitempty"`
	DateHeureEtatCourant              *time.Time  `json:"dateHeureEtatCourant,omitempty"`
	DesignationDestinataire           string      `json:"designationDestinataire,omitempty"`
	DesignationFournisseur            string      `json:"designationFournisseur,omitempty"`
	DesignationMOA                    string      `json:"designationMOA,omitempty"`
	DesignationMOE                    string      `json:"designationMOE,omitempty"`
	Devise                            string      `json:"devise,omitempty"`
	FactureTelechargeeParDestinataire bool        `json:"factureTelechargeeParDestinataire,omitempty"`
	IdDestinataire                    int64       `json:"idDestinataire,omitempty"`
	IdFacture                         int64       `json:"idFacture,omitempty"`
	IdServiceExecutant                int64       `json:"idServiceExecutant,omitempty"`
	MontantAPayer                     float32     `json:"montantAPayer,omitempty"`
	MontantHT                         float32     `json:"montantHT,omitempty"`
	MontantTTC                        float32     `json:"montantTTC,omitempty"`
	NomServiceExecutant               string      `json:"nomServiceExecutant,omitempty"`
	NomServiceFournisseur             string      `json:"nomServiceFournisseur,omitempty"`
	NumeroBonCommande                 string      `json:"numeroBonCommande,omitempty"`
	NumeroFacture                     string      `json:"numeroFacture,omitempty"`
	NumeroMarche                      string      `json:"numeroMarche,omitempty"`
	Statut                            string      `json:"statut,omitempty"`
	Taille                            int64       `json:"taille,omitempty"`
	TypeDemandePaiement               string      `json:"typeDemandePaiement,omitempty"`
	TypeFacture                       TypeFacture `json:"typeFacture,omitempty"`
	TypeFactureTravaux                string      `json:"typeFactureTravaux,omitempty"`
	TypeIdentifiantFournisseur        string      `json:"typeIdentifiantFournisseur,omitempty"`
	TypeIdentifiantMOA                string      `json:"typeIdentifiantMOA,omitempty"`
	TypeIdentifiantMOE                string      `json:"typeIdentifiantMOE,omitempty"`
}

func (s *FacturesService) RechercherFactureParRecipiendaire(ctx context.Context, opts RechercherFactureRecipiendaireOptions) (*RechercherFactureRecipiendaireResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/rechercher/recipiendaire", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherFactureRecipiendaireResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
