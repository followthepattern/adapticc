import { updateUser as graphQL } from "@/graphql/users/mutation";
import { MutationResponse, UpdateMutationResult } from "@/graphql/mutation";
import useUpdate from "@/graphql/hooks/useUpdate";
import { User } from "@/models/user";

export default function useUpdateUser(): UpdateMutationResult<User,number | undefined> {
const parseResult = (data?: MutationResponse | null) : number | undefined  => {
        return data?.users?.update?.code;
    }

    return useUpdate({parseResult, graphQL})
}