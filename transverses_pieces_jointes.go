package choruspro

import (
	"context"
	"errors"
	"fmt"
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

type ObjectPJ string

const (
	ObjetPJDemandePaiement             ObjectPJ = "DEMANDE_PAIEMENT"
	ObjetPJEngagementJuridique         ObjectPJ = "ENGAGEMENT_JURIDIQUE"
	ObjetPJSolicitation                ObjectPJ = "SOLLICITATION"
	ObjetPJStructurePieceJointe        ObjectPJ = "STRUCTURE_PIECEJOINTE"
	ObjetPJStructureMandat             ObjectPJ = "STRUCTURE_MANDAT"
	ObjetPJStructureCoordonneeBancaire ObjectPJ = "STRUCTURE_COORDONNEE_BANCAIRE"
	ObjetPJUtilisateur                 ObjectPJ = "UTILISATEUR"
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
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/typespj", opts)
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

type ListePiecesJointesStructure struct {
	CodeRetour    int32                  `json:"codeRetour"`
	Libelle       string                 `json:"libelle"`
	PiecesJointes []PieceJointeStructure `json:"listePiecesJointes,omitempty"`
	Pagiantion    *PaginationResponse    `json:"parametresRetour,omitempty"`
}

type PieceJointeStructure struct {
	Designation     string `json:"pieceJointeDesignation"`
	Extension       string `json:"pieceJointeExtension"`
	Id              int64  `json:"pieceJointeId"`
	ObjetId         int64  `json:"pieceJointeObjetId"`
	StructureCode   string `json:"pieceJointeStructureCode"`
	StructureId     int64  `json:"pieceJointeStructureId"`
	StructureRaison string `json:"pieceJointeStructureRaisonSociale"`
	TypeCode        string `json:"pieceJointeTypeCode"`
	TypeLibelle     string `json:"pieceJointeTypeLibelle"`
	IdUtitilisateur int64  `json:"pieceJointeUtilisateurId"`
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

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/pj/structure", opts)
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

type ListePiecesJointesMonCompte struct {
	CodeRetour    int32                  `json:"codeRetour"`
	Libelle       string                 `json:"libelle"`
	PiecesJointes []PieceJointeMonCompte `json:"listePiecesJointes,omitempty"`
	Pagiantion    *PaginationResponse    `json:"parametresRetour,omitempty"`
}

type PieceJointeMonCompte struct {
	Designation   string `json:"pieceJointeDesignation"`
	Extension     string `json:"pieceJointeExtension"`
	Id            int64  `json:"pieceJointeId"`
	ObjetId       int64  `json:"pieceJointeObjetId"`
	TypeCode      string `json:"pieceJointeTypeCode"`
	TypeLibelle   string `json:"pieceJointeTypeLibelle"`
	IdUtilisateur int64  `json:"pieceJointeUtilisateurId"`
}

type ListePiecesJointesMonCompteOptions struct {
	CodeLangue             CodeLangue         `json:"codeLangue,omitempty"`
	PieceJointeDesignation string             `json:"pieceJointeDesignation,omitempty"`
	PieceJointeTypeId      int64              `json:"pieceJointeTypeId"`
	Pagination             *PaginationOptions `json:"parametresRecherche,omitempty"`
}

func (o ListePiecesJointesMonCompteOptions) Validate() error {
	if o.PieceJointeTypeId == 0 {
		return errors.New("choruspro: PieceJointeTypeId is required")
	}

	return nil
}

func (s *TransversesService) RechercherPiecesJointesMonCompte(ctx context.Context, opts ListePiecesJointesMonCompteOptions) (*ListePiecesJointesMonCompte, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/pj/moncompte", opts)
	if err != nil {
		return nil, err
	}

	pieces := new(ListePiecesJointesMonCompte)

	err = s.client.doRequest(ctx, req, pieces)
	if err != nil {
		return nil, err
	}

	return pieces, nil
}

type TelechargerPieceJointeOptions struct {
	Objet ObjectPJ `json:"objetPieceJointe"`

	DemandePaiement *struct {
		// 2 values possible : "OUI" or "NON"
		AvecPJCompltementaires string `json:"avecPiecesJointesComplementaires"`

		// 2 values possible : "PDF" or "PIVOT"
		Format       string `json:"format"`
		ListeFacture []struct {
			IdFacture int64 `json:"idFacture"`
		} `json:"listeFacture"`
	} `json:"demandePaiement"`

	EngagementJuridique *struct {
		Numero string `json:"numeroEngagementJuridique"`
	} `json:"engagementJuridique"`

	Sollicitation *struct {
		Id int64 `json:"idSollicitation"`

		// 2 values possible : "OUI" or "NON"
		PieceJointeSollicitation string `json:"pieceJointeSollicitation"`
	} `json:"sollicitation"`

	Structure *struct {
		Id              int64  `json:"idStructure"`
		Identifiant     string `json:"identifiantStructure"`
		TypeIdentifiant string `json:"typeIdentifiantStructure"`
	} `json:"structure"`
}

type PieceJointe struct {
	PieceJointe string `json:"pieceJointe"` // base64 encoded
}

func (o TelechargerPieceJointeOptions) Validate() error {
	if o.Objet == "" {
		return errors.New("choruspro: ObjetPieceJointe is required")
	}

	// Demande paiement validation
	if o.Objet == ObjetPJDemandePaiement {
		if o.DemandePaiement == nil {
			return fmt.Errorf("choruspro: DemandePaiement is required because ObjetPieceJointe is set to %s", ObjetPJDemandePaiement)
		}

		if o.DemandePaiement.AvecPJCompltementaires != "OUI" && o.DemandePaiement.AvecPJCompltementaires != "NON" {
			return errors.New("choruspro: DemandePaiement.AvecPiecesJointesComplementaires must be set to OUI or NON")
		}

		if o.DemandePaiement.Format != "PDF" && o.DemandePaiement.Format != "PIVOT" {
			return errors.New("choruspro: DemandePaiement.Format must be set to PDF or PIVOT")
		}

		if len(o.DemandePaiement.ListeFacture) == 0 {
			return errors.New("choruspro: DemandePaiement.ListeFacture must contain at least one element")
		}
	}

	// Engagement juridique validation
	if o.Objet == ObjetPJEngagementJuridique {
		if o.EngagementJuridique == nil {
			return fmt.Errorf("choruspro: EngagementJuridique is required because ObjetPieceJointe is set to %s", ObjetPJEngagementJuridique)
		}
	}

	// Sollicitation validation
	if o.Objet == ObjetPJSolicitation {
		if o.Sollicitation == nil {
			return fmt.Errorf("choruspro: Sollicitation is required because ObjetPieceJointe is set to %s", ObjetPJSolicitation)
		}

		if o.Sollicitation.Id == 0 {
			return errors.New("choruspro: Sollicitation.IdSollicitation is required")
		}

		if o.Sollicitation.PieceJointeSollicitation != "OUI" && o.Sollicitation.PieceJointeSollicitation != "NON" {
			return errors.New("choruspro: Sollicitation.PieceJointeSollicitation must be set to OUI or NON")
		}
	}

	// Structure validation
	if o.Objet == ObjetPJStructurePieceJointe ||
		o.Objet == ObjetPJStructureMandat ||
		o.Objet == ObjetPJStructureCoordonneeBancaire {
		if o.Structure == nil {
			return fmt.Errorf("choruspro: Structure is required because ObjetPieceJointe is set to %s", o.Objet)
		}

		if o.Structure.Id == 0 && (o.Structure.Identifiant == "" && o.Structure.TypeIdentifiant == "") {
			return errors.New("choruspro: Structure.IdStructure or Structure.IdentifiantStructure + Structure.TypeIdentifiantStructure are required")
		}

		if o.Structure.Id != 0 && (o.Structure.Identifiant != "" || o.Structure.TypeIdentifiant != "") {
			return errors.New("choruspro: Structure.IdStructure and Structure.IdentifiantStructure + Structure.TypeIdentifiantStructure are mutually exclusive")
		}

		if o.Structure.Identifiant != "" && o.Structure.TypeIdentifiant == "" {
			return errors.New("choruspro: Structure.TypeIdentifiantStructure is required when Structure.IdentifiantStructure is set")
		}

		if o.Structure.Identifiant == "" && o.Structure.TypeIdentifiant != "" {
			return errors.New("choruspro: Structure.IdentifiantStructure is required when Structure.TypeIdentifiantStructure is set")
		}
	}

	return nil
}

func (s *TransversesService) TelechargerPieceJointe(ctx context.Context, opts TelechargerPieceJointeOptions) (*PieceJointe, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

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
