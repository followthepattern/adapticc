import { gql } from "@apollo/client"
import { List } from "../utils/list"

export interface Product {
	id?: string
	title?: string
	description?: string
}

export interface Products {
	single?: Product
	list?: List<Product>
}

export const getSingleProduct = gql(`
query GetSingleProduct($id: String!) {
	products {
		single (id: $id) {
			id
			title
			description
		}
	}
}`)

export const getProducts = gql(`
query GetProducts($id: String, $pageSize: Uint, $page: Uint, $search: String) {
	products {
		list(
			filter: { id: $id, pageSize: $pageSize, page: $page, search: $search }
		) {
			count
			data {
				id
				title
				description
			}
			page
			pageSize
		}
	}
}
`)