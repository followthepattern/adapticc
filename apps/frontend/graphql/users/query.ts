import { List } from "../utils/list"

export interface User {
	id?: string
	email?: string
	firstName?: string
	lastName?: string
}

export interface Users {
	profile?: User
	single?: User
	list?: List<User>
}

export const getUserProfile = `
query {
	users {
		profile {
			id
			email
			firstName
			lastName
		}
	}
}`

export const getSingleUser = `
query GetSingleUser($id: String!) {
	users {
		single (id: $id) {
			id
			email
			firstName
			lastName
		}
	}
}`

export const getUsers = `
query GetUsers($search: String, $pageSize: Uint, $page: Uint) {
	users {
		list (filter:{
			search: $search,
			pageSize: $pageSize,
			page: $page,
		}
		) {
			count
			data {
				id
				email
				firstName
				lastName
			}
			page
			pageSize
		}
	}
}
`