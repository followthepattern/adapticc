import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import UserContext from "../../contexts/UserContext";

function WithAuthorization(
  Component: (props: any) => JSX.Element
): JSX.Element {
  // download account details ...

  const navigate = useNavigate();

  let isAuthenticated = true;

  useEffect(() => {
    if (!isAuthenticated) {
      navigate("/login");
    }
  }, []);

  return (
    <UserContext.Provider
      value={{
        isAuthenticated: isAuthenticated,
      }}
    >
      <Component />
    </UserContext.Provider>
  );
}

export default WithAuthorization;
