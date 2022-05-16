package centerBank

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlBankJSON = "https://cdn.cur.su/api/cbr.json"

//request запрос в банк о курсах валют
func request() ([]byte, error) {
	res, err := http.Get(urlBankJSON)
	if err != nil {
		return nil, errors.New("ошибка запроса на сервео ЦБ")
	}
	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("ошибка получения тела документа")
	}

	return bodyByte, nil
}

//getCourse курс
func GetCourse(currencies string) (string, error) {
	bodyByte, err := request()
	if err != nil {
		return "", err
	}

	var course map[string]map[string]interface{}
	json.Unmarshal(bodyByte, &course)

	return fmt.Sprintf("%.2f", course["rates"][currencies]), nil
}
