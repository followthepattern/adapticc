import { gql } from "@apollo/client"
import { MutationResult } from "../utils/mutationResult"

export interface ProductMutation {
	update?: MutationResult
	create?: MutationResult
	delete?: MutationResult
}

export const updateProduct = gql(`
mutation UpdateProduct($model: ProductInput!) {
	products {
		update(model: $model) {
			code
		}
	}
}`)

export const createProduct = gql(`
mutation CreateProduct($model: ProductInput!) {
	products {
		create(model: $model) {
			code
		}
	}
}`)


export const deleteProduct = gql(`
mutation DeleteProduct($id: String!) {
	products {
		delete(id: $id) {
			code
		}
	}
}`)