package domain

type Product struct {
	id          string
	name        string
	description string
	price       float64
	stock       int
	createdAt   int64
	updatedAt   int64
	deletedAt   int64
}

func NewProduct(id, name, description string, price float64, stock int, createdAt, updatedAt, deletedAt int64) *Product {
	return &Product{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

func (p *Product) Id() string {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Stock() int {
	return p.stock
}

func (p *Product) CreatedAt() int64 {
	return p.createdAt
}

func (p *Product) UpdatedAt() int64 {
	return p.updatedAt
}

func (p *Product) DeletedAt() int64 {
	return p.deletedAt
}
