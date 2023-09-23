package choruspro

type PaginationResponse struct {
	NbResultatsParPage int32 `json:"nbResultatsParPage"`
	PageCourante       int32 `json:"pageCourante"`
	Pages              int32 `json:"pages"`
	Total              int32 `json:"total"`
}

type PaginationOptions struct {
	NbResultatsParPage   int32  `json:"nbResultatsParPage,omitempty"`
	PageResultatDemandee int32  `json:"pageResultatDemandee,omitempty"`
	TriColonne           string `json:"triColonne,omitempty"`
	TriSens              string `json:"triSense,omitempty"`
}
