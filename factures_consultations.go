package choruspro

import (
	"context"
	"net/http"
)

type ConsulterFactureOptions struct {
	CodeLangue              CodeLangue                       `json:"codeLangue,omitempty"`
	IdUtilisateurCourant    int                              `json:"idUtilisateurCourant,omitempty"`
	IdentifiantFactureCPP   int                              `json:"identifiantFactureCPP,omitempty"`
	PaginationLignesPoste   *PaginationLignesPosteOptions    `json:"rechercheLignePoste,omitempty"`
	PaginationPiecesJointes *PaginationPiecesJointesOptions  `json:"recherchePj,omitempty"`
	PaginationLignesTVA     *PaginationLignesRecapTVAOptions `json:"rechercheligneTvaRecap,omitempty"`
}

type CadreDeFacturation struct {
	Code                     CadreFac        `json:"codeCadreFacturation,omitempty"`
	CodeServiceMoa           string          `json:"codeServiceMoa,omitempty"`
	CodeServiceMoe           string          `json:"codeServiceMoe,omitempty"`
	CodeValideur1            string          `json:"codeValideur1,omitempty"`
	CodeValideur2            string          `json:"codeValideur2,omitempty"`
	DateValidation1          *Date           `json:"dateValidation1,omitempty"`
	DateValidation2          *Date           `json:"dateValidation2,omitempty"`
	DesignationMoa           string          `json:"designationMoa,omitempty"`
	DesignationMoe           string          `json:"designationMoe,omitempty"`
	DesignationValideur2     string          `json:"designationValideur2,omitempty"`
	IdMoa                    int64           `json:"idMoa,omitempty"`
	IdMoe                    int64           `json:"idMoe,omitempty"`
	IdServiceMoa             int64           `json:"idServiceMoa,omitempty"`
	IdServiceMoe             int64           `json:"idServiceMoe,omitempty"`
	IdValideur1              int64           `json:"idValideur1,omitempty"`
	IdValideur2              int64           `json:"idValideur2,omitempty"`
	IdentifiantMoa           string          `json:"identifiantMoa,omitempty"`
	IdentifiantMoe           string          `json:"identifiantMoe,omitempty"`
	NomMoe                   string          `json:"nomMoe,omitempty"`
	NomServiceMoa            string          `json:"nomServiceMoa,omitempty"`
	NomServiceMoe            string          `json:"nomServiceMoe,omitempty"`
	NomValideur1             string          `json:"nomValideur1,omitempty"`
	NomValideur2             string          `json:"nomValideur2,omitempty"`
	PrenomMoe                string          `json:"prenomMoe,omitempty"`
	PrenomValideur1          string          `json:"prenomValideur1,omitempty"`
	PrenomValideur2          string          `json:"prenomValideur2,omitempty"`
	RaisonSocialeMoa         string          `json:"raisonSocialeMoa,omitempty"`
	RaisonSocialeMoe         string          `json:"raisonSocialeMoe,omitempty"`
	RaisonSocialeValideur1   string          `json:"raisonSocialeValideur1,omitempty"`
	RaisonSocialeValideur2   string          `json:"raisonSocialeValideur2,omitempty"`
	TypeIdentifiantMoa       TypeIdentifiant `json:"typeIdentifiantMoa,omitempty"`
	TypeIdentifiantMoe       TypeIdentifiant `json:"typeIdentifiantMoe,omitempty"`
	TypeIdentifiantValideur1 TypeIdentifiant `json:"typeIdentifiantValideur1,omitempty"`
	TypeIdentifiantValideur2 TypeIdentifiant `json:"typeIdentifiantValideur2,omitempty"`
}

type Affactureur struct {
	Code            string          `json:"affactureurCode,omitempty"`
	CodePays        string          `json:"affactureurCodePays,omitempty"`
	Id              int64           `json:"affactureurId,omitempty"`
	RaisonSociale   string          `json:"affactureurRaisonSociale,omitempty"`
	TypeIdentifiant TypeIdentifiant `json:"affactureurTypeIdentifiant,omitempty"`
}

type PieceJointeFacture struct {
	Designation        string `json:"pieceJointeDesignation,omitempty"`
	Extension          string `json:"pieceJointeExtension,omitempty"`
	Id                 int    `json:"pieceJointeId,omitempty"`
	IdLiaison          int    `json:"pieceJointeIdLiaison,omitempty"`
	NumeroLigneFacture int    `json:"pieceJointeNumeroLigneFacture,omitempty"`
	TypeCode           string `json:"pieceJointeTypeCode,omitempty"`
	TypeLibelle        string `json:"pieceJointeTypeLibelle,omitempty"`
}

type LignePoste struct {
	Denomination          string  `json:"lignePosteDenomination,omitempty"`
	MontantHtApresRemise  float32 `json:"lignePosteMontantHtApresRemise,omitempty"`
	MontantRemiseHT       float32 `json:"lignePosteMontantRemiseHT,omitempty"`
	MontantTtcApresRemise float32 `json:"lignePosteMontantTtcApresRemise,omitempty"`
	MontantTva            float32 `json:"lignePosteMontantTva,omitempty"`
	MontantUnitaireHT     float32 `json:"lignePosteMontantUnitaireHT,omitempty"`
	Numero                int     `json:"lignePosteNumero,omitempty"`
	Quantite              float32 `json:"lignePosteQuantite,omitempty"`
	Reference             string  `json:"lignePosteReference,omitempty"`
	TauxTva               string  `json:"lignePosteTauxTva,omitempty"`
	TauxTvaManuel         float32 `json:"lignePosteTauxTvaManuel,omitempty"`
	UniteCode             int     `json:"lignePosteUniteCode,omitempty"`
	UniteLibelle          string  `json:"lignePosteUniteLibelle,omitempty"`
}

type LigneTva struct {
	MontantBaseHtParTaux float32 `json:"ligneTvaMontantBaseHtParTaux,omitempty"`
	MontantTvaParTaux    float32 `json:"ligneTvaMontantTvaParTaux,omitempty"`
	TauxManuel           float32 `json:"ligneTvaTauxManuel,omitempty"`
	TauxRefCode          string  `json:"ligneTvaTauxRefCode,omitempty"`
	TauxRefId            int     `json:"ligneTvaTauxRefId,omitempty"`
	TauxRefLibelle       string  `json:"ligneTvaTauxRefLibelle,omitempty"`
	TauxRefValeur        float32 `json:"ligneTvaTauxRefValeur,omitempty"`
}

type MontantTotal struct {
	MontantAPayer                  float32 `json:"montantAPayer,omitempty"`
	MontantHtTotal                 float32 `json:"montantHtTotal,omitempty"`
	MontantRemiseGlobaleTTC        float32 `json:"montantRemiseGlobaleTTC,omitempty"`
	MontantTVA                     float32 `json:"montantTVA,omitempty"`
	MontantTtcAvantRemiseGlobalTTC float32 `json:"montantTtcAvantRemiseGlobalTTC,omitempty"`
	MontantTtcTotal                float32 `json:"montantTtcTotal,omitempty"`
	MotifRemiseGlobaleTTC          string  `json:"motifRemiseGlobaleTTC,omitempty"`
}

type PieceJointePrincipale struct {
	IdLiaison int `json:"idLiaisonPieceJointePrincipale,omitempty"`
	Id        int `json:"idPieceJointePrincipale,omitempty"`
}

type PiecePrecedente struct {
	CadreFacturation          string `json:"cadreFacturationPiecePrecedente,omitempty"`
	CodeServiceExecutant      string `json:"codeServiceExecutantPiecePrecedente,omitempty"`
	IdDestinataire            int    `json:"idDestinatairePiecePrecedente,omitempty"`
	Id                        int    `json:"idPiecePrecedente,omitempty"`
	IdServiceExecutant        int    `json:"idServiceExecutantPiecePrecedente,omitempty"`
	IdentifiantDestinataire   string `json:"identifiantDestinatairePiecePrecedente,omitempty"`
	NomServiceExecutant       string `json:"nomServiceExecutantPiecePrecedente,omitempty"`
	Numero                    string `json:"numeroPiecePrecedente,omitempty"`
	RaisonSocialeDestinataire string `json:"raisonSocialeDestinatairePiecePrecedente,omitempty"`
}

type PieceSuivante struct {
	CadreFacturation string `json:"cadreFacturationPieceSuivante,omitempty"`
	Id               int    `json:"idPieceSuivante,omitempty"`
	Numero           string `json:"numeroPieceSuivante,omitempty"`
}

type References struct {
	CodeDevise                 string      `json:"codeDeviseFacture,omitempty"`
	DateCreationFacture        *Datetime   `json:"dateCreationFacture,omitempty"`
	DateDepot                  *Date       `json:"dateDepot,omitempty"`
	DateEcheancePaiement       *Date       `json:"dateEcheancePaiement,omitempty"`
	DateFacture                *Date       `json:"dateFacture,omitempty"`
	LibelleDevise              string      `json:"libelleDeviseFacture,omitempty"`
	LibelleMotifExonerationTva string      `json:"libelleMotifExonerationTva,omitempty"`
	ModePaiement               string      `json:"modePaiement,omitempty"`
	MotifExonerationTva        string      `json:"motifExonerationTva,omitempty"`
	NumeroBonCommande          string      `json:"numeroBonCommande,omitempty"`
	NumeroDpMandat             string      `json:"numeroDpMandat,omitempty"`
	NumeroFactureOrigine       string      `json:"numeroFactureOrigine,omitempty"`
	NumeroMarche               string      `json:"numeroMarche,omitempty"`
	TailleTotalePiecesJointes  float32     `json:"tailleTotalePiecesJointes,omitempty"`
	TypeFacture                TypeFacture `json:"typeFacture,omitempty"`
	TypeFactureTravaux         string      `json:"typeFactureTravaux,omitempty"`
	TypeTva                    TypeTva     `json:"typeTva,omitempty"`
}

type Destinataire struct {
	AdresseCodePays         string `json:"adresseDestinataireCodePays,omitempty"`
	AdresseCodePostal       string `json:"adresseDestinataireCodePostal,omitempty"`
	AdresseComplement1      string `json:"adresseDestinataireComplement1,omitempty"`
	AdresseComplement2      string `json:"adresseDestinataireComplement2,omitempty"`
	AdresseDetail           string `json:"adresseDestinataireDetail,omitempty"`
	AdresseId               int    `json:"adresseDestinataireId,omitempty"`
	AdresseLibellePays      string `json:"adresseDestinataireLibellePays,omitempty"`
	AdresseVille            string `json:"adresseDestinataireVille,omitempty"`
	Code                    string `json:"codeDestinataire,omitempty"`
	CodeServiceExecutant    string `json:"codeServiceExecutant,omitempty"`
	Etat                    string `json:"destinataireEtat,omitempty"`
	Id                      int    `json:"idDestinataire,omitempty"`
	IdServiceExecutant      int    `json:"idServiceExecutant,omitempty"`
	Libelle                 string `json:"libelleDestinataire,omitempty"`
	LibelleServiceExecutant string `json:"libelleServiceExecutant,omitempty"`
}

type Fournisseur struct {
	AdresseCodePays              string          `json:"adresseFournisseurCodePays,omitempty"`
	AdresseCodePostal            string          `json:"adresseFournisseurCodePostal,omitempty"`
	AdresseComplement1           string          `json:"adresseFournisseurComplement1,omitempty"`
	AdresseComplement2           string          `json:"adresseFournisseurComplement2,omitempty"`
	AdresseDetail                string          `json:"adresseFournisseurDetail,omitempty"`
	AdresseId                    int64           `json:"adresseFournisseurId,omitempty"`
	AdresseLibellePays           string          `json:"adresseFournisseurLibellePays,omitempty"`
	AdresseVille                 string          `json:"adresseFournisseurVille,omitempty"`
	Affactureur                  Affactureur     `json:"affactureur,omitempty"`
	Code                         string          `json:"codeFournisseur,omitempty"`
	CodeService                  string          `json:"codeServiceFournisseur,omitempty"`
	CoordBancairesBicSwift       string          `json:"coordBancairesBicSwift,omitempty"`
	CoordBancairesCodeGuichet    string          `json:"coordBancairesCodeGuichet,omitempty"`
	CoordBancairesCleIban        string          `json:"coordBancairesFournisseurCleIban,omitempty"`
	CoordBancairesCleRib         string          `json:"coordBancairesFournisseurCleRib,omitempty"`
	CoordBancairesCodeBanque     string          `json:"coordBancairesFournisseurCodeBanque,omitempty"`
	CoordBancairesCodePays       string          `json:"coordBancairesFournisseurCodePays,omitempty"`
	CoordBancairesCompteBancaire string          `json:"coordBancairesFournisseurCompteBancaire,omitempty"`
	CoordBancairesId             int64           `json:"coordBancairesFournisseurId,omitempty"`
	CoordBancairesLibelle        string          `json:"coordBancairesFournisseurLibelle,omitempty"`
	CoordBancairesType           string          `json:"coordBancairesFournisseurType,omitempty"`
	Designation                  string          `json:"designationFournisseur,omitempty"`
	Id                           int64           `json:"idFournisseur,omitempty"`
	IdService                    int64           `json:"idServiceFournisseur,omitempty"`
	LibelleService               string          `json:"libelleServiceFournisseur,omitempty"`
	NumeroRcs                    string          `json:"numeroRcsFournisseur,omitempty"`
	TypeIdentifiant              TypeIdentifiant `json:"typeIdentifiantFournisseur,omitempty"`
}

type ListeLignesDePoste struct {
	Lignes             []LignePoste `json:"lignePoste,omitempty"`
	NbResultatsParPage int64        `json:"nbResultatsParPageLignesPoste,omitempty"`
	PageCourante       int64        `json:"pageCouranteLignesPoste,omitempty"`
	Pages              int64        `json:"pagesLignesPoste,omitempty"`
	Total              int64        `json:"totalLignesPoste,omitempty"`
}

type ListePiecesJointesFacture struct {
	PiecesJointes      []PieceJointeFacture `json:"pieceJointe,omitempty"`
	NbResultatsParPage int64                `json:"nbResultatsParPageListePiecesJointe,omitempty"`
	PageCourante       int64                `json:"pageCouranteListePiecesJointe,omitempty"`
	Pages              int64                `json:"pagesListePiecesJointe,omitempty"`
	Total              int64                `json:"totalListePiecesJointe,omitempty"`
}

type ListeLignesRecapTVA struct {
	Lignes             []LigneTva `json:"ligneTva,omitempty"`
	NbResultatsParPage int64      `json:"nbResultatsParPageLignesRecapitulativesTVA,omitempty"`
	PageCourante       int64      `json:"pageCouranteLignesRecapitulativesTVA,omitempty"`
	Pages              int64      `json:"pagesLignesRecapitulativesTVA,omitempty"`
	Total              int64      `json:"totalLignesRecapitulativesTVA,omitempty"`
}

type Facture struct {
	CadreDeFacturation    CadreDeFacturation         `json:"cadreDeFacturation,omitempty"`
	Commentaire           string                     `json:"commentaire,omitempty"`
	Destinataire          Destinataire               `json:"destinataire,omitempty"`
	Fournisseur           Fournisseur                `json:"fournisseur,omitempty"`
	IdentifiantFactureCPP int                        `json:"identifiantFactureCPP,omitempty"`
	LignesDePoste         *ListeLignesDePoste        `json:"lignesDePoste,omitempty"`
	ListeDesPiecesJointes *ListePiecesJointesFacture `json:"listeDesPiecesJointes,omitempty"`
	ModeDepot             string                     `json:"modeDepot,omitempty"`
	MontantTotal          MontantTotal               `json:"montantTotal,omitempty"`
	NumeroFacture         string                     `json:"numeroFacture,omitempty"`
	PieceJointePrincipale *PieceJointePrincipale     `json:"pieceJointePrincipale,omitempty"`
	PiecePrecedente       *PiecePrecedente           `json:"piecePrecedente,omitempty"`
	PieceSuivante         *PieceSuivante             `json:"pieceSuivante,omitempty"`
	RecapitulatifDeTva    *ListeLignesRecapTVA       `json:"recapitulatifDeTva,omitempty"`
	References            References                 `json:"references,omitempty"`
	StatutFacture         string                     `json:"statutFacture,omitempty"`
}

type ConsulterFactureResponse struct {
	CodeRetour int     `json:"codeRetour,omitempty"`
	Facture    Facture `json:"facture,omitempty"`
	Libelle    string  `json:"libelle,omitempty"`
}

func (s *FacturesService) ConsulterFactureParRecipiendaire(ctx context.Context, opts ConsulterFactureOptions) (*ConsulterFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/consulter/recipiendaire", opts)
	if err != nil {
		return nil, err
	}

	res := new(ConsulterFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *FacturesService) ConsulterFactureParValideur(ctx context.Context, opts ConsulterFactureOptions) (*ConsulterFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/consulter/valideur", opts)
	if err != nil {
		return nil, err
	}

	res := new(ConsulterFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ConsulterFactureFournisseurOptions struct {
	CodeLangue              CodeLangue                       `json:"codeLangue,omitempty"`
	IdUtilisateurCourant    int                              `json:"idUtilisateurCourant,omitempty"`
	IdentifiantFactureCPP   int                              `json:"identifiantFactureCPP,omitempty"`
	PaginationLignesPoste   *PaginationLignesPosteOptions    `json:"pLignesPoste,omitempty"`
	PaginationPiecesJointes *PaginationPiecesJointesOptions  `json:"pListePiecesJointes,omitempty"`
	PaginationLignesTVA     *PaginationLignesRecapTVAOptions `json:"pLigneRecapTva,omitempty"`
}

func (s *FacturesService) ConsulterFactureParFournisseur(ctx context.Context, opts ConsulterFactureFournisseurOptions) (*ConsulterFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/consulter/fournisseur", opts)
	if err != nil {
		return nil, err
	}

	res := new(ConsulterFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ConsulterHistoriqueOptions struct {
	IdEspace                                 int64                                            `json:"idEspace"`
	IdFacture                                int64                                            `json:"idFacture"`
	IdUtilisateurCourant                     int64                                            `json:"idUtilisateurCourant"`
	PaginationHistoActionsUtilisateurs       *PaginationHistoActionsUtilisateursOptions       `json:"paramRechercheHistoActionsUtilisateurs"`
	PaginationHistoEvenementsComplementaires *PaginationHistoEvenementsComplementairesOptions `json:"paramRechercheHistoEvenementsComplementaires"`
	PaginationHistoStatuts                   *PaginationHistoStatutsOptions                   `json:"paramRechercheHistoStatuts"`
}

type DerniereAction struct {
	Code string `json:"derniereActionCode"`
	Id   int64  `json:"derniereActionId"`
}

type HistoAction struct {
	Code                 string    `json:"histoActionCode"`
	DateHeure            *Datetime `json:"histoActionDateHeure"`
	Id                   int32     `json:"histoActionId"`
	UtilisateurEmail     string    `json:"histoActionUtilisateurEmail"`
	UtilisateurNom       string    `json:"histoActionUtilisateurNom"`
	UtilisateurPrenom    string    `json:"histoActionUtilisateurPrenom"`
	UtilisateurTelephone string    `json:"histoActionUtilisateurTelephone"`
}

type HistoriquesDesActionsUtilisateurs struct {
	HistoAction                   []HistoAction `json:"histoAction"`
	NbResultatsParPageHistoAction int32         `json:"nbResultatsParPageHistoAction"`
	PageCouranteHistoAction       int32         `json:"pageCouranteHistoAction"`
	PagesHistoAction              int32         `json:"pagesHistoAction"`
	TotalHistoAction              int32         `json:"totalHistoAction"`
}

type Evenement struct {
	EvenementDateHeure *Datetime `json:"evenementDateHeure"`
	EvenementId        int64     `json:"evenementId"`
	EvenementLibelle   string    `json:"evenementLibelle"`
	EvenementQui       string    `json:"evenementQui"`
}

type HistoriquesDesEvenementsComplementaires struct {
	Evenement                   []Evenement `json:"evenement"`
	NbResultatsParPageEvenement int32       `json:"nbResultatsParPageEvenement"`
	PageCouranteEvenement       int32       `json:"pageCouranteEvenement"`
	PagesEvenement              int32       `json:"pagesEvenement"`
	TotalEvenement              int32       `json:"totalEvenement"`
}

type HistoStatut struct {
	Code                 string    `json:"histoStatutCode"`
	Commentaire          string    `json:"histoStatutCommentaire"`
	DatePassage          *Datetime `json:"histoStatutDatePassage"`
	Id                   int32     `json:"histoStatutId"`
	UtilisateurEmail     string    `json:"histoStatutUtilisateurEmail"`
	UtilisateurNom       string    `json:"histoStatutUtilisateurNom"`
	UtilisateurPrenom    string    `json:"histoStatutUtilisateurPrenom"`
	UtilisateurTelephone string    `json:"histoStatutUtilisateurTelephone"`
}

type HistoriquesDesStatuts struct {
	HistoStatut                   []HistoStatut `json:"histoStatut"`
	NbResultatsParPageHistoStatut int32         `json:"nbResultatsParPageHistoStatut"`
	PageCouranteHistoStatut       int32         `json:"pageCouranteHistoStatut"`
	PagesHistoStatut              int32         `json:"pagesHistoStatut"`
	TotalHistoStatut              int32         `json:"totalHistoStatut"`
}

type ConsulterHistoriqueResponse struct {
	CodeRetour                              int32                                    `json:"codeRetour"`
	DerniereAction                          *DerniereAction                          `json:"derniereAction"`
	HistoriquesDesActionsUtilisateurs       *HistoriquesDesActionsUtilisateurs       `json:"historiquesDesActionsUtilisateurs"`
	HistoriquesDesEvenementsComplementaires *HistoriquesDesEvenementsComplementaires `json:"historiquesDesEvenementsComplementaires"`
	HistoriquesDesStatuts                   *HistoriquesDesStatuts                   `json:"historiquesDesStatuts"`
	IdFacture                               int64                                    `json:"idFacture"`
	Libelle                                 string                                   `json:"libelle"`
	ModeDepot                               string                                   `json:"modeDepot"`
	NumeroFacture                           string                                   `json:"numeroFacture"`
	StatutCourantCode                       string                                   `json:"statutCourantCode"`
}

func (s *FacturesService) ConsulterHistorique(ctx context.Context, opts ConsulterHistoriqueOptions) (*ConsulterHistoriqueResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/consulter/historique", opts)
	if err != nil {
		return nil, err
	}

	res := new(ConsulterHistoriqueResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
