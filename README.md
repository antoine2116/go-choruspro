# go-choruspro

![tests](https://github.com/antoine2116/go-choruspro/actions/workflows/tests.yml/badge.svg)
[![codecov](https://codecov.io/gh/antoine2116/go-choruspro/graph/badge.svg?token=E35WE8EF7Y)](https://codecov.io/gh/antoine2116/go-choruspro)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-choruspro is a [Go](https://go.dev/) client library for accessing the [ChorusPro API](https://api.gouv.fr/les-api/chorus-pro).

> Chorus Pro is the shared invoicing solution that has been set up for all suppliers (private or public) in the public sector (State, local authorities, etc.) to meet the legal requirements for electronic invoicing.
> (translated from [economie.gouv.fr](https://www.economie.gouv.fr/cedef/chorus-pro-aide#:~:text=Chorus%20Pro%20est%20la%20solution,en%20mati%C3%A8re%20de%20facturation%20%C3%A9lectronique.))

See the [Chorus Pro Community](https://communaute.chorus-pro.gouv.fr/) for more information (documentation, support, status, etc.)

## Installation

```bash
go get github.com/antoine2116/go-choruspro
```

## Usage

```go
import "github.com/antoine2116/go-choruspro"
```

Construct a new ChorusPro client. You must provide the following parameters:

```go
cfg := &choruspro.ClientConfig{
		BaseUrl:      "<piste_url>",
		AuthUrl:      "<piste_oauth_url>",
		ClientId:     "<piste_client_id>",
		ClientSecret: "<piste_client_secret>",
		Login:        "<chorus_pro_technical_credentials>",
	}

c, err := choruspro.NewClient().WithConfig(cfg)
```

### Examples

#### Get the list of currencies available

```go
res, err := c.Transverses.RecupererDevises(context.Background(), choruspro.CodeLangueFr)
```

#### Upload an invoice

```go
invoice := SoumettreFactureOptions{
  NumeroFactureSaisi: "123456",
  LignesPoste: &[]SoumettreFactureLignePoste{
    {
      Quantite: 1,
      PrixUnitaire: 100,
      Designation: "Test",
      CodeTVA: "TVA20",
    },
  },
}

res, err := c.Factures.SoumettreFacture(context.Background(), invoice)
```

## Services available

For now, only the following services are available:

- ✅ Factures
- ✅ Structures
- ✅ Transverses
- ✅ Utilisateurs
- ❌ Engagements
- ❌ FacturesTravaux

The other services will be added soon (PRs are welcome!)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Disclaimer

This project is not affiliated with the Chorus Pro project or the French government.
