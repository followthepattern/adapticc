import { createContext } from "react";

interface UserProfile {
    isAuthenticated: boolean,
    email?: string,
    firstName?: string,
    lastName?: string,
}

const UserContext = createContext<UserProfile>({
    isAuthenticated: false,
});

export default UserContext;
