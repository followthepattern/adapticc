export const createUser = `
mutation CreateUser($model: CreateUserInput!) {
	users {
		create(model: $model) {
			code
		}
	}
}
`

export const updateUser = `
mutation UpdateUser($id: String!, $model: UpdateUserInput!) {
	users {
		update(id: $id, model: $model) {
			code
		}
	}
}`


export const deleteUser = `mutation DeleteUser($id: String!) {
	users {
		delete(id: $id) {
			code
		}
	}
}`