import SectionHeader from '../components/listPage/sectionHeader/sectionHeader';
import List from '../components/listPage/listPage';
import { ListPageComponentProperties } from '../components/listPage/listPageWrapper/listPageWrapper';
import { SortLabel, SetPageParams, SetSearchPatternParams, SetSortPatternParrams } from '../components/listPage/listPageWrapper/listingFunctions';
import useListUsers from './hooks/listUser';
import { User } from '@/models/user';
import CreateTable, { CreateTableProperties} from '../components/listPage/table/table';

export const RESOURCE_NAME = "Users"
export const RESOURCE_URL = "/users"

const sortByLables: SortLabel[] = [
  {
    code: "email",
    name: "Email",
    asc: true,
  },
  {
    code: "first_name",
    name: "First name",
    asc: true,
  },
  {
    code: "last_name",
    name: "Last name",
    asc: true,
  }
];

function mapUserToRowCells(user: User): string[] {
  const email = user.email ?? ""
  const name = `${user.firstName} ${user.lastName}`
  return [email, name]
}

function userViewLink(user: User): string {
  return `/users/${user.id}`
}

function userEditLink(user: User): string {
  return `/users/${user.id}/edit`
}

export default function Users(props: ListPageComponentProperties) {
  const sortOnChange = (sortLabel: SortLabel) => {
    SetSortPatternParrams(props.searchParams, props.setSearchParams, sortLabel);
  }

  const searchInputOnChange = (searchString: string) => {
    SetSearchPatternParams(props.searchParams, props.setSearchParams, searchString);
  }

  const pageOnChange = (page: number) => {
    SetPageParams(props.searchParams, props.setSearchParams, page);
  }

  const selectedSortLabel = sortByLables.find(l => l.code == props.sortProps.sortLabel?.code);

  const createTableProperties: CreateTableProperties<User> = {
    headerColumns: ["Email", "Name"],
    getViewLink: userViewLink,
    getEditLink: userEditLink,
    getCells: mapUserToRowCells,
  }

  const userTable = CreateTable(createTableProperties);

  return (
    <div>
      <SectionHeader
        resourceName={RESOURCE_NAME}
        resourceUrl={RESOURCE_URL}
        searchInputOnChange={searchInputOnChange}
        sortOnChange={sortOnChange}
        searchInput={props.filterProps.searchString}
        sortByLables={sortByLables}
        selectedSortLabel={selectedSortLabel}
      />
      <div className="flow-root mt-8 overflow-hidden">
        <div className="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
          <List {...props}
            sortProps={{ sortLabel: selectedSortLabel }}
            onPageChange={pageOnChange}
            useList={useListUsers}
            tableComponent={userTable}
          />
        </div>
      </div>
    </div>
  )
}