package marvel

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
	),
)

type (
	MarvelCharacterDetails struct {
		Code            string `json:"code"`
		Status          string `json:"status"`
		Copyright       string `json:"copyright"`
		AttributionText string `json:"attributionText"`
		AttributionHTML string `json:"attributionHTML"`
		Data            Data   `json:"data"`
		Etag            string `json:"etag"`
	}
	Urls struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}
	Thumbnail struct {
		Path      string `json:"path"`
		Extension string `json:"extension"`
	}
	Items struct {
		ResourceURI string `json:"resourceURI"`
		Name        string `json:"name"`
		Type        string `json:"type,omitempty"`
	}
	Comics struct {
		Available     string  `json:"available"`
		Returned      string  `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Stories struct {
		Available     string  `json:"available"`
		Returned      string  `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Events struct {
		Available     string  `json:"available"`
		Returned      string  `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Series struct {
		Available     string  `json:"available"`
		Returned      string  `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Results struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Modified    string    `json:"modified"`
		ResourceURI string    `json:"resourceURI"`
		Urls        []Urls    `json:"urls"`
		Thumbnail   Thumbnail `json:"thumbnail"`
		Comics      Comics    `json:"comics"`
		Stories     Stories   `json:"stories"`
		Events      Events    `json:"events"`
		Series      Series    `json:"series"`
	}
	Data struct {
		Offset  string    `json:"offset"`
		Limit   string    `json:"limit"`
		Total   string    `json:"total"`
		Count   string    `json:"count"`
		Results []Results `json:"results"`
	}
	// ReqBody for searching
	Payload struct {
		Name          string `json:"name"`
		NameStartWith string `json:"nameStartWith"`
		ModifiedSince string `json:"modifiedSince"`
		Comics        string `json:"comics"`
		Series        string `json:"series"`
		Events        string `json:"events"`
		Stories       string `json:"stories"`
		OrderBy       string `json:"orderBy"`
		Limit         string `json:"limit"`
		Offset        string `json:"offset"`
		Apikey        string `json:"apikey"`
		Hash          string `json:"hash"`
		Ts            string `json:"ts"`
	}
)
