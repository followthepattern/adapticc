import { Link } from "react-router-dom";

interface Entity {
    id?: string
}

export interface TableProperties<T = Entity> {
    entities: T[]
}

export interface CreateTableProperties<T> {
    headerColumns: string[]
    getEditLink: (entity: T) => string
    getViewLink: (entity: T) => string
    getCells: (entity: T) => string[]
}

export default function CreateTable<T extends Entity>({
    headerColumns, getEditLink,
    getViewLink, getCells }: CreateTableProperties<T>): React.ComponentType<TableProperties<T>> {
    return function ({ entities }: TableProperties<T>) {
        return (
            <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
                <table className="w-full text-sm text-left text-gray-500">
                    <thead className="text-xs text-gray-700 uppercase bg-gray-50">
                        <tr>
                            {headerColumns.map(column => (
                                <th scope="col" key={column} className="px-6 py-3">
                                    {column}
                                </th>)
                            )}
                            <th scope="col" key="action" className="px-6 py-3 text-right">Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {entities.map(entity => (
                            <tr key={entity.id} className="bg-white border-b hover:bg-gray-50">
                                {getCells(entity).map(cell => (
                                    <td key={cell} className="px-6 py-4 font-medium text-gray-900 whitespace-nowra">
                                        <Link to={getViewLink(entity)} className="hover:text-blue-900">
                                            {cell}
                                        </Link>
                                    </td>
                                ))}
                                <td className="px-6 py-4 text-right">
                                    <Link to={getEditLink(entity)} className="font-medium text-blue-600 hover:underline">
                                        Edit
                                    </Link>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div >
        )
    }
}