import { RouteObject } from "react-router-dom"
import Login from "./(auth)/login/page"
import AuthLayout from "./(auth)/layout"
import AccountLayout from "./(account)/layout"
import Products from "./(account)/products/page"
import Profile from "./(account)/profile/page"
import Product from "./(account)/products/[id]/page"
import ProductEdit from "./(account)/products/[id]/edit/page"
import ProductNew from "./(account)/products/new/page"
import { ListPageWrapper } from "./(account)/components/listPage/listPageWrapper/listPageWrapper"
import UserNew from "./(account)/users/new/page"
import Users from "./(account)/users/page"
import User from "./(account)/users/[id]/page"
import UserEdit from "./(account)/users/[id]/edit/page"

export const Routes: RouteObject[] = [
    {
        path: "/",
        element: <AuthLayout />,
        children: [{
            path: "/",
            element: <Login />,

    }]
    },
    {
        path: "/",
        element: <AuthLayout />,
        children: [
            {
                path: "/login",
                element: <Login />
            }
        ]
    },
    {
        path: "/",
        element: <AccountLayout />,
        children: [
            {
                path: "/users",
                element: <ListPageWrapper Component={Users} />
            },
            {
                path: "/users/:id",
                element: <ListPageWrapper Component={User} />
            },
            {
                path: "/users/new",
                element: <UserNew />
            },
            {
                path: "/users/:id/edit",
                element: <UserEdit />
            },
            {
                path: "products",
                element: <ListPageWrapper Component={Products} />,
            },
            {
                path: "products/:id",
                element: <Product />
            },
            {
                path: "products/new",
                element: <ProductNew />
            },
            {
                path: "products/:id/edit",
                element: <ProductEdit />
            },
            {
                path: "/profile",
                element: <Profile />
            },
        ]
    }
]