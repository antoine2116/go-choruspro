package choruspro

import (
	"context"
	"net/http"
	"time"
)

// RechercherDemandePaiementOptions est la structure de données utilisée pour appeler
// les méthodes ConsulterFactureParRecipiendaire et ConsulterFactureParValideur.
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

// StructureDestinataireParam est la structure de données représentant
// les paramètres de recherche sur la structure destinataire.
// Elle est utilisée dans la structure RechercherDemandePaiementOptions.
type StructureDestinataireParam struct {
	IdStructure              int64  `json:"idStructureDestinataire,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureDestinataire,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureDestinataire,omitempty"`
}

// StructureFournisseurParam est la structure de données représentant
// les paramètres de recherche sur la structure fournisseur.
// Elle est utilisée dans la structure RechercherDemandePaiementOptions.
type StructureFournisseurParam struct {
	IdStructure              int64  `json:"idStructureFournisseur,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureFournisseur,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureFournisseur,omitempty"`
}

// StructureMoaParam est la structure de données représentant
// les paramètres de recherche sur la structure MOA.
// Elle est utilisée dans la structure RechercherDemandePaiementOptions.
type StructureMoaParam struct {
	IdStructure              int64  `json:"idStructureMoa,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureMoa,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureMoa,omitempty"`
}

// StructureMoeParam est la structure de données représentant
// les paramètres de recherche sur la structure MOE.
// Elle est utilisée dans la structure RechercherDemandePaiementOptions.
type StructureMoeParam struct {
	IdStructure              int64  `json:"idStructureMoe,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureMoe,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureMoe,omitempty"`
}

// StructureValideurParam est la structure de données représentant
// les paramètres de recherche sur la structure valideur.
// Elle est utilisée dans la structure RechercherDemandePaiementOptions.
type StructureValideurParam struct {
	IdStructure              int64  `json:"idStructureValideur,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructureValideur,omitempty"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructureValideur,omitempty"`
}

// DemandePaie est la structure de données représentant une demande de paiement.
// Elle est utilisée dans la structure RechercherDemandePaiementResponse.
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

// RechercherDemandePaiementResponse est la structure de données représentant
// la réponse de la méthode RechercherDemandePaiement.
type RechercherDemandePaiementResponse struct {
	CodeRetour       int                `json:"codeRetour,omitempty"`
	Libelle          string             `json:"libelle,omitempty"`
	DemandesPaiement *[]DemandePaiement `json:"listeDemandePaiement,omitempty"`
}

// La méthode RechercherDemandePaiement permet de retourner des demandes de
// paiement en base archive et en base vivante pour l'utilisateur connecté
// en fonction de certains critères (service exposé ne fonctionnant pas
// sur l'espace de qualification).
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

// RechercherFactureParRecipiendaireOptions est la structure de données représentant
// les critères de recherche de factures pour un récipiendaire.
// Elle est utilisée par la méthode RechercherFactureParRecipiendaire.
type RechercherFactureParRecipiendaireOptions struct {
	CadreFacturation                  CadreFac           `json:"cadreFacturation,omitempty"`
	FactureTelechargeeParDestinataire bool               `json:"factureTelechargeeParDestinataire,omitempty"`
	IdDestinataire                    int64              `json:"idDestinataire,omitempty"`
	IdServiceExecutant                int64              `json:"idServiceExecutant,omitempty"`
	IdStructureValideur               int64              `json:"idStructureValideur,omitempty"`
	IdUtilisateurCourant              int64              `json:"idUtilisateurCourant,omitempty"`
	ListeFournisseurs                 []FournisseurParam `json:"listeFournisseurs,omitempty"`
	ListeTypesFacture                 []string           `json:"listeTypeFacture,omitempty"`
	ListeTypesFactureTravaux          []string           `json:"listeTypeFactureTravaux,omitempty"`
	MontantAPayerMax                  float32            `json:"montantApayerMax,omitempty"`
	MontantAPayerMin                  float32            `json:"montantApayerMin,omitempty"`
	MontantHTMax                      float32            `json:"montantHTMax,omitempty"`
	MontantHTMin                      float32            `json:"montantHTMin,omitempty"`
	MontantTTCMax                     float32            `json:"montantTTCMax,omitempty"`
	MontantTTCMin                     float32            `json:"montantTTCMin,omitempty"`
	NumeroBonCommande                 string             `json:"numeroBonCommande,omitempty"`
	NumeroFacture                     string             `json:"numeroFacture,omitempty"`
	NumeroMarche                      string             `json:"numeroMarche,omitempty"`
	Pagination                        *PaginationOptions `json:"paramRecherche,omitempty"`
	PeriodeDateDepotAu                *Date              `json:"periodeDateDepotAu,omitempty"`
	PeriodeDateDepotDu                *Date              `json:"periodeDateDepotDu,omitempty"`
	PeriodeDateFactureAu              *Date              `json:"periodeDateFactureAu,omitempty"`
	PeriodeDateFactureDu              *Date              `json:"periodeDateFactureDu,omitempty"`
	PeriodeDateHeureEtatCourantAu     *time.Time         `json:"periodeDateHeureEtatCourantAu,omitempty"`
	PeriodeDateHeureEtatCourantDu     *time.Time         `json:"periodeDateHeureEtatCourantDu,omitempty"`
	RechercheSirenFournisseur         bool               `json:"rechercheSirenFournisseur,omitempty"`
	RecupererTaille                   bool               `json:"recupererTaille,omitempty"`
	StatutCourant                     []string           `json:"statutCourant,omitempty"`
	TypeDemandePaiement               string             `json:"typeDemandePaiement,omitempty"`
}

// FournisseurParam est la structure de données représentant
// les critères de recherche de factures pour un récipiendaire.
// Elle est utilisée dans la méthode RechercherFactureParRecipiendaire.
type FournisseurParam struct {
	IdFournisseur  int64   `json:"idFournisseur,omitempty"`
	ListeIdService []int64 `json:"listeIdService,omitempty"`
}

// FactureParRecipiendaire est la structure de données représentant
// une facture pour un récipiendaire.
// Elle est utilisée dans la méthode RechercherFactureParRecipiendaire.
type FactureParRecipiendaire struct {
	CodeDestinataire                  string          `json:"codeDestinataire,omitempty"`
	CodeFournisseur                   string          `json:"codeFournisseur,omitempty"`
	CodeMOA                           string          `json:"codeMOA,omitempty"`
	CodeMOE                           string          `json:"codeMOE,omitempty"`
	CodeServiceExecutant              string          `json:"codeServiceExecutant,omitempty"`
	CodeServiceFournisseur            string          `json:"codeServiceFournisseur,omitempty"`
	CommentaireEtatCourant            string          `json:"commentaireEtatCourant,omitempty"`
	DateDepot                         *Date           `json:"dateDepot,omitempty"`
	DateFacture                       *Date           `json:"dateFacture,omitempty"`
	DateHeureEtatCourant              *time.Time      `json:"dateHeureEtatCourant,omitempty"`
	DesignationDestinataire           string          `json:"designationDestinataire,omitempty"`
	DesignationFournisseur            string          `json:"designationFournisseur,omitempty"`
	DesignationMOA                    string          `json:"designationMOA,omitempty"`
	DesignationMOE                    string          `json:"designationMOE,omitempty"`
	Devise                            string          `json:"devise,omitempty"`
	FactureTelechargeeParDestinataire bool            `json:"factureTelechargeeParDestinataire,omitempty"`
	IdDestinataire                    int64           `json:"idDestinataire,omitempty"`
	IdFacture                         int64           `json:"idFacture,omitempty"`
	IdServiceExecutant                int64           `json:"idServiceExecutant,omitempty"`
	MontantAPayer                     float32         `json:"montantAPayer,omitempty"`
	MontantHT                         float32         `json:"montantHT,omitempty"`
	MontantTTC                        float32         `json:"montantTTC,omitempty"`
	NomServiceExecutant               string          `json:"nomServiceExecutant,omitempty"`
	NomServiceFournisseur             string          `json:"nomServiceFournisseur,omitempty"`
	NumeroBonCommande                 string          `json:"numeroBonCommande,omitempty"`
	NumeroFacture                     string          `json:"numeroFacture,omitempty"`
	NumeroMarche                      string          `json:"numeroMarche,omitempty"`
	Statut                            string          `json:"statut,omitempty"`
	Taille                            int64           `json:"taille,omitempty"`
	TypeDemandePaiement               string          `json:"typeDemandePaiement,omitempty"`
	TypeFacture                       TypeFacture     `json:"typeFacture,omitempty"`
	TypeFactureTravaux                string          `json:"typeFactureTravaux,omitempty"`
	TypeIdentifiantFournisseur        TypeIdentifiant `json:"typeIdentifiantFournisseur,omitempty"`
	TypeIdentifiantMOA                TypeIdentifiant `json:"typeIdentifiantMOA,omitempty"`
	TypeIdentifiantMOE                TypeIdentifiant `json:"typeIdentifiantMOE,omitempty"`
}

// RechercherFactureParRecipiendaireResponse est la structure de données représentant
// la réponse de la méthode RechercherFactureParRecipiendaire.
type RechercherFactureParRecipiendaireResponse struct {
	CodeRetour         int                        `json:"codeRetour,omitempty"`
	Libelle            string                     `json:"libelle,omitempty"`
	Factures           *[]FactureParRecipiendaire `json:"listeFactures,omitempty"`
	NbResultatsParPage int                        `json:"nbResultatsParPage,omitempty"`
	PageCourante       int                        `json:"pageCourante,omitempty"`
	Pages              int                        `json:"pages,omitempty"`
	Total              int                        `json:"total,omitempty"`
}

// La méthode RechercherFactureParRecipiendaire permet d'afficher les factures reçues
// correspondant aux paramètres de recherche renseignés. Il est notamment possible
// de sélectionner des factures en fonction de l'état de la facture (téléchargée, non téléchargée).
func (s *FacturesService) RechercherFactureParRecipiendaire(ctx context.Context, opts RechercherFactureParRecipiendaireOptions) (*RechercherFactureParRecipiendaireResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/rechercher/recipiendaire", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherFactureParRecipiendaireResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// RechercherFactureParValideurOptions est la structure de données représentant
// les critères de recherche de factures pour un valideur.
// Elle est utilisée dans la méthode RechercherFactureParValideur.
type RechercherFactureParValideurOptions struct {
	CadreFacturation           CadreDeFacturation `json:"cadreFacturation"`
	IdDestinataire             int64              `json:"idDestinataire"`
	IdFournisseur              int64              `json:"idFournisseur"`
	IdServiceExecutant         int64              `json:"idServiceExecutant"`
	IdServiceFournisseur       int64              `json:"idServiceFournisseur"`
	IdStructureMoa             int64              `json:"idStructureMoa"`
	IdStructureMoe             int64              `json:"idStructureMoe"`
	IdStructureValideur        int64              `json:"idStructureValideur"`
	IdUtilisateurCourant       int64              `json:"idUtilisateurCourant"`
	ListeTypeFacture           []string           `json:"listeTypeFacture"`
	ListeTypeFactureTravaux    []string           `json:"listeTypeFactureTravaux"`
	MontantAPayerMax           float32            `json:"montantAPayerMax"`
	MontantAPayerMin           float32            `json:"montantAPayerMin"`
	MontantHTMax               float32            `json:"montantHTMax"`
	MontantHTMin               float32            `json:"montantHTMin"`
	MontantTTCMax              float32            `json:"montantTTCMax"`
	MontantTTCMin              float32            `json:"montantTTCMin"`
	NumeroEngagement           string             `json:"numeroEngagement"`
	NumeroFacture              string             `json:"numeroFacture"`
	NumeroMarche               string             `json:"numeroMarche"`
	Pagination                 *PaginationOptions `json:"pagination"`
	PeriodeDateDepotAu         *Date              `json:"periodeDateDepotAu"`
	PeriodeDateDepotDu         *Date              `json:"periodeDateDepotDu"`
	PeriodeDateFactureAu       *time.Time         `json:"periodeDateFactureAu"`
	PeriodeDateFactureDu       *time.Time         `json:"periodeDateFactureDu"`
	RechercheSirenDestinataire bool               `json:"rechercheSirenDestinataire"`
	RechercheSirenFournisseur  bool               `json:"rechercheSirenFournisseur"`
	StatutCourant              []string           `json:"statutCourant"`
	TypeDemandePaiement        string             `json:"typeDemandePaiement"`
}

// FactureParValideur est la structure de données représentant
// une facture pour un valideur.
type FactureParValideur struct {
	CodeServiceExecutant        string          `json:"codeServiceExecutant"`
	CodeServiceFournisseur      string          `json:"codeServiceFournisseur"`
	CodeServiceMoa              string          `json:"codeServiceMoa"`
	CodeServiceMoe              string          `json:"codeServiceMoe"`
	DateDepot                   *Date           `json:"dateDepot"`
	DateFacture                 *Date           `json:"dateFacture"`
	DesignationDestinataire     string          `json:"designationDestinataire"`
	DesignationFournisseur      string          `json:"designationFournisseur"`
	DesignationMoa              string          `json:"designationMoa"`
	DesignationMoe              string          `json:"designationMoe"`
	DesignationValideur         string          `json:"designationValideur"`
	Devise                      string          `json:"devise"`
	IdDestinataire              int64           `json:"idDestinataire"`
	IdFacture                   int64           `json:"idFacture"`
	IdFournisseur               int64           `json:"idFournisseur"`
	IdMoa                       int64           `json:"idMoa"`
	IdMoe                       int64           `json:"idMoe"`
	IdServiceExecutant          int64           `json:"idServiceExecutant"`
	IdServiceMoa                int64           `json:"idServiceMoa"`
	IdServiceMoe                int64           `json:"idServiceMoe"`
	IdValideur                  int64           `json:"idValideur"`
	IdentifiantDestinataire     string          `json:"identifiantDestinataire"`
	IdentifiantFournisseur      string          `json:"identifiantFournisseur"`
	IdentifiantMoa              string          `json:"identifiantMoa"`
	IdentifiantMoe              string          `json:"identifiantMoe"`
	IdentifiantValideur         string          `json:"identifiantValideur"`
	MontantAPayer               float32         `json:"montantAPayer"`
	MontantHT                   float32         `json:"montantHT"`
	MontantTTC                  float32         `json:"montantTTC"`
	NomServiceExecutant         string          `json:"nomServiceExecutant"`
	NomServiceFournisseur       string          `json:"nomServiceFournisseur"`
	NomServiceMoa               string          `json:"nomServiceMoa"`
	NomServiceMoe               string          `json:"nomServiceMoe"`
	NumeroEngagement            string          `json:"numeroEngagement"`
	NumeroFacture               string          `json:"numeroFacture"`
	NumeroMarche                string          `json:"numeroMarche"`
	Statut                      string          `json:"statut"`
	TypeDemandePaiement         string          `json:"typeDemandePaiement"`
	TypeFacture                 TypeFacture     `json:"typeFacture"`
	TypeFactureTravaux          TypeIdentifiant `json:"typeFactureTravaux"`
	TypeIdentifiantDestinataire TypeIdentifiant `json:"typeIdentifiantDestinataire"`
	TypeIdentifiantFournisseur  TypeIdentifiant `json:"typeIdentifiantFournisseur"`
	TypeIdentifiantMoa          TypeIdentifiant `json:"typeIdentifiantMoa"`
	TypeIdentifiantMoe          TypeIdentifiant `json:"typeIdentifiantMoe"`
	TypeIdentifiantValideur     TypeIdentifiant `json:"typeIdentifiantValideur"`
}

// RechercherFactureParValideurResponse est la structure de données représentant
// la réponse de la méthode RechercherFactureParValideur.
type RechercherFactureParValideurResponse struct {
	CodeRetour int32                `json:"codeRetour"`
	Libelle    string               `json:"libelle"`
	Factures   []FactureParValideur `json:"listeFactures"`
	Pagination *PaginationResponse  `json:"parametresRetour"`
}

// Le service rechercherFactureParValideur permet de rechercher, en tant que
// valideur, les factures et/ou les factures de travaux correspondant aux paramètres.
func (s *FacturesService) RechercherFactureParValideur(ctx context.Context, opts RechercherFactureParValideurOptions) (*RechercherFactureParValideurResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/rechercher/valideur", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherFactureParValideurResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DestinataireParam est la structure de données représentant
// les paramètres de recherche d'un destinataire.
// Elle est utilisée dans la structure RechercherFactureParFournisseurOptions.
type DestinataireParam struct {
	IdDestinataire          int64   `json:"idDestinataire"`
	ListeIdServiceExecutant []int64 `json:"listeIdServiceExecutant"`
	RaisonSocialeStuctDmd   string  `json:"raisonSocialeStuctDmd"`
}

// RechercherFactureParFournisseurOptions est la structure de données représentant
// les critères de recherche d'une facture pour un fournisseur.
// Elle est utilisée dans la méthode RechercherFactureParFournisseur.
type RechercherFactureParFournisseurOptions struct {
	CadreFacturation               CadreDeFacturation  `json:"cadreFacturation"`
	CoordonneeBancaire             int64               `json:"coordonneeBancaire"`
	IdFournisseur                  int64               `json:"idFournisseur"`
	IdServiceFournisseur           int64               `json:"idServiceFournisseur"`
	IdStructureValideur            int64               `json:"idStructureValideur"`
	IdUtilisateurCourant           int64               `json:"idUtilisateurCourant"`
	IdUtilisateurCreateur          []int64             `json:"idUtilisateurCreateur"`
	ListeDestinataire              []DestinataireParam `json:"listeDestinataire"`
	ModeDepot                      string              `json:"modeDepot"`
	MontantAPayerMax               float32             `json:"montantApayerMax"`
	MontantAPayerMin               float32             `json:"montantApayerMin"`
	MontantHTMax                   float32             `json:"montantHTMax"`
	MontantHTMin                   float32             `json:"montantHTMin"`
	MontantTTCMax                  float32             `json:"montantTTCMax"`
	MontantTTCMin                  float32             `json:"montantTTCMin"`
	NumeroBonCommande              string              `json:"numeroBonCommande"`
	NumeroFacture                  string              `json:"numeroFacture"`
	NumeroFactureOrigine           string              `json:"numeroFactureOrigine"`
	NumeroFluxDepot                string              `json:"numeroFluxDepot"`
	NumeroMarche                   string              `json:"numeroMarche"`
	PeriodeDateDepotAu             string              `json:"periodeDateDepotAu"`
	PeriodeDateDepotDu             string              `json:"periodeDateDepotDu"`
	PeriodeDateFactureAu           *Date               `json:"periodeDateFactureAu"`
	PeriodeDateFactureDu           *Date               `json:"periodeDateFactureDu"`
	PeriodeDateHeureEtatCourantAu  *time.Time          `json:"periodeDateHeureEtatCourantAu"`
	PeriodeDateHeureEtatCourantDu  *time.Time          `json:"periodeDateHeureEtatCourantDu"`
	RechercheFactureParFournisseur *PaginationOptions  `json:"rechercheFactureParFournisseur"`
	RechercheSirenDestinataire     bool                `json:"rechercheSirenDestinataire"`
	RecupererTaille                bool                `json:"recupererTaille"`
	RejetTraite                    bool                `json:"rejetTraite"`
	StatutCourant                  []string            `json:"statutCourant"`
	TypeFacture                    TypeFacture         `json:"typeFacture"`
}

// FactureParFournisseur est la structure de données représentant
// une facture pour un fournisseur.
// Elle est utilisée dans la structure RechercherFactureParFournisseurResponse.
type FactureParFournisseur struct {
	AffactureurCode                         string          `json:"affactureurCode"`
	AffactureurRaisonSociale                string          `json:"affactureurRaisonSociale"`
	AffactureurTypeIdentifiant              string          `json:"affactureurTypeIdentifiant"`
	CodeDestinataire                        string          `json:"codeDestinataire"`
	CodeFournisseur                         string          `json:"codeFournisseur"`
	CodeServiceExecutant                    string          `json:"codeServiceExecutant"`
	CodeServiceFournisseur                  string          `json:"codeServiceFournisseur"`
	CodeValideur1                           string          `json:"codeValideur1"`
	CodeValideur2                           string          `json:"codeValideur2"`
	CommentaireEtatCourant                  string          `json:"commentaireEtatCourant"`
	CoordBancairesFournisseurCleIban        string          `json:"coordBancairesFournisseurCleIban"`
	CoordBancairesFournisseurCleRib         string          `json:"coordBancairesFournisseurCleRib"`
	CoordBancairesFournisseurCodeBanque     string          `json:"coordBancairesFournisseurCodeBanque"`
	CoordBancairesFournisseurCodePays       string          `json:"coordBancairesFournisseurCodePays"`
	CoordBancairesFournisseurCompteBancaire string          `json:"coordBancairesFournisseurCompteBancaire"`
	CoordBancairesFournisseurNomCb          string          `json:"coordBancairesFournisseurNomCb"`
	DateDepot                               *Date           `json:"dateDepot"`
	DateFacture                             *Date           `json:"dateFacture"`
	DateHeureEtatCourant                    *time.Time      `json:"dateHeureEtatCourant"`
	DateValidation1                         *Date           `json:"dateValidation1"`
	DateValidation2                         *Date           `json:"dateValidation2"`
	DesignationDestinataire                 string          `json:"designationDestinataire"`
	DesignationFournisseur                  string          `json:"designationFournisseur"`
	Devise                                  string          `json:"devise"`
	IdDestinataire                          int64           `json:"idDestinataire"`
	IdServiceExecutant                      int64           `json:"idServiceExecutant"`
	IdentifiantFactureCPP                   int64           `json:"identifiantFactureCPP"`
	ModeDepot                               string          `json:"modeDepot"`
	MontantAPayer                           float32         `json:"montantAPayer"`
	MontantHT                               float32         `json:"montantHT"`
	MontantTTC                              float32         `json:"montantTTC"`
	NomPrenomUtilisateurCreateur            string          `json:"nomPrenomUtilisateurCreateur"`
	NomServiceExecutant                     string          `json:"nomServiceExecutant"`
	NomServiceFournisseur                   string          `json:"nomServiceFournisseur"`
	NomValideur1                            string          `json:"nomValideur1"`
	NomValideur2                            string          `json:"nomValideur2"`
	NumeroBonCommande                       string          `json:"numeroBonCommande"`
	NumeroFacture                           string          `json:"numeroFacture"`
	NumeroFactureOrigine                    string          `json:"numeroFactureOrigine"`
	NumeroFluxDepot                         string          `json:"numeroFluxDepot"`
	NumeroMarche                            string          `json:"numeroMarche"`
	PrenomValideur1                         string          `json:"prenomValideur1"`
	PrenomValideur2                         string          `json:"prenomValideur2"`
	RaisonSocialeValideur1                  string          `json:"raisonSocialeValideur1"`
	RaisonSocialeValideur2                  string          `json:"raisonSocialeValideur2"`
	RejetTraite                             bool            `json:"rejetTraite"`
	Statut                                  string          `json:"statut"`
	Taille                                  int64           `json:"taille"`
	TypeDemandePaiement                     string          `json:"typeDemandePaiement"`
	TypeFacture                             string          `json:"typeFacture"`
	TypeIdentifiantFournisseur              TypeIdentifiant `json:"typeIdentifiantFournisseur"`
	TypeIdentifiantValideur1                TypeIdentifiant `json:"typeIdentifiantValideur1"`
	TypeIdentifiantValideur2                TypeIdentifiant `json:"typeIdentifiantValideur2"`
}

// RechercherFactureParFournisseurResponse est la structure de données représentant
// la réponse de la méthode RechercherFactureParFournisseur.
type RechercherFactureParFournisseurResponse struct {
	CodeRetour int32                   `json:"codeRetour"`
	Libelle    string                  `json:"libelle"`
	Factures   []FactureParFournisseur `json:"listeFactures"`
	Pagination *PaginationResponse     `json:"parametresRetour"`
}

// La méthode RechercherFactureParFournisseur permet d'afficher les factures
// émises correspondant aux paramètres de recherche renseignés.
func (s *FacturesService) RechercherFactureParFournisseur(ctx context.Context, opts RechercherFactureParFournisseurOptions) (*RechercherFactureParFournisseurResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/rechercher/fournisseur", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherFactureParFournisseurResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
