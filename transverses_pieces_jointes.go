package choruspro

import (
	"context"
	"net/http"
)

// TypeObjet est le type de l'objet d'une pièce jointe.
type TypeObjet string

const (
	TypeObjetStructure         TypeObjet = "STRUCTURE"
	TypeObjetUtilisateur       TypeObjet = "UTILISATEUR"
	TypeObjetCb                TypeObjet = "CB"
	TypeObjetMandat            TypeObjet = "MANDAT"
	TypeObjetFicheRaccordement TypeObjet = "FICHE_RACCORDEMENT"
	TypeObjetEj                TypeObjet = "EJ"
	TypeObjetMemoire           TypeObjet = "MEMOIRE"
	TypeObjetDmdRbst           TypeObjet = "DMD_RBST"
	TypeObjetTransverse        TypeObjet = "TRANSVERSE"
	TypeObjetFactureTravaux    TypeObjet = "FACTURE_TRAVAUX"
	TypeObjetFacture           TypeObjet = "FACTURE"
	TypeObjetDmdeCpteDevSiExt  TypeObjet = "DMDE_CPTE_DEV_SI_EXT"
	TypeObjetSolicitation      TypeObjet = "SOLLICITATION"
	TypeObjetCertificat        TypeObjet = "CERTIFICAT"
	TypeObjetPjRejetEdi        TypeObjet = "PJ_REJET_EDI"
)

// ObjetPJ est l'ojet d'une pièce jointe.
type ObjetPJ string

const (
	ObjetPJDemandePaiement             ObjetPJ = "DEMANDE_PAIEMENT"
	ObjetPJEngagementJuridique         ObjetPJ = "ENGAGEMENT_JURIDIQUE"
	ObjetPJSollicitation               ObjetPJ = "SOLLICITATION"
	ObjetPJStructurePieceJointe        ObjetPJ = "STRUCTURE_PIECEJOINTE"
	ObjetPJStructureMandat             ObjetPJ = "STRUCTURE_MANDAT"
	ObjetPJStructureCoordonneeBancaire ObjetPJ = "STRUCTURE_COORDONNEE_BANCAIRE"
	ObjetPJUtilisateur                 ObjetPJ = "UTILISATEUR"
)

// ListeTypesPieceJointeResponse est la structure de données représentant
// la réponse de la méthode RecupererTypesPieceJointe.
type ListeTypesPieceJointeResponse struct {
	CodeRetour int32             `json:"codeRetour"`
	Libelle    string            `json:"libelle"`
	Types      []TypePieceJointe `json:"listeTypePieceJointe"`
}

// TypePieceJointe est le type d'une pièce jointe.
type TypePieceJointe struct {
	Code      string    `json:"codeTypePieceJointe"`
	Libelle   string    `json:"libelleTypePieceJointe"`
	Id        int64     `json:"idTypePieceJointe"`
	TypeObjet TypeObjet `json:"typeObjet"`
}

// ListeTypesPieceJointeOptions est la structure de données représentant
// les options de la méthode RecupererTypesPieceJointe.
type ListeTypesPieceJointeOptions struct {
	CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	TypeObjet  TypeObjet  `json:"typeObjet,omitempty"`
}

// La méthode RecupererTypesPieceJointe permet de récupérer le type de pièce-jointe
// pouvant être déposé en fonction de l'espace sur lequel l'objet peut être rattaché.
func (s *TransversesService) RecupererTypesPieceJointe(ctx context.Context, opts ListeTypesPieceJointeOptions) (*ListeTypesPieceJointeResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/typespj", opts)
	if err != nil {
		return nil, err
	}

	types := new(ListeTypesPieceJointeResponse)

	err = s.client.doRequest(ctx, req, types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// AjouterPieceResponse est la structure de données représentant
// la réponse de la méthode AjouterPieceJointe.
type AjouterPieceResponse struct {
	CodeRetour    int32  `json:"codeRetour"`
	Libelle       string `json:"libelle"`
	IdPieceJointe int64  `json:"pieceJointeId"`
}

// AjouterPieceOptions est la structure de données utilisée pour
// appeler la méthode AjouterPieceJointe.
type AjouterPieceOptions struct {
	IdUtilisateurCourant int64  `json:"idUtilisateurCourant,omitempty"`
	Extension            string `json:"pieceJointeExtension"`
	Fichier              string `json:"pieceJointeFichier"`
	Nom                  string `json:"pieceJointeNom"`
	TypeMime             string `json:"pieceJointeTypeMime"`
}

// La méthode AjouterPieceJointe permet d'ajouter une pièce-jointe au
// compte utilisateur courant et d'en obtenir l'identifiant technique.
// La pièce jointe ne doit pas dépasser 10Mo. Si le fichier dépasse cette
// taille, une erreur 20003 sera remontée.
func (s *TransversesService) AjouterPieceJointe(ctx context.Context, opts AjouterPieceOptions) (*AjouterPieceResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/ajouter/fichier", opts)
	if err != nil {
		return nil, err
	}

	piece := new(AjouterPieceResponse)

	err = s.client.doRequest(ctx, req, piece)
	if err != nil {
		return nil, err
	}

	return piece, nil
}

// SupprimerPieceOptions est la structure de données représentant
// la réponse de la méthode RechercherPiecesJointesStructure.
type ListePiecesJointesStructureResponse struct {
	CodeRetour    int32                  `json:"codeRetour"`
	Libelle       string                 `json:"libelle"`
	PiecesJointes []PieceJointeStructure `json:"listePiecesJointes,omitempty"`
	Pagination    *PaginationResponse    `json:"parametresRetour,omitempty"`
}

// PieceJointeStructure est la structure de données représentant
// une pièce jointe rattachée à une structure.
type PieceJointeStructure struct {
	Designation            string `json:"pieceJointeDesignation"`
	Extension              string `json:"pieceJointeExtension"`
	Id                     int64  `json:"pieceJointeId"`
	ObjetId                int64  `json:"pieceJointeObjetId"`
	StructureCode          string `json:"pieceJointeStructureCode"`
	StructureId            int64  `json:"pieceJointeStructureId"`
	StructureRaisonSociale string `json:"pieceJointeStructureRaisonSociale"`
	TypeCode               string `json:"pieceJointeTypeCode"`
	TypeLibelle            string `json:"pieceJointeTypeLibelle"`
	IdUtilisateur          int64  `json:"pieceJointeUtilisateurId"`
}

// ListePiecesJointesStructureOptions est la structure de données représentant
// les options de la méthode RechercherPiecesJointesStructure.
type ListePiecesJointesStructureOptions struct {
	CodeLangue             CodeLangue         `json:"codeLangue,omitempty"`
	IdStructureCPP         int64              `json:"idStructureCPP"`
	PieceJointeDesignation string             `json:"pieceJointeDesignation,omitempty"`
	PieceJointeTypeId      int64              `json:"pieceJointeTypeId,omitempty"`
	Pagination             *PaginationOptions `json:"parametresRecherche,omitempty"`
}

// La méthode RechercherPiecesJointesStructure permet de rechercher une pièce-jointe
// associée à une structure à laquelle est habilité l'utilisateur.
func (s *TransversesService) RechercherPiecesJointesStructure(ctx context.Context, opts ListePiecesJointesStructureOptions) (*ListePiecesJointesStructureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/pj/structure", opts)
	if err != nil {
		return nil, err
	}

	pieces := new(ListePiecesJointesStructureResponse)

	err = s.client.doRequest(ctx, req, pieces)
	if err != nil {
		return nil, err
	}

	return pieces, nil
}

// ListePiecesJointesMonCompteResponse est la structure de données représentant
// la réponse de la méthode RechercherPiecesJointesMonCompte.
type ListePiecesJointesMonCompteResponse struct {
	CodeRetour    int32                  `json:"codeRetour"`
	Libelle       string                 `json:"libelle"`
	PiecesJointes []PieceJointeMonCompte `json:"listePiecesJointes,omitempty"`
	Pagination    *PaginationResponse    `json:"parametresRetour,omitempty"`
}

// PieceJointeMonCompte est la structure de données représentant
// une pièce jointe rattachée à un compte utilisateur.
type PieceJointeMonCompte struct {
	Designation   string `json:"pieceJointeDesignation"`
	Extension     string `json:"pieceJointeExtension"`
	Id            int64  `json:"pieceJointeId"`
	ObjetId       int64  `json:"pieceJointeObjetId"`
	TypeCode      string `json:"pieceJointeTypeCode"`
	TypeLibelle   string `json:"pieceJointeTypeLibelle"`
	IdUtilisateur int64  `json:"pieceJointeUtilisateurId"`
}

// ListePiecesJointesMonCompteOptions est la structure de données utilisée
// pour appeler la méthode RechercherPiecesJointesMonCompte.
type ListePiecesJointesMonCompteOptions struct {
	CodeLangue             CodeLangue         `json:"codeLangue,omitempty"`
	PieceJointeDesignation string             `json:"pieceJointeDesignation,omitempty"`
	PieceJointeTypeId      int64              `json:"pieceJointeTypeId"`
	Pagination             *PaginationOptions `json:"parametresRecherche,omitempty"`
}

// La méthode RechercherPiecesJointesMonCompte permet de rechercher
// un fichier associé au compte utilisateur.
func (s *TransversesService) RechercherPiecesJointesMonCompte(ctx context.Context, opts ListePiecesJointesMonCompteOptions) (*ListePiecesJointesMonCompteResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/pj/moncompte", opts)
	if err != nil {
		return nil, err
	}

	pieces := new(ListePiecesJointesMonCompteResponse)

	err = s.client.doRequest(ctx, req, pieces)
	if err != nil {
		return nil, err
	}

	return pieces, nil
}

// TelechargerPieceJointeDemandePaiementOptions est la structure de données utilisée
// pour appeler la méthode TelechargerPieceJointeDemandePaiement.
type TelechargerPieceJointeDemandePaiementOptions struct {
	// 2 values possible : "OUI" or "NON"
	AvecPJCompltementaires string        `json:"avecPiecesJointesComplementaires"`
	Format                 FormatFichier `json:"format"`
	ListeFacture           []IdFacture   `json:"listeFacture"`
}

// IdFacture est la structure de données représentant
// l'identifiant d'une facture.
type IdFacture struct {
	IdFacture int64 `json:"idFacture"`
}

// TelechargerPieceJointeSollicitationOptions est la structure de données utilisée
// pour appeler la méthode TelechargerPieceJointe.
// Elle est utilisée dans la structure TelechargerPieceJointeOptions.
type TelechargerPieceJointeSollicitationOptions struct {
	Id int64 `json:"idSollicitation"`

	// 2 values possible : "OUI" or "NON"
	PieceJointeSollicitation string `json:"pieceJointeSollicitation"`
}

// TelechargerPieceJointeEngagementJuridiqueOptions est la structure de données utilisée
// pour appeler la méthode TelechargerPieceJointe.
// Elle est utilisée dans la structure TelechargerPieceJointeOptions.
type TelechargerPieceJointeEngagementJuridiqueOptions struct {
	Numero string `json:"numeroEngagementJuridique"`
}

// TelechargerPieceJointeStructureOptions est la structure de données utilisée
// pour appeler la méthode TelechargerPieceJointe.
// Elle est utilisée dans la structure TelechargerPieceJointeOptions.
type TelechargerPieceJointeStructureOptions struct {
	Id              int64  `json:"idStructure"`
	Identifiant     string `json:"identifiantStructure"`
	TypeIdentifiant string `json:"typeIdentifiantStructure"`
}

// TelechargerPieceJointeOptions est la structure de données utilisée
// pour appeler la méthode TelechargerPieceJointe.
type TelechargerPieceJointeOptions struct {
	Objet               ObjetPJ                                           `json:"objetPieceJointe"`
	DemandePaiement     *TelechargerPieceJointeDemandePaiementOptions     `json:"demandePaiement,omitempty"`
	EngagementJuridique *TelechargerPieceJointeEngagementJuridiqueOptions `json:"engagementJuridique,omitempty"`
	Sollicitation       *TelechargerPieceJointeSollicitationOptions       `json:"sollicitation,omitempty"`
	Structure           *TelechargerPieceJointeStructureOptions           `json:"structure,omitempty"`
}

// PieceJointe est la structure de données représentant
// une pièce jointe.
type PieceJointe struct {
	PieceJointe string `json:"pieceJointe"` // base64 encoded
}

// La méthode TelechargerPieceJointe permet à un utilisateur de télécharger
// les pieces jointes des types objets suivants : sollicitation, demande de
// remboursement, facture, facture travaux, mémoire, coordonnée bancaire,
// structure, mandat, utilisateur (courant), engagement juridique.
func (s *TransversesService) TelechargerPieceJointe(ctx context.Context, opts TelechargerPieceJointeOptions) (*PieceJointe, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/telecharger/pieceJointe", opts)
	if err != nil {
		return nil, err
	}

	piece := new(PieceJointe)

	err = s.client.doRequest(ctx, req, piece)
	if err != nil {
		return nil, err
	}

	return piece, nil
}
