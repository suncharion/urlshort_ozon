package service

import (
	"math/rand"
	"net/http"
	"net/url"
	"ozon/internal/model"
	"ozon/internal/storage"
	"strings"
	"time"
	"unicode/utf8"
)

type Service struct {
	storage storage.Storage
}

func NewService(strg storage.Storage) (*Service, error) {
	return &Service{
		storage: strg,
	}, nil
}

func (s *Service) getUrl(url string) (*model.ShortenedUrl, error) {
	return s.storage.Get(url)
}

func (s *Service) setUrl() {

}

const (
	numbersOffset    = 48
	upperCaseOffset  = 65 - 10
	lowerCaseOffset  = 97 - 36
	underscoreOffset = 95
)

func (s *Service) generateShort(full string) string {
	data := [10]byte{}
	rand.Seed(time.Now().Unix())

	// генерируем случайные символы из ascii-таблицы
	for i := range data {
		// всего у нас доступно 63 символа (10 цифр, 26 строчных букв и 26 заглавных, и подчеркивание)
		// генерируем случайное число от 0 до 62
		tmp := rand.Intn(63)

		/* символы в ascii идут группами, поэтому тоже мы разобъем ряд случайных чисел 0:63 на группы
		 	1) 0:9 - цифры
			2) 10:35 - заглавные буквы
			3) 36:61 - строчные буквы
			4) 62 - подчеркивание

			далее определяем группу, к которой принадлежит текущее число, и применяем к нему соответствующий сдвиг
			сдвиги посчитаны следующим образом: номер первого символа группы в ascii-таблице минус начало текущей группы
			наших случайных чисел. В итоге получаем код случайного ascii символа из списка допустимых
		*/
		if tmp <= 9 {
			tmp += numbersOffset
		} else if tmp <= 35 {
			tmp += upperCaseOffset
		} else if tmp <= 61 {
			tmp += lowerCaseOffset
		} else {
			tmp = underscoreOffset
		}
		data[i] = byte(tmp)
	}
	return string(data[:])
}

func (s *Service) HttpHandlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong http method", http.StatusNotImplemented)
	}

	fullStr := r.FormValue("full-url")
	_, err := url.ParseRequestURI(fullStr)
	if err != nil {
		http.Error(w, "Url is not valid", http.StatusBadRequest)
		return
	}
	short := model.ShortenedUrl{
		Original: fullStr,
	}
	for i := 0; i < 5; i++ { // пять попыток сгенерировать уникальный ключ
		short.Short = s.generateShort(fullStr)

		err := s.storage.Put(&short)
		if err == nil {
			break
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
	Your short link:<br>
	<a href="` + short.Short + `">` + short.Short + `</a>

	<br>
	<br>
	<a href="/">Generate another one</a>
	`))

}

func (s *Service) HttpHandlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Wrong http method", http.StatusNotImplemented)
	}

	urlArr := strings.Split(r.URL.Path, "/")
	urlArrClean := []string{}
	for i := range urlArr {
		urlArr[i] = strings.TrimSpace(urlArr[i])
		if urlArr[i] != "" {
			urlArrClean = append(urlArrClean, urlArr[i])
		}
	}
	if len(urlArrClean) == 0 {
		s.httpHandlerIndex(w, r)
		return
	}
	if utf8.RuneCountInString(urlArr[1]) != 10 {
		http.Error(w, "Url alias not present", http.StatusBadRequest)
		return
	}
	url, err := s.getUrl(urlArr[1])
	if err != nil {
		http.Error(w, "Url alias not present", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.Original, http.StatusSeeOther)
	return
}

func (s *Service) httpHandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(
		`<form method="POST" action="/save"><input type="text" name="full-url" placeholder="your url goes here" /><input type="submit"></form>`,
	))
}
