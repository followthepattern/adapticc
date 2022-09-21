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

export const getUsers =
`query GetUsers (
	$page: Uint,
	$pageSize: Uint
	$search: String) {
	users {
		list (filter: {
			pageSize: $pageSize,
			page: $page,
			search: $search,
		}) {
			page
			pageSize
			count
			data {
				id
				email
				firstName
				lastName
			}
		}
	}
}`