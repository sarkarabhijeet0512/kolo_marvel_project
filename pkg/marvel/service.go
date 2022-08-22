package marvel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kolo_marvel_project/er"
	"kolo_marvel_project/pkg/cache"
	utils "kolo_marvel_project/utils/common"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf *viper.Viper
	log  *logrus.Logger
	// Repo         Repository
	CacheService *cache.Service
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, CacheService *cache.Service) *Service {
	return &Service{
		conf: conf,
		log:  log,
		// Repo:         dbRepo,
		CacheService: CacheService,
	}
}

const (
	itemsPerPage = 10
)

// IsDBActive gets user data by her userID
func (s *Service) FetchCharacterDetails(payload *Payload) (mcd *MarvelCharacterDetails, err error) {

	payload.Ts = 1
	payload.Apikey = s.conf.GetString("MARVEL_PUBLIC_KEY")
	payload.Hash = utils.GetMD5Hash(fmt.Sprint(payload.Ts) + s.conf.GetString("MARVEL_PRIVATE_KEY") + s.conf.GetString("MARVEL_PUBLIC_KEY"))

	err = s.CacheService.Repo.Get(payload.NameStartsWith+fmt.Sprint(payload.Page), &mcd)
	paginate := GetDataPage(payload.Page)
	payload.Limit = paginate.Limit
	payload.Offset = paginate.Offset

	if err != nil && err.Error() == "cache: key not found." {
		mcd, err = s.MarvelCharacterList(payload)
		if err != nil {
			s.log.Info("err :", err, " ,obj :", mcd, " ,payload :", payload)
			err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
			return
		}
		err = s.CacheService.Repo.Set(payload.NameStartsWith+fmt.Sprint(payload.Page), mcd, 5*time.Minute)
		if err != nil {
			err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
			return
		}
	}
	return
}

// pages start at 1, can't be 0 or less.
func GetDataPage(page int) (pagination Pagination) {
	pagination.Offset = (page - 1) * itemsPerPage
	pagination.Limit = pagination.Offset + itemsPerPage
	return
}
func (s *Service) MarvelCharacterList(payload *Payload) (mcd *MarvelCharacterDetails, err error) {

	var (
		client = &http.Client{}
		params = url.Values{}
	)

	base, _ := url.Parse(s.conf.GetString("MARVEL_BASE_SVC") + "/v1/public/characters")

	params.Add("apikey", payload.Apikey)
	params.Add("hash", payload.Hash)
	params.Add("ts", fmt.Sprint(payload.Ts))
	params.Add("limit", fmt.Sprint(payload.Limit))
	params.Add("offset", fmt.Sprint(payload.Offset))

	if payload.NameStartsWith != "" {
		params.Add("nameStartsWith", payload.NameStartsWith)
	}
	base.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &mcd)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusInternalServerError)
		return
	}
	return
}
