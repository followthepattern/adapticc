'use client';

import { ListPageProperties } from "@/components/listPage";
import { QueryResponse } from "@/graphql/query";
import { getUsers } from "@/graphql/users/query";
import { gql, useLazyQuery } from "@apollo/client";
import { Link } from "react-router-dom";
import { useEffect } from "react";

export default function Users() {
  const [executeGetUsers, { data, called, loading, error }] = useLazyQuery<QueryResponse>(gql(getUsers))

  useEffect(() => {
    if (!called) {
      executeGetUsers({
        variables: {
          page: 1,
          pageSize: 10,
        }
      });
    }
  }, [called, executeGetUsers]);

  if (error) {
    return <div>{JSON.stringify(error)}</div>
  }

  // if (loading) {
  //   return <div>Loading...</div>
  // }

  const list = data?.users?.list;

  return (
    <>
      <h1 className="">
        Users
      </h1>
      <div>List</div>
      <div>
        {JSON.stringify(list)}
      </div>
      <Link to="/products">Products</Link>
    </>
  )
}