package choruspro

import (
	"context"
	"net/http"
)

// CompleterFactureOptions est la structure de données utilisée pour
// appeler la méthode CompleterFacture.
type CompleterFactureOptions struct {
	Commentaire                  string                       `json:"commentaire,omitempty"`
	IdUtilisateurCourant         int64                        `json:"idUtilisateurCourant,omitempty"`
	IdentifiantFactureCPP        int64                        `json:"identifiantFactureCPP,omitempty"`
	PiecesJointesComplementaires []*PieceJointeComplementaire `json:"pieceJointeComplementaire,omitempty"`
}

// CompleterFactureResponse est la structure de données représentant
// la réponse de la méthode CompleterFacture.
type CompleterFactureResponse struct {
	CodeRetour            int    `json:"codeRetour,omitempty"`
	DateTraitement        *Date  `json:"dateTraitement,omitempty"`
	IdentifiantFactureCPP int64  `json:"identifiantFactureCPP,omitempty"`
	Libelle               string `json:"libelle,omitempty"`
	NumeroFacture         string `json:"numeroFacture,omitempty"`
}

// La méthode CompleterFacture permet de modifier une facture au statut
// "SUSPENDUE" en ajoutant des pièces jointes et/ou en modifiant le champ
// "commentaire"
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

// RecyclerFactureOptions est la structure de données utilisée pour
// appeler la méthode RecyclerFacture.
type RecyclerFactureOptions struct {
	Commentaire           string                        `json:"commentaire"`
	Destinataire          *SoumettreFactureDestinataire `json:"destinataire"`
	IdUtilisateurCourant  int                           `json:"idUtilisateurCourant"`
	IdentifiantFactureCPP int                           `json:"identifiantFactureCPP"`
	Reference             *SoumettreFactureReferences   `json:"reference"`
}

// RecyclerFactureResponse est la structure de données représentant
// la réponse de la méthode RecyclerFacture.
type RecyclerFactureResponse struct {
	CodeRetour            int32  `json:"codeRetour"`
	DateDepot             *Date  `json:"dateDepot"`
	IdentifiantFactureCPP int32  `json:"identifiantFactureCPP"`
	Libelle               string `json:"libelle"`
	NumeroFacture         string `json:"numeroFacture"`
	StatutFacture         string `json:"statutFacture"`
}

// La méthode RecyclerFacture permet de modifier les données
// d'acheminement d'une facture au statut "A_Recycler"
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

// SoumettreFactureOptions est la structure de données utilisée pour
// appeler la méthode SoumettreFacture.
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

// SoumettreFactureCadreDeFacturation est la structure de données réprésentant
// le cadre de facturation d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type SoumettreFactureCadreDeFacturation struct {
	Code                  CadreFac `json:"codeCadreFacturation,omitempty"`
	CodeServiceValideur   string   `json:"codeServiceValideur,omitempty"`
	CodeStructureValideur string   `json:"codeStructureValideur,omitempty"`
}

// SoumettreFactureDestinataire est la structure de données réprésentant
// le destinataire d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type SoumettreFactureDestinataire struct {
	Code                 string `json:"codeDestinataire,omitempty"`
	CodeServiceExecutant string `json:"codeServiceExecutant,omitempty"`
}

// SoumettreFactureFournisseur est la structure de données réprésentant
// le fournisseur d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type SoumettreFactureFournisseur struct {
	CodeCoordonneesBancairesFournisseur int64 `json:"codeCoordonneesBancairesFournisseur,omitempty"`
	IdFournisseur                       int64 `json:"idFournisseur,omitempty"`
	IdServiceFournisseur                int64 `json:"idServiceFournisseur,omitempty"`
}

// SoumettreFactureLignePoste est la structure de données réprésentant
// une ligne de poste d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
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

// SoumettreFactureLigneTva est la structure de données réprésentant
// une ligne de TVA d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type SoumettreFactureLigneTva struct {
	MontantBaseHtParTaux float32 `json:"ligneTvaMontantBaseHtParTaux,omitempty"`
	MontantTvaParTaux    float32 `json:"ligneTvaMontantTvaParTaux,omitempty"`
	Taux                 string  `json:"ligneTvaTaux,omitempty"`
	TauxManuel           float32 `json:"ligneTvaTauxManuel,omitempty"`
}

// SoumettreFactureMontantTotal est la structure de données réprésentant
// le montant total d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type SoumettreFactureMontantTotal struct {
	MontantAPayer           float32 `json:"montantAPayer,omitempty"`
	MontantHtTotal          float32 `json:"montantHtTotal,omitempty"`
	MontantRemiseGlobaleTTC float32 `json:"montantRemiseGlobaleTTC,omitempty"`
	MontantTVA              float32 `json:"montantTVA,omitempty"`
	MontantTtcTotal         float32 `json:"montantTtcTotal,omitempty"`
	MotifRemiseGlobaleTTC   string  `json:"motifRemiseGlobaleTTC,omitempty"`
}

// PieceJointeComplementaire est la structure de données réprésentant
// une pièce jointe complémentaire d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type PieceJointeComplementaire struct {
	Designation        string `json:"pieceJointeComplementaireDesignation,omitempty"`
	Id                 int64  `json:"pieceJointeComplementaireId,omitempty"`
	IdLiaison          int64  `json:"pieceJointeComplementaireIdLiaison,omitempty"`
	NumeroLigneFacture int64  `json:"pieceJointeComplementaireNumeroLigneFacture,omitempty"`
	Type               string `json:"pieceJointeComplementaireType,omitempty"`
}

// SoumettreFacturePJPrincipale est la structure de données réprésentant
// la pièce jointe principale d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
type SoumettreFacturePJPrincipale struct {
	Designation string `json:"pieceJointePrincipaleDesignation,omitempty"`
	Id          int64  `json:"pieceJointePrincipaleId,omitempty"`
}

// SoumettreFactureReferences est la structure de données réprésentant
// les références d'une facture. Elle est utilisée dans la méthode
// SoumettreFacture.
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

// SoumettreFactureResponse est la structure de données réprésentant
// la réponse de la méthode SoumettreFacture.
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

// La méthode SoumettreFacture permet de soumettre une facture à la
// solution mutualisée-CPP 2017 en renseignant les données nécessaires
// à la constitution d'un flux.
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

// CorrigerValideurFactureOptions est la structure de données utlisée
// pour la méthode CorrigerValideurFacture.
type CorrigerValideurFactureOptions struct {
	IdFacture                int64  `json:"idFacture"`
	IdStructure              int64  `json:"idStructure"`
	IdentifiantStructure     string `json:"identifiantStructure"`
	TypeIdentifiantStructure string `json:"typeIdentifiantStructure"`
}

// CorrigerValideurFactureResponse est la structure de données réprésentant
// la réponse de la méthode CorrigerValideurFacture.
type CorrigerValideurFactureResponse struct {
	CodeRetour int32  `json:"codeRetour"`
	IdFacture  int64  `json:"idFacture"`
	Libelle    string `json:"libelle"`
}

// La méthode CorrigerValideurFacture permet à un fournisseur de corriger
// le valideur initialement renseigné sur une facture rejetée pour
// renseignement d'un mauvais valideur.
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

// DeposerFluxFactureOptions est la structure de données utlisée
// pour la méthode DeposerFluxFacture.
type DeposerFluxFactureOptions struct {
	AvecSignature        bool        `json:"avecSignature,omitempty"`
	FichierFlux          string      `json:"fichierFlux,omitempty"`
	IdUtilisateurCourant int64       `json:"idUtilisateurCourant,omitempty"`
	NomFichier           string      `json:"nomFichier,omitempty"`
	SyntaxeFlux          SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

// DeposerFluxFactureResponse est la structure de données réprésentant
// la réponse de la méthode DeposerFluxFacture.
type DeposerFluxFactureResponse struct {
	CodeRetour      int         `json:"codeRetour,omitempty"`
	DateDepot       *Date       `json:"dateDepot,omitempty"`
	Libelle         string      `json:"libelle,omitempty"`
	NumeroFluxDepot string      `json:"numeroFluxDepot,omitempty"`
	SyntaxeFlux     SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

// La méthode DeposerFluxFacture permet de déposer un fichier XML ou
// PDF/A3 permettant de renseigner les données nécessaires à la
// constitution d'un flux facture.
func (s *FacturesService) DeposerFluxFacture(ctx context.Context, opts DeposerFluxFactureOptions) (*DeposerFluxFactureResponse, error) {
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
