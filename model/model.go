package model

// Согласно задаче нужно детектировать тип объект в рамках единого хендлера,
// типы мы опишем тут чтобы централизованно менять их имена
const ByerType = "buyer"
const ShopType = "shop"

type Buyer struct {
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	Age         int    `json:"age"`
	Registation string `json:"registation"`
}

type Shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Working bool   `json:"working"`
	Owner   string `json:"owner"`
}

// согласно задаче детекция объектов обзятельна
// из-за условия QQ с использованием одной функции,
// умеющей принимать в качестве входного
// значения оба типа записи. QQ
// но с учётом других противоречащих друг другу условий задачи,
// нет иного выхода, кроме как переложить это на плечи прользователя
type SearchRequest struct {
	Type     string `json:"type"`
	LastName string `json:"first_name"` // Для покупателя
	Name     string `json:"name"`       // Для магазина
}
