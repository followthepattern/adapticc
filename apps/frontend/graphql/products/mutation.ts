export const updateProduct = `
mutation UpdateProduct($model: ProductInput!) {
	products {
		update(model: $model) {
			code
		}
	}
}`

export const createProduct = `
mutation CreateProduct($model: ProductInput!) {
	products {
		create(model: $model) {
			code
		}
	}
}`


export const deleteProduct = `
mutation DeleteProduct($id: String!) {
	products {
		delete(id: $id) {
			code
		}
	}
}`