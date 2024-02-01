export interface MutationResult {
    code?: number
}

export interface MutationResourceResponse {
    create?: MutationResult
	update?: MutationResult
	delete?: MutationResult
}

export type CreateMutationResult<Entity = any, TResult = any> = [
    (model: Entity) => void,
    {
        createLoading: boolean;
        createResult?: TResult;
        createError?: any;
    }
];

export type UpdateMutationResult<Entity = any, TResult = any> = [
    (id: string, model: Entity) => void,
    {
        updateLoading: boolean;
        updateResult?: TResult;
        updateError?: any;
    }
];

export type DeleteMutationResult<Entity = any, TResult = any> = [
    (id: Entity) => void,
    {
        deleteLoading: boolean;
        deleteResult?: TResult;
        deleteError?: any;
    }
];

export interface MutationResponse {
    users: MutationResourceResponse
    products: MutationResourceResponse
}