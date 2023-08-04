import { BarsArrowUpIcon, ChevronDownIcon } from '@heroicons/react/20/solid'
import { Listbox, Transition } from '@headlessui/react';
import { Fragment, useState } from 'react';
import { CheckIcon } from '@heroicons/react/24/outline';

export interface SortLabel {
    name: string
    asc: boolean
    code: string
}

export interface SortButtonProperties {
    sortByLables: SortLabel[]
    selectedSortLabel?: SortLabel
    sortOnChange: (v: SortLabel) => void
}

export default function SortButton({sortByLables, selectedSortLabel, sortOnChange}: SortButtonProperties) {
    let selectedValue = null;

    if (selectedSortLabel != null) {
        let sort = sortByLables.find(l => l.code == selectedSortLabel.code);
        if (sort) {
            selectedValue = sort;
        }
    }

    const OnChange = (v: SortLabel) => {
        sortOnChange(v);
    }

    return (
        <div className='relative'>
            <Listbox value={selectedValue} onChange={OnChange}>
                <Listbox.Button className="-ml-px inline-flex items-center gap-x-1.5 rounded-r-md px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
                    <span className="sr-only">{selectedValue ? selectedValue.name : "Sort"}</span>
                    <BarsArrowUpIcon className="-ml-0.5 h-5 w-5 text-gray-400" aria-hidden="true" />
                    {selectedValue ? selectedValue.name : "Sort"}
                    <ChevronDownIcon className="-mr-1 h-5 w-5 text-gray-400" aria-hidden="true" />
                </Listbox.Button>
                <Transition
                    as={Fragment}
                    enter="transition ease-out duration-100"
                    enterFrom="transform opacity-0 scale-95"
                    enterTo="transform opacity-100 scale-100"
                    leave="transition ease-in duration-75"
                    leaveFrom="transform opacity-100 scale-100"
                    leaveTo="transform opacity-0 scale-95"
                >
                    <Listbox.Options className="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                        <div className="py-1">
                            {sortByLables.map((sortLabel) => (
                                <Listbox.Option
                                    key={sortLabel.code}
                                    value={sortLabel}
                                    as={Fragment}
                                >
                                    {({ selected }) => (
                                        <span
                                            className="flex text-gray-700 justify-between px-4 py-2 text-sm hover:bg-gray-100 hover:text-gray-900"
                                        >{sortLabel.name} {selected && <CheckIcon className="h-5 w-5" />}</span>
                                    )}
                                </Listbox.Option>
                            ))}
                        </div>
                    </Listbox.Options>
                </Transition>
            </Listbox>
        </div>
    )
}