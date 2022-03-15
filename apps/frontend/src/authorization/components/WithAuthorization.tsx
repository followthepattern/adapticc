import { useLazyQuery, gql } from "@apollo/client";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import UserContext from "../../contexts/UserContext";
import { getUserProfile } from "../../graphql/user/query";
import { useUserStore } from "../../utils/store";

function WithAuthorization(
  Component: (props: any) => JSX.Element
): JSX.Element {
  const [executeGetUserProfile, { data, called, loading, error }] =
    useLazyQuery(gql(getUserProfile));
  const navigate = useNavigate();
  const { token, removeToken } = useUserStore();

  useEffect(() => {
    if (!called) {
      executeGetUserProfile();
    }

    if (!token) {
      navigate("/login");
    }

    if (called && !error && !loading) {
      let id = data?.users?.profile?.id;
      if (!id) {
        removeToken();
        navigate("/login");
      }
    }
  }, [data]);

  if (error) {
    return <div>{JSON.stringify(error)}</div>;
  }

  return called && !loading ? (
    <UserContext.Provider value={data.users.profile}>
      <Component />
    </UserContext.Provider>
  ) : (
    <div>Loading...</div>
  );
}

export default WithAuthorization;
