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
		Code            int    `json:"code"`
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
		Available     int     `json:"available"`
		Returned      int     `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Stories struct {
		Available     int     `json:"available"`
		Returned      int     `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Events struct {
		Available     int     `json:"available"`
		Returned      int     `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Series struct {
		Available     int     `json:"available"`
		Returned      int     `json:"returned"`
		CollectionURI string  `json:"collectionURI"`
		Items         []Items `json:"items"`
	}
	Results struct {
		ID          int       `json:"id"`
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
		Offset  int       `json:"offset"`
		Limit   int       `json:"limit"`
		Total   int       `json:"total"`
		Count   int       `json:"count"`
		Results []Results `json:"results"`
	}
	// ReqBody for searching
	Payload struct {
		Name           string `form:"name"`
		NameStartsWith string `form:"nameStartsWith"`
		ModifiedSince  string `form:"modifiedSince"`
		Comics         string `form:"comics"`
		Series         string `form:"series"`
		Events         string `form:"events"`
		Stories        string `form:"stories"`
		OrderBy        string `form:"orderBy"`
		Limit          int    `form:"limit"`
		Offset         int    `form:"offset"`
		Apikey         string `form:"apikey"`
		Hash           string `form:"hash"`
		Ts             int    `form:"ts"`
		Page           int    `form:"page,default:1"`
	}
	Pagination struct {
		Limit          int  `json:"limit"`
		Offset         int  `json:"offset"`
		TotalPages     int  `json:"totalPages"`
		TotalDataCount int  `json:"totalDataCount"`
		CurrentPage    int  `json:"currentPage"`
		Proceed        bool `json:"proceed"`
	}
)
