import { MagnifyingGlassIcon } from '@heroicons/react/20/solid';
import classNames from 'classnames';
import { Controller, useForm } from 'react-hook-form';

interface SearchInputValues {
    search: string;
}

interface SearchInputProperties {
    onChange: (s: string) => void
    search?: string
}

export default function SearchInput(props: SearchInputProperties) {
    const { getValues, control } = useForm<SearchInputValues>();

    const commonClasses = "w-full rounded-none rounded-l-md border-0 py-1.5 pl-10 text-gray-900 ring-1 ring-gray-300 ring-inset placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600";

    const handleInputChange = (_: any) => {
        const formValues = getValues();

        props.onChange(formValues.search);
    };

    return (
        <div className="relative flex-grow focus-within:z-10">
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                <MagnifyingGlassIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
            </div>
            <Controller
                control={control}
                name="search"
                defaultValue={props.search}
                render={({ field: { onChange } }) => (
                    <input
                        type="text"
                        className={classNames("block ring-inset", commonClasses)}
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