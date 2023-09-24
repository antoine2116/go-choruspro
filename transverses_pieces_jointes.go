package choruspro

import (
	"context"
	"errors"
	"net/http"
)

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

type ListeTypesPieceJointe struct {
	CodeRetour int32             `json:"codeRetour"`
	Libelle    string            `json:"libelle"`
	Types      []TypePieceJointe `json:"listeTypePieceJointe"`
}

type TypePieceJointe struct {
	Code      string    `json:"codeTypePieceJointe"`
	Libelle   string    `json:"libelleTypePieceJointe"`
	Id        int64     `json:"idTypePieceJointe"`
	TypeObjet TypeObjet `json:"typeObjet"`
}

type ListeTypesPieceJointeOptions struct {
	CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	TypeObjet  TypeObjet  `json:"typeObjet,omitempty"`
}

func (s *TransversesService) RecupererTypesPieceJointe(ctx context.Context, opts ListeTypesPieceJointeOptions) (*ListeTypesPieceJointe, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/typespj", opts)
	if err != nil {
		return nil, err
	}

	types := new(ListeTypesPieceJointe)

	err = s.client.doRequest(ctx, req, types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

type ListePiecesJointesStructure struct {
	CodeRetour    int32               `json:"codeRetour"`
	Libelle       string              `json:"libelle"`
	PiecesJointes []PieceJointe       `json:"listePiecesJointes,omitempty"`
	Pagiantion    *PaginationResponse `json:"parametresRetour,omitempty"`
}

type PieceJointe struct {
	Designation     string `json:"pieceJointeDesignation"`
	Extension       string `json:"pieceJointeExtension"`
	Id              int64  `json:"pieceJointeId"`
	ObjetId         int64  `json:"pieceJointeObjetId"`
	StructureCode   string `json:"pieceJointeStructureCode"`
	StructureId     int64  `json:"pieceJointeStructureId"`
	StructureRaison string `json:"pieceJointeStructureRaisonSociale"`
	TypeCode        string `json:"pieceJointeTypeCode"`
	TypeLibelle     string `json:"pieceJointeTypeLibelle"`
}

type ListePiecesJointesStructureOptions struct {
	CodeLangue             CodeLangue         `json:"codeLangue,omitempty"`
	IdStructureCPP         int64              `json:"idStructureCPP"`
	PieceJointeDesignation string             `json:"pieceJointeDesignation,omitempty"`
	PieceJointeTypeId      int64              `json:"pieceJointeTypeId,omitempty"`
	Pagination             *PaginationOptions `json:"parametresRecherche,omitempty"`
}

func (o ListePiecesJointesStructureOptions) Validate() error {
	if o.IdStructureCPP == 0 {
		return errors.New("choruspro: IdStructureCPP is required")
	}

	return nil
}

func (s *TransversesService) RechercherPiecesJointesStructure(ctx context.Context, opts ListePiecesJointesStructureOptions) (*ListePiecesJointesStructure, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/rechercher/pj/structure", opts)
	if err != nil {
		return nil, err
	}

	pieces := new(ListePiecesJointesStructure)

	err = s.client.doRequest(ctx, req, pieces)
	if err != nil {
		return nil, err
	}

	return pieces, nil
}

type AjouterPieceResponse struct {
	CodeRetour    int32  `json:"codeRetour"`
	Libelle       string `json:"libelle"`
	PieceJointeId int64  `json:"pieceJointeId"`
}

type AjouterPieceOptions struct {
	IdUtilisateurCourant int64  `json:"idUtilisateurCourant,omitempty"`
	Extension            string `json:"pieceJointeExtension"`
	Fichier              string `json:"pieceJointeFichier"`
	Nom                  string `json:"pieceJointeNom"`
	TypeMime             string `json:"pieceJointeTypeMime"`
}

func (o AjouterPieceOptions) Validate() error {
	if o.Extension == "" {
		return errors.New("choruspro: Extension is required")
	}

	if o.Fichier == "" {
		return errors.New("choruspro: Fichier is required")
	}

	if o.Nom == "" {
		return errors.New("choruspro: Nom is required")
	}

	if o.TypeMime == "" {
		return errors.New("choruspro: TypeMime is required")
	}

	return nil
}

func (s *TransversesService) AjouterPieceJointe(ctx context.Context, opts AjouterPieceOptions) (*AjouterPieceResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/ajouter/fichier", opts)
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
