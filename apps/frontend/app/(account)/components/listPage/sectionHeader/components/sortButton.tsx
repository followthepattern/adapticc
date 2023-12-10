import { BarsArrowUpIcon } from '@heroicons/react/20/solid'
import { Listbox, Transition } from '@headlessui/react';
import { Fragment } from 'react';
import { CheckIcon } from '@heroicons/react/24/outline';
import { SortLabel } from '../../listPageWrapper/listingFunctions';
import ChevronDownIcon from '@/app/icons/chevronDown';
import classNames from 'classnames';

export interface SortButtonProperties {
    sortByLables: SortLabel[]
    selectedSortLabel?: SortLabel
    sortOnChange: (v: SortLabel) => void
    className?: string
}

export default function SortButton(props: SortButtonProperties) {
    const OnChange = (v: SortLabel) => {
        props.sortOnChange(v);
    }

    const initValue: SortLabel | null = props.selectedSortLabel ? props.selectedSortLabel : null;

    return (
        <div className={classNames(props.className, "relative")}>
            <Listbox value={initValue} onChange={OnChange}>
                <Listbox.Button className="inline-flex items-center justify-center w-full px-4 py-2 text-sm font-semibold rounded-lg gap-x-2 ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
                    {initValue ? initValue.name : "Sort"}
                    <ChevronDownIcon className="w-4 h-4 text-gray-400" aria-hidden="true" />
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
                    <Listbox.Options className="absolute right-0 z-10 w-40 mt-2 origin-top-right bg-white border rounded-lg shadow-lg border-gray-50">
                        <div className="py-1">
                            {props.sortByLables.map((sortLabel) => (
                                <Listbox.Option
                                    key={sortLabel.code}
                                    value={sortLabel}
                                    as={Fragment}
                                >
                                    {({ selected }) => (
                                        <span
                                            className="flex justify-between px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                                        >{sortLabel.name} {selected && <CheckIcon className="w-5 h-5" />}</span>
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