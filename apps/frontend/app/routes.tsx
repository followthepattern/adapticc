import { RouteObject } from "react-router-dom"
import Home from "./(public)/home"
import Login from "./(auth)/login/page"
import AuthLayout from "./(auth)/layout"
import AccountLayout from "./(account)/layout"
import Products from "./(account)/products/page"
import Users from "./(account)/users/page"
import Settings from "./(account)/settings/page"
import Profile from "./(account)/profile/page"
import Product from "./(account)/products/[id]/page"
import ProductEdit from "./(account)/products/[id]/edit/page"
import ProductNew from "./(account)/products/new/page"

export const Routes: RouteObject[] = [
    {
        path: "/",
        element: <Home />,
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
                element: <Users />
            },
            {
                path: "products",
                element: <Products />,
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
                path: "/settings",
                element: <Settings />
            },
            {
                path: "/profile",
                element: <Profile />
            },
        ]
    }
]