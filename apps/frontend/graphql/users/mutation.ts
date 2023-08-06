export const createUser = `
mutation CreateUser($model: UserInput!) {
	users {
		create(model: $model) {
			code
		}
	}
}
`

export const updateUser = `
mutation UpdateUser($id: String!, $model: UserInput!) {
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