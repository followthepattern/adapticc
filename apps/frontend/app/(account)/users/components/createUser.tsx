import { CreateMutationResult, MutationResponse } from "@/graphql/mutation";
import { User } from "@/models/user";
import useCreate from "@/graphql/hooks/useCreate";
import { createUser as graphQL } from "@/graphql/users/mutation";

export default function useCreateUsers(): CreateMutationResult<User, number> {
    const parseResult = (data?: MutationResponse | null) : number | undefined  => {
        return data?.users?.create?.code;
    }

    return useCreate<User>({parseResult, graphQL})
}