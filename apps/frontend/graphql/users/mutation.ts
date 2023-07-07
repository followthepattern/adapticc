import { gql } from "@apollo/client"
import { MutationResult } from "../utils/mutationResult"

export interface Users {
	update?: MutationResult
	delete?: MutationResult
}

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