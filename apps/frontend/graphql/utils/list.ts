export interface OrderBy {
	name: string
	desc?: boolean
}

export interface ListQueryParams {
	page?: number
	pageSize?: number
	search?: string
	orderBy?: OrderBy[]
}

export interface ListResponse<T> {
	count?: number
	page?: number
	pageSize?: number
	search?: string
	data?: T[]
}

export type ListQueryResult<Data = any> = [
    (params: ListQueryParams) => void,
    {
        loading: boolean;
        data?: Data;
        error?: any;
        itemNotFound?: boolean;
        called: boolean;
    }
];