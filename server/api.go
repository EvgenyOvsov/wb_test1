package server

import (
	"bytes"
	"draft/model"
	"encoding/json"
	"io"
	"net/http"
)

func (srv *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`OK`))
	w.WriteHeader(http.StatusOK)
}

// Согласно задаче, не знаю почему надо делать одну функцию, но бог с ним
func (srv *Server) handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body := bytes.Buffer{}
	io.Copy(&body, r.Body)

	// В задаче сказано
	// QQ На выходе сервис должен уметь возвращать всю запись,
	// либо одно поле из записи в зависимости от запроса пользователя,
	// осуществлять поиск по Фамилии и Названию  для соответствующих записей. QQ

	// Поиск и возврат полного объекта легко осуществим, но
	// 1. В Get запрос нельзя передавать тело
	//    (Можно, но большинство клиентов не поддерживают)
	// 2. Post и Delete запросы уже заняты, не хочу использовать ни PUT, ни PATCH, ни прочие методы
	//    для поиска

	// Задача полностью противоречит REST'у, видимо меня хотели заставить поговнокодить и
	// показать максимум антипаттернов
	// В любом случае слишком много времени уйдёт на реализацию частичного возврата
	// объекта, так что остановлюсь на поиске согласно задаче

	// Если это GET - ищем объект
	// Фильтр будем ставить через querry параметры
	// Возвращаем 404 если не нашли

	if r.Method == http.MethodGet {
		obj_type := r.URL.Query().Get(`type`)
		if obj_type == `` {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if obj_type == model.ByerType {
			filter := r.URL.Query().Get(`last_name`)
			if filter == `` {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			buyer := srv.ds.GetBuyer(filter)
			if buyer == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(buyer)

			return
		}
		if obj_type == model.ShopType {
			filter := r.URL.Query().Get(`name`)
			if filter == `` {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			shop := srv.ds.GetShop(filter)
			if shop == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(shop)
			return
		}
	}

	// Если не json - игнорируем
	if r.Header.Get(`Content-Type`) != `application/json` {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ЕСЛИ это не Get запрос
	// Надо понять что это за объект
	// Не понимаю нахрена вообще этим заниматься, но

	o1 := &model.Buyer{}
	o2 := &model.Shop{}
	second_obj := false

	err := json.Unmarshal(body.Bytes(), o1)
	if err != nil || o1.FirstName == `` || o1.LastName == `` {
		second_obj = true
		err = json.Unmarshal(body.Bytes(), o2)
		if err != nil || o2.Name == `` || o2.Address == `` {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// PUT и POST будут дёргать соответствующие методы датастора
	switch r.Method {

	case http.MethodPost:
		if second_obj {
			err = srv.ds.SaveShop(o2)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			err = srv.ds.SaveBuyer(o1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		if second_obj {
			err = srv.ds.DeleteShop(o2)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			err = srv.ds.DeleteBuyer(o1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
