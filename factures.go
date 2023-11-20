package choruspro

// FacturesService gère les appels à l'API Facture
// API docs : https://developer.aife.economie.gouv.fr/api-catalog-sandbox?filter=Factures
type FacturesService service

// Action Facture est le type des actions possibles sur une facture
type ActionFacture string

const (
	ActionFactureRefus           ActionFacture = "REFUS"
	ActionFactureValidation      ActionFacture = "VALIDATION"
	ActionFactureMiseEnPaiement  ActionFacture = "MISE_EN_PAIEMENT"
	ActionFactureMiseEnRecyclage ActionFacture = "MISE_EN_RECYCLAGE"
	ActionFactureSuspension      ActionFacture = "SUSPENSION"
	ActionFactureRejet           ActionFacture = "REJET"
)

// CadreFac est le type des cadres de facturation
type CadreFac string

const (
	CodeFacA1  CadreFac = "A1_FACTURE_FOURNISSEUR"
	CodeFacA2  CadreFac = "A2_FACTURE_FOURNISSEUR_DEJA_PAYEE"
	CodeFacA3  CadreFac = "A3_MEMOIRE_JUSTICE"
	CodeFacA4  CadreFac = "A4_PROJET_DECOMPTE_MENSUEL_FOURNISSEUR"
	CodeFacA5  CadreFac = "A5_ETAT_ACOMPTE_FOURNISSEUR"
	CodeFacA6  CadreFac = "A6_ETAT_ACOMPTE_VALIDE_FOURNISSEUR"
	CodeFacA7  CadreFac = "A7_PROJET_DECOMPTE_FINAL_FOURNISSEUR"
	CodeFacA8  CadreFac = "A8_DECOMPTE_GENERAL_DEFINITIF_FOURNISSEUR"
	CodeFacA9  CadreFac = "A9_FACTURE_SOUSTRAITANT"
	CodeFacA10 CadreFac = "A10_PROJET_DECOMPTE_MENSUEL_SOUSTRAITANT"
	CodeFacA11 CadreFac = "A11_PROJET_DECOMPTE_FINAL_SOUSTRAITANT"
	CodeFacA12 CadreFac = "A12_FACTURE_COTRAITANT"
	CodeFacA13 CadreFac = "A13_PROJET_DECOMPTE_MENSUEL_COTRAITANT"
	CodeFacA14 CadreFac = "A14_PROJET_DECOMPTE_FINAL_COTRAITANT"
	CodeFacA15 CadreFac = "A15_ETAT_ACOMPTE_MOE"
	CodeFacA16 CadreFac = "A16_ETAT_ACOMPTE_VALIDE_MOE"
	CodeFacA17 CadreFac = "A17_PROJET_DECOMPTE_GENERAL_MOE"
	CodeFacA18 CadreFac = "A18_DECOMPTE_GENERAL_MOE"
	CodeFacA19 CadreFac = "A19_ETAT_ACOMPTE_VALIDE_MOA"
	CodeFacA20 CadreFac = "A20_DECOMPTE_GENERAL_MOA"
	CodeFacA21 CadreFac = "A21_DEMANDE_REMBOURSEMENT_TIC"
	CodeFacA22 CadreFac = "A22_PROJET_DECOMPTE_GENERAL_FOURNISSEUR_PROCEDURE_TACITE"
	CodeFacA23 CadreFac = "A23_DECOMPTE_GENERAL_DEFINITIF_TACITE_FOURNISSEUR"
	CodeFacA24 CadreFac = "A24_DECOMPTE_GENERAL_DEFINITIF_TACITE_MOE"
	CodeFacA25 CadreFac = "A25_DECOMPTE_GENERAL_DEFINITIF_MOE_PROCEDURE_TACITE"
)

// RoleUtilisateur est le type des rôles possibles pour un utilisateur
type RoleUtilisateur string

const (
	RoleFournisseur  RoleUtilisateur = "FOURNISSEUR"
	RoleMoe          RoleUtilisateur = "MOE"
	RoleMoa          RoleUtilisateur = "MOA"
	RoleValidateur   RoleUtilisateur = "VALIDATEUR"
	RoleDestinataire RoleUtilisateur = "DESTINATAIRE"
)

// TypeFacture est le type des factures
type TypeFacture string

const (
	TypeFactureFacture TypeFacture = "FACTURE"
	TypeFactureAvoir   TypeFacture = "AVOIR"
	TypeFactureAcompte TypeFacture = "ACOMPTE"
)

// TypeTva est le type des TVA
type TypeTva string

const (
	TypeTvaSurDebit        TypeTva = "TVA_SUR_DEBIT"
	TypeTvaSurEncaissement TypeTva = "TVA_SUR_ENCAISSEMENT"
	TypeTvaExoneration     TypeTva = "EXONERATION"
	TypeTvaSansTva         TypeTva = "SANS_TVA"
)

// TypeIdentifiant est le type des identifiants pour identifier un tiers
type TypeIdentifiant string

const (
	TypeIdentifiantSiret        TypeIdentifiant = "SIRET"
	TypeIdentifiantUEHorsFrance TypeIdentifiant = "UE_HORS_FRANCE"
	TypeIdentifiantHorsUE       TypeIdentifiant = "HORS_UE"
	TypeIdentifiantRidet        TypeIdentifiant = "RIDET"
	TypeIdentifiantTahiti       TypeIdentifiant = "TAHITI"
	TypeIdentifiantAutre        TypeIdentifiant = "AUTRE"
	TypeIdentifiantParticulier  TypeIdentifiant = "PARTICULIER"
)
