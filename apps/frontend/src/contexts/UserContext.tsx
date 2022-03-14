import { createContext } from "react";

interface UserProfile {
    email?: string,
    firstName?: string,
    lastName?: string,
}

const UserContext = createContext<UserProfile>({});

export default UserContext;
