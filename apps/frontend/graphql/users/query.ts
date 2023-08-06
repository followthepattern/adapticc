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
query GetUsers(
	$search: String
	$pageSize: Uint
	$page: Uint
	$orderBy: [OrderBy!]
) {
	users {
		list(
			pagination: { pageSize: $pageSize, page: $page }
			filter: { search: $search }
			orderBy: $orderBy
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
}`