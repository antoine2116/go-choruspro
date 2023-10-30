package choruspro

type TriSens string

const (
	TriSensAsc  TriSens = "Ascendant "
	TriSensDesc TriSens = "Descendant"
)

type PaginationResponse struct {
	NbResultatsParPage int `json:"nbResultatsParPage"`
	PageCourante       int `json:"pageCourante"`
	Pages              int `json:"pages"`
	Total              int `json:"total"`
}

type PaginationOptions struct {
	NbResultatsParPage   int     `json:"nbResultatsParPage,omitempty"`
	PageResultatDemandee int     `json:"pageResultatDemandee,omitempty"`
	TriColonne           string  `json:"triColonne,omitempty"`
	TriSens              TriSens `json:"triSense,omitempty"`
}

type PaginationLignesPosteOptions struct {
	NbResultatsParPage   int     `json:"nbResultatsParPageLignesPoste,omitempty"`
	PageResultatDemandee int     `json:"pageResultatDemandeeLignesPoste,omitempty"`
	TriColonne           string  `json:"triColonneListeLignesPoste,omitempty"`
	TriSens              TriSens `json:"triSensListeLignesPoste,omitempty"`
}

type PaginationPiecesJointesOptions struct {
	NbResultatsParPage   int     `json:"nbResultatsParPageListePieceJointe,omitempty"`
	PageResultatDemandee int     `json:"pageResultatDemandeeListePieceJointe,omitempty"`
	TriColonne           string  `json:"triColonneListePiecesJointes,omitempty"`
	TriSens              TriSens `json:"triSensListePiecesJointes,omitempty"`
}

type PaginationLignesRecapTVAOptions struct {
	NbResultatsParPage   int     `json:"nbResultatsParPageListeListeRecapitulatifTVA,omitempty"`
	PageResultatDemandee int     `json:"pageResultatDemandeeListeRecapitulatifTVA,omitempty"`
	TriColonne           string  `json:"triColonneListeRecapitulatifTVA,omitempty"`
	TriSens              TriSens `json:"triSensListeRecapitulatifTVA,omitempty"`
}
