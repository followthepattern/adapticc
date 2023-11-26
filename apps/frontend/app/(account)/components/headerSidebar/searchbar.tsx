import SearchIcon from "@/app/icons/search";

interface SearchbarProperties { }

export default function Searchbar(props: SearchbarProperties) {
    return (
        <div className="relative flex flex-auto">
            <SearchIcon
                className="absolute left-0 w-5 h-full text-gray-600 pointer-events-none"
            />
            <input
                className="block w-full h-full py-0 pl-10 border-0 placeholder:text-gray-400 focus:ring-0"
                placeholder="Resources..."
                type="text"
                name="search"
            />
        </div>
    )
}