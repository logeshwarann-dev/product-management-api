package store

type Product struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Price       int    `json:"price,omitempty"`
}

type Products struct {
	List []Product `json:"products,omitempty"`
}
