import { gql } from "@apollo/client"

export const updateUser = `
mutation UpdateUser($id: String!, $firstName: String, $lastName: String){
	users {
		update (id: $id, firstName: $firstName, lastName: $lastName) {
			code
		}
	}
}`


export const deleteUser = gql(`mutation DeleteUser($id: String!) {
	users {
		delete(id: $id) {
			code
		}
	}
}`)