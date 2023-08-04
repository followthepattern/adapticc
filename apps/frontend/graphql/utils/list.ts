export interface OrderBy {
	name: string
	desc?: boolean
}

export interface ListFilter {
	page?: number
	pageSize?: number
	search?: string
	orderBy?: OrderBy[]
}

export interface List<T> {
	count?: number
	page?: number
	pageSize?: number
	search?: string
	data?: T[]
}