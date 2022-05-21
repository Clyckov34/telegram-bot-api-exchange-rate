package check

import "errors"

//Flag проверка на пустоту флагов
func Flag(data string) error {
	if len(data) != 0 {
		return nil
	}
	return errors.New("ошибка: не указаны параметры флага")
}
