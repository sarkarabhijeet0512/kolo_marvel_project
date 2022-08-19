package marvel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
func (s *Service) FetchCharacterDetails(payload *Payload) (mcd *MarvelCharacterDetails) {
	url := "https://gateway.marvel.com:443/v1/public/characters?apikey=07216b5045a3252f244a86a0de131be3&hash=ff4aeb8af63fc49e26997e7f4a4f9991&ts=1&name=3-D%20Man"
	method := "GET"
	client := &http.Client{}
	b, _ := json.Marshal(payload)
	par := strings.NewReader(string(b))
	req, err := http.NewRequest(method, url, par)

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
