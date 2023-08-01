import { Link } from "react-router-dom";
import { Product } from "@/graphql/products/query";

interface TableProperties {
    entities: Product[]
}

export default function Table({entities}: TableProperties) {
    return (
        <table className="w-full text-left">
            <thead className="bg-white">
                <tr>
                    <th scope="col" className="relative isolate py-3.5 pr-3 text-left text-sm font-semibold text-gray-900">
                        Title
                        <div className="absolute inset-y-0 right-full -z-10 w-screen border-b border-b-gray-200" />
                        <div className="absolute inset-y-0 left-0 -z-10 w-screen border-b border-b-gray-200" />
                    </th>
                    <th
                        scope="col"
                        className="hidden px-3 py-3.5 text-left text-sm font-semibold text-gray-900 md:table-cell"
                    >
                        Description
                    </th>
                    <th scope="col" className="relative py-3.5 pl-3">
                        <span className="sr-only">Edit</span>
                    </th>
                </tr>
            </thead>
            <tbody>
                {entities.map((product) => (
                    <tr key={product.id}>
                        <td className="relative py-4 pr-3 text-sm font-medium text-gray-900">
                            <Link to={`/products/${product.id}`} className="hover:text-indigo-900">
                                {product.title}
                            </Link>
                            <div className="absolute bottom-0 right-full h-px w-screen bg-gray-100" />
                            <div className="absolute bottom-0 left-0 h-px w-screen bg-gray-100" />
                        </td>
                        <td className="hidden px-3 py-4 text-sm text-gray-500 md:table-cell">{product.description}</td>
                        <td className="relative py-4 pl-3 text-right text-sm font-medium">
                            <Link to={`/products/${product.id}/edit`} className="text-indigo-600 hover:text-indigo-900">
                                Edit<span className="sr-only">, {product.title}</span>
                            </Link>
                        </td>
                    </tr>
                ))}
            </tbody>
        </table>
    )
}