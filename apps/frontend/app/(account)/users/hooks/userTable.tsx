import { User } from "@/models/user";
import { Link } from "react-router-dom";

interface UserTableProperties {
    entities: User[]
}

export default function UserTable({ entities }: UserTableProperties) {
    return (
        <table className="w-full text-left">
            <thead className="bg-white">
                <tr>
                    <th scope="col" className="relative isolate py-3.5 pr-3 text-left text-sm font-semibold text-gray-900">
                        Email
                        <div className="absolute inset-y-0 right-full -z-10 w-screen border-b border-b-gray-200" />
                        <div className="absolute inset-y-0 left-0 -z-10 w-screen border-b border-b-gray-200" />
                    </th>
                    <th
                        scope="col"
                        className="hidden px-3 py-3.5 text-left text-sm font-semibold text-gray-900 md:table-cell"
                    >
                        Name
                    </th>
                    <th scope="col" className="relative py-3.5 pl-3">
                        <span className="sr-only">Edit</span>
                    </th>
                </tr>
            </thead>
            <tbody>
                {entities.map((user) => (
                    <tr key={user.id}>
                        <td className="relative py-4 pr-3 text-sm font-medium text-gray-900">
                            <Link to={`/users/${user.id}`} className="hover:text-indigo-900">
                                {user.email}
                            </Link>
                            <div className="absolute bottom-0 right-full h-px w-screen bg-gray-100" />
                            <div className="absolute bottom-0 left-0 h-px w-screen bg-gray-100" />
                        </td>
                        <td className="hidden px-3 py-4 text-sm text-gray-500 md:table-cell">{user.firstName} {user.lastName}</td>
                        <td className="relative py-4 pl-3 text-right text-sm font-medium">
                            <Link to={`/users/${user.id}/edit`} className="text-indigo-600 hover:text-indigo-900">
                                Edit<span className="sr-only">, {user.email}</span>
                            </Link>
                        </td>
                    </tr>
                ))}
            </tbody>
        </table>
    )
}