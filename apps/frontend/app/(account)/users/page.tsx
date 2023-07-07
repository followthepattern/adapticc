'use client';

import { ListPageProperties } from "@/components/listPage";
import { QueryResponse } from "@/graphql/query";
import { getUsers } from "@/graphql/users/query";
import { gql, useLazyQuery } from "@apollo/client";
import Link from "next/link";
import { useEffect } from "react";

export default function Page({searchParams}: ListPageProperties) {
  const [executeGetUsers, { data, called, loading, error }] = useLazyQuery<QueryResponse>(gql(getUsers))

  useEffect(() => {
    if (!called) {
      executeGetUsers({
        variables: {
          page: searchParams.page,
          pageSize: searchParams.pageSize,
        }
      });
    }
  }, [called, executeGetUsers, searchParams]);

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
      <Link href="/products">Products</Link>
    </>
  )
}