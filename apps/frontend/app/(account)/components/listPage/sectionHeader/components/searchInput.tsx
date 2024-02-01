import SearchIcon from '@/app/icons/search';
import classNames from 'classnames';
import { Controller, useForm } from 'react-hook-form';

interface SearchInputValues {
    search: string;
}

interface SearchInputProperties {
    onChange: (s: string) => void
    search?: string
    className?: string
}

export default function SearchInput(props: SearchInputProperties) {
    const { getValues, control } = useForm<SearchInputValues>();

    const handleInputChange = (_: any) => {
        const formValues = getValues();

        props.onChange(formValues.search);
    };

    return (
        <div className={classNames(props.className, "relative flex-grow focus-within:z-10")}>
            <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                <SearchIcon className="w-4 h-4 text-gray-400" aria-hidden="true" />
            </div>
            <Controller
                control={control}
                name="search"
                defaultValue={props.search}
                render={({ field: { onChange } }) => (
                    <input
                        type="text"
                        className="w-full h-full py-2 pl-10 text-gray-900 border border-gray-300 rounded-lg placeholder:text-gray-400 focus:ring-0 focus:ring-inset"
                        placeholder="Search"
                        onChange={(e) => {
                            onChange(e);
                            handleInputChange(e);
                        }}
                        defaultValue={props.search}
                    />
                )}
            />
        </div >
    )
}