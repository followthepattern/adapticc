export const getSingleProduct = `
query GetSingleProduct($id: String!) {
	products {
		single (id: $id) {
			id
			title
			description
		}
	}
}`

export const getProducts = `
query GetProducts(
	$pageSize: Uint
	$page: Uint
	$search: String
	$orderBy: [OrderBy!]
) {
	products {
		list(
			pagination: { pageSize: $pageSize, page: $page }
			filter: { search: $search }
			orderBy: $orderBy
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
}`