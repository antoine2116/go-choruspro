package choruspro

import (
	"context"
	"net/http"
)

type CompleterFactureOptions struct {
	Commentaire                  string                       `json:"commentaire,omitempty"`
	IdUtilisateurCourant         int64                        `json:"idUtilisateurCourant,omitempty"`
	IdentifiantFactureCPP        int64                        `json:"identifiantFactureCPP,omitempty"`
	PiecesJointesComplementaires []*PieceJointeComplementaire `json:"pieceJointeComplementaire,omitempty"`
}

type CompleterFactureResponse struct {
	CodeRetour            int    `json:"codeRetour,omitempty"`
	DateTraitement        *Date  `json:"dateTraitement,omitempty"`
	IdentifiantFactureCPP int64  `json:"identifiantFactureCPP,omitempty"`
	Libelle               string `json:"libelle,omitempty"`
	NumeroFacture         string `json:"numeroFacture,omitempty"`
}

func (s *FacturesService) CompleterFacture(ctx context.Context, opts CompleterFactureOptions) (*CompleterFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/completer", opts)
	if err != nil {
		return nil, err
	}

	res := new(CompleterFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type RecyclerFactureOptions struct {
	Commentaire           string                        `json:"commentaire"`
	Destinataire          *SoumettreFactureDestinataire `json:"destinataire"`
	IdUtilisateurCourant  int                           `json:"idUtilisateurCourant"`
	IdentifiantFactureCPP int                           `json:"identifiantFactureCPP"`
	Reference             *SoumettreFactureReferences   `json:"reference"`
}

type RecyclerFactureResponse struct {
	CodeRetour            int32  `json:"codeRetour"`
	DateDepot             *Date  `json:"dateDepot"`
	IdentifiantFactureCPP int32  `json:"identifiantFactureCPP"`
	Libelle               string `json:"libelle"`
	NumeroFacture         string `json:"numeroFacture"`
	StatutFacture         string `json:"statutFacture"`
}

func (s *FacturesService) RecyclerFacture(ctx context.Context, opts RecyclerFactureOptions) (*RecyclerFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/recycler", opts)
	if err != nil {
		return nil, err
	}

	res := new(RecyclerFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SoumettreFactureOptions struct {
	CadreDeFacturation           SoumettreFactureCadreDeFacturation `json:"cadreDeFacturation,omitempty"`
	Commentaire                  string                             `json:"commentaire,omitempty"`
	DateFacture                  *Date                              `json:"dateFacture,omitempty"`
	Destinataire                 SoumettreFactureDestinataire       `json:"destinataire,omitempty"`
	Fournisseur                  SoumettreFactureFournisseur        `json:"fournisseur,omitempty"`
	IdUtilisateurCourant         int64                              `json:"idUtilisateurCourant,omitempty"`
	LignesPoste                  *[]SoumettreFactureLignePoste      `json:"lignePoste,omitempty"`
	LignesTva                    *[]SoumettreFactureLigneTva        `json:"ligneTva,omitempty"`
	ModeDepot                    string                             `json:"modeDepot,omitempty"`
	MontantTotal                 SoumettreFactureMontantTotal       `json:"montantTotal,omitempty"`
	NumeroFactureSaisi           string                             `json:"numeroFactureSaisi,omitempty"`
	PiecesJointesComplementaires *[]PieceJointeComplementaire       `json:"pieceJointeComplementaire,omitempty"`
	PiecesJointesPrincipales     SoumettreFacturePJPrincipale       `json:"pieceJointePrincipale,omitempty"`
	References                   SoumettreFactureReferences         `json:"references,omitempty"`
}

type SoumettreFactureCadreDeFacturation struct {
	Code                  CadreFac `json:"codeCadreFacturation,omitempty"`
	CodeServiceValideur   string   `json:"codeServiceValideur,omitempty"`
	CodeStructureValideur string   `json:"codeStructureValideur,omitempty"`
}

type SoumettreFactureDestinataire struct {
	Code                 string `json:"codeDestinataire,omitempty"`
	CodeServiceExecutant string `json:"codeServiceExecutant,omitempty"`
}

type SoumettreFactureFournisseur struct {
	CodeCoordonneesBancairesFournisseur int64 `json:"codeCoordonneesBancairesFournisseur,omitempty"`
	IdFournisseur                       int64 `json:"idFournisseur,omitempty"`
	IdServiceFournisseur                int64 `json:"idServiceFournisseur,omitempty"`
}

type SoumettreFactureLignePoste struct {
	Denomination      string  `json:"lignePosteDenomination,omitempty"`
	MontantRemiseHT   float32 `json:"lignePosteMontantRemiseHT,omitempty"`
	MontantUnitaireHT float32 `json:"lignePosteMontantUnitaireHT,omitempty"`
	Numero            int32   `json:"lignePosteNumero,omitempty"`
	Quantite          float32 `json:"lignePosteQuantite,omitempty"`
	Reference         string  `json:"lignePosteReference,omitempty"`
	TauxTva           string  `json:"lignePosteTauxTva,omitempty"`
	TauxTvaManuel     float32 `json:"lignePosteTauxTvaManuel,omitempty"`
	Unite             string  `json:"lignePosteUnite,omitempty"`
}

type SoumettreFactureLigneTva struct {
	MontantBaseHtParTaux float32 `json:"ligneTvaMontantBaseHtParTaux,omitempty"`
	MontantTvaParTaux    float32 `json:"ligneTvaMontantTvaParTaux,omitempty"`
	Taux                 string  `json:"ligneTvaTaux,omitempty"`
	TauxManuel           float32 `json:"ligneTvaTauxManuel,omitempty"`
}

type SoumettreFactureMontantTotal struct {
	MontantAPayer           float32 `json:"montantAPayer,omitempty"`
	MontantHtTotal          float32 `json:"montantHtTotal,omitempty"`
	MontantRemiseGlobaleTTC float32 `json:"montantRemiseGlobaleTTC,omitempty"`
	MontantTVA              float32 `json:"montantTVA,omitempty"`
	MontantTtcTotal         float32 `json:"montantTtcTotal,omitempty"`
	MotifRemiseGlobaleTTC   string  `json:"motifRemiseGlobaleTTC,omitempty"`
}

type PieceJointeComplementaire struct {
	Designation        string `json:"pieceJointeComplementaireDesignation,omitempty"`
	Id                 int64  `json:"pieceJointeComplementaireId,omitempty"`
	IdLiaison          int64  `json:"pieceJointeComplementaireIdLiaison,omitempty"`
	NumeroLigneFacture int64  `json:"pieceJointeComplementaireNumeroLigneFacture,omitempty"`
	Type               string `json:"pieceJointeComplementaireType,omitempty"`
}

type SoumettreFacturePJPrincipale struct {
	Designation string `json:"pieceJointePrincipaleDesignation,omitempty"`
	Id          int64  `json:"pieceJointePrincipaleId,omitempty"`
}

type SoumettreFactureReferences struct {
	DeviseFacture        string `json:"deviseFacture,omitempty"`
	ModePaiement         string `json:"modePaiement,omitempty"`
	MotifExonerationTva  string `json:"motifExonerationTva,omitempty"`
	NumeroBonCommande    string `json:"numeroBonCommande,omitempty"`
	NumeroFactureOrigine string `json:"numeroFactureOrigine,omitempty"`
	NumeroMarche         string `json:"numeroMarche,omitempty"`
	TypeFacture          string `json:"typeFacture,omitempty"`
	TypeTva              string `json:"typeTva,omitempty"`
}

type SoumettreFactureResponse struct {
	CodeRetour               int    `json:"codeRetour,omitempty"`
	DateDepot                *Date  `json:"dateDepot,omitempty"`
	EmpreinteCertificatDepot string `json:"empreinteCertificatDepot,omitempty"`
	IdentifiantFactureCPP    int64  `json:"identifiantFactureCPP,omitempty"`
	IdentifiantStructure     string `json:"identifiantStructure,omitempty"`
	Libelle                  string `json:"libelle,omitempty"`
	NumeroFacture            string `json:"numeroFacture,omitempty"`
	StatutFacture            string `json:"statutFacture,omitempty"`
}

func (s *FacturesService) SoumettreFacture(ctx context.Context, opts SoumettreFactureOptions) (*SoumettreFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/soumettre", opts)
	if err != nil {
		return nil, err
	}

	res := new(SoumettreFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type DeposerFluxFactureOptions struct {
	AvecSignature        bool        `json:"avecSignature,omitempty"`
	FichierFlux          string      `json:"fichierFlux,omitempty"`
	IdUtilisateurCourant int64       `json:"idUtilisateurCourant,omitempty"`
	NomFichier           string      `json:"nomFichier,omitempty"`
	SyntaxeFlux          SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

type DeposerFluxFactureResponse struct {
	CodeRetour      int         `json:"codeRetour,omitempty"`
	DateDepot       *Date       `json:"dateDepot,omitempty"`
	Libelle         string      `json:"libelle,omitempty"`
	NumeroFluxDepot string      `json:"numeroFluxDepot,omitempty"`
	SyntaxeFlux     SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

func (s *FacturesService) DeposerFlux(ctx context.Context, opts DeposerFluxFactureOptions) (*DeposerFluxFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/deposer/flux", opts)
	if err != nil {
		return nil, err
	}

	res := new(DeposerFluxFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CorrigerValideurFactureOptions struct {
	IdFacture                int64  `json:"idFacture"`
	IdStructure              int64  `json:"idStructure"`
	IdentifiantStructure     string `json:"identifiantStructure"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructure"`
}

type CorrigerValideurFactureResponse struct {
	CodeRetour int32  `json:"codeRetour"`
	IdFacture  int64  `json:"idFacture"`
	Libelle    string `json:"libelle"`
}

func (s *FacturesService) CorrigerValideurFacture(ctx context.Context, opts CorrigerValideurFactureOptions) (*CorrigerValideurFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/corriger/valideur/facture", opts)
	if err != nil {
		return nil, err
	}

	res := new(CorrigerValideurFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
