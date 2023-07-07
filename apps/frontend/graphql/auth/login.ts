export const login = `
mutation Login($email: String!, $password: String!) {
	authentication {
		login(email: $email, password: $password) {
			jwt
			expires_at
		}
	}
}`