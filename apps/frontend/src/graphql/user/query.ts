export const getUser = `{
	users {
		single (id: "21efde9d-8091-477a-919a-6156bad4ba23") {
			email
			firstName
			lastName
			updatedAt
			createdAt
		}
	}
}`

export const getUserProfile = `{
	users {
		profile {
			id
			email
			firstName
			lastName
		}
	}
}`