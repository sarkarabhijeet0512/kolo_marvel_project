package marvel

import (
	"fmt"
	"io/ioutil"
	utils "kolo_marvel_project/utils/common"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf *viper.Viper
	log  *logrus.Logger
	Repo Repository
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, dbRepo Repository) *Service {
	return &Service{
		conf: conf,
		log:  log,
		Repo: dbRepo,
	}
}

// IsDBActive gets user data by her userID
func (s *Service) FetchCharacterDetails(payload *Payload) (mcd *MarvelCharacterDetails, err error) {
	payload.Ts = 1
	payload.Apikey = s.conf.GetString("MARVEL_PUBLIC_KEY")
	payload.Hash = utils.GetMD5Hash(fmt.Sprint(payload.Ts) + s.conf.GetString("MARVEL_PRIVATE_KEY") + s.conf.GetString("MARVEL_PUBLIC_KEY"))
	payload.Limit = 10

	base, _ := url.Parse(s.conf.GetString("MARVEL_BASE_SVC") + "/v1/public/characters")

	client := &http.Client{}
	params := url.Values{}

	params.Add("apikey", payload.Apikey)
	params.Add("hash", payload.Hash)
	params.Add("ts", fmt.Sprint(payload.Ts))
	params.Add("limit", fmt.Sprint(payload.Limit))
	if payload.Name != "" {
		params.Add("name", payload.Name)
	}
	if payload.NameStartsWith != "" {
		params.Add("nameStartsWith", payload.NameStartsWith)
	}
	base.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	return
}
