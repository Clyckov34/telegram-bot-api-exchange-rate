package centerBank

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlBankJSON = "https://cdn.cur.su/api/cbr.json"

//getCourse курс
func GetCourse(currencies string) (string, error) {
	res, err := http.Get(urlBankJSON)
	if err != nil {
		return "", errors.New("ошибка запроса на сервео ЦБ")
	}
	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("ошибка получения тела документа")
	}

	var course map[string]map[string]interface{}
	json.Unmarshal(bodyByte, &course)

	return fmt.Sprintf("%.2f", course["rates"][currencies]), nil
}
