package choruspro

import (
	"context"
	"net/http"
	"time"
)

// AdressePostale est la structure de données représentant une adresse postale
// d'une structure ou d'un service.
type AdressePostale struct {
	Adresse            string `json:"adresse"`
	CodePostal         string `json:"codePostal"`
	ComplementAdresse1 string `json:"complementAdresse1"`
	ComplementAdresse2 string `json:"complementAdresse2"`
	Fax                string `json:"fax"`
	IndicatifFax       string `json:"indicatifFax"`
	IndicatifTelephone string `json:"indicatifTelephone"`
	Pays               string `json:"pays"`
	Telephone          string `json:"telephone"`
	Ville              string `json:"ville"`
}

// StructureDupliquee est la structure de données représentant une l'identifiant
// d'une structure dupliquée.
type StructureDupliquee struct {
	StructureDupliquee int64 `json:"structureDupliquee"`
}

// InformationGeneralesStructure est la structure de données représentant les
// informations générales d'une structure.
type InformationGeneralesStructure struct {
	DateExpirationMotPasseCompteTech *time.Time           `json:"dateExpirationMotPasseCompteTech"`
	Email                            string               `json:"emailStructure"`
	EstValideurDeleguee              bool                 `json:"estValideurDeleguee"`
	IdCPP                            int64                `json:"idStructureCPP"`
	IdOrigine                        int64                `json:"idStructureOrigine"`
	Identifiant                      string               `json:"identifiantStructure"`
	Libelle                          string               `json:"libelleStructure"`
	ListeStructureDupliquees         []StructureDupliquee `json:"listeStructureDupliquees"`
	Nom                              string               `json:"nomStructure"`
	NumeroRcs                        string               `json:"numeroRcsStructure"`
	NumeroTVA                        string               `json:"numeroTVA"`
	Prenom                           string               `json:"prenomStructure"`
	RaisonSociale                    string               `json:"raisonSocialeStructure"`
	PriveePublique                   string               `json:"structurePriveePublique"`
	TypeIdentifiant                  string               `json:"typeIdentifiantStructure"`
}

// ParametrageStructure est la structure de données représentant les
// paramètres d'une structure.
type ParametrageStructure struct {
	CodeServiceDoitEtreRenseigne       bool `json:"codeServiceDoitEtreRenseigne"`
	ConnexionEDI                       bool `json:"connexionEDI"`
	EstMoaUniquement                   bool `json:"estMoaUniquement"`
	EstValideurDelegue                 bool `json:"estValideurDelegue"`
	GestionNumeroEJOuCodeService       bool `json:"gestionNumeroEJOuCodeService"`
	GetaMoa                            bool `json:"getaMoa"`
	NonDiffusibleInsee                 bool `json:"nonDiffusibleInsee"`
	NumeroEJDoitEtreRenseigne          bool `json:"numeroEJDoitEtreRenseigne"`
	RecevoirDonneesViaEDI              bool `json:"recevoirDonneesViaEDI"`
	StatutMiseEnPaiementNestPasRemonte bool `json:"statutMiseEnPaiementNestPasRemonte"`
}

// ConsulterStructureResponse est la structure de données représentant la
// réponse de la méthode ConsulterStructure.
type ConsulterStructureResponse struct {
	CodeRetour           int32                          `json:"codeRetour"`
	Libelle              string                         `json:"libelle"`
	AdressePostale       *AdressePostale                `json:"adressePostaleDuSiege"`
	InformationGenerales *InformationGeneralesStructure `json:"informationGenerales"`
	Parametrage          *ParametrageStructure          `json:"parametres"`
}

// ConsulterStructureOptions est la structure de données utilisée pour
// appeler la méthode ConsulterStructure.
type ConsulterStructureOptions struct {
	CodeLangue  CodeLangue `json:"codeLangue"`
	IdStructure int64      `json:"idStructure"`
}

// La méthode ConsulterStructure permet de consulter les informations relatives
// à une structure à laquelle l'utilisateur est rattaché ou une structure publique.
func (s *StructuresService) ConsulterStructure(ctx context.Context, opts ConsulterStructureOptions) (*ConsulterStructureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/structures/v1/consulter", opts)
	if err != nil {
		return nil, err
	}

	res := new(ConsulterStructureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// InformationsGeneralesService est la structure de données représentant les
// informations générales d'un service.
type InformationsGeneralesService struct {
	Code        string `json:"codeService"`
	Description string `json:"descriptionService"`
	Nom         string `json:"nomService"`
}

// ParametrageService est la structure de données représentant les paramètres
// d'un service.
type ParametrageService struct {
	DateCreation      *Datetime `json:"dateCreation"`
	DateDebutValidite *Datetime `json:"dateDebutValidite"`
	DateFinValidite   *Datetime `json:"dateFinValidite"`
	DateModification  *Datetime `json:"dateModification"`
	MiseEnPaiement    bool      `json:"miseEnPaiement"`
	NumeroEngagement  bool      `json:"numeroEngagement"`
}

// ConsulterServiceResponse est la structure de données représentant la réponse
// de la méthode ConsulterService.
type ConsulterServiceResponse struct {
	CodeRetour            int32                         `json:"codeRetour"`
	Libelle               string                        `json:"libelle"`
	AdressePostale        *AdressePostale               `json:"adressePostale"`
	InformationsGenerales *InformationsGeneralesService `json:"informationsGenerales"`
	Parametrage           *ParametrageService           `json:"parametres"`
}

// ConsulterServiceOptions est la structure de données utilisée pour appeler
// la méthode ConsulterService.
type ConsulterServiceOptions struct {
	CodeLangue  CodeLangue `json:"codeLangue"`
	IdService   int64      `json:"idService"`
	IdStructure int64      `json:"idStructure"`
}

// La méthode ConsulterService permet de consulter :
// - les informations relatives à un service auquel l'utilisateur est rattaché
// - les informations relatives à un service d'une structure publique.
func (s *StructuresService) ConsulterService(ctx context.Context, opts ConsulterServiceOptions) (*ConsulterServiceResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/structures/v1/consulter/service", opts)
	if err != nil {
		return nil, err
	}

	res := new(ConsulterServiceResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
