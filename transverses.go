package choruspro

type TransversesService service

type CodeLangue string

const (
	CodeLangueFr CodeLangue = "FR"
	CodeLangueEn CodeLangue = "EN"
)

type SyntaxeFlux string

const (
	SyntaxeFluxE1UblInvoice    SyntaxeFlux = "IN_DP_E1_UBL_INVOICE"
	SyntaxeFluxE1Cii16B        SyntaxeFlux = "IN_DP_E1_CII_16B"
	SyntaxeFluxE2UblInvoiceMin SyntaxeFlux = "IN_DP_E2_UBL_INVOICE_MIN"
	SyntaxeFluxE2CppFactureMin SyntaxeFlux = "IN_DP_E2_CPP_FACTURE_MIN"
	SyntaxeFluxE2CiiMin16B     SyntaxeFlux = "IN_DP_E2_CII_MIN_16B"
	SyntaxeFluxE2CiiFacturx    SyntaxeFlux = "IN_DP_E2_CII_FACTURX"
	SyntaxeFluxE3UblInvoice    SyntaxeFlux = "IN_DP_E3_UBL_INVOICE"
)
