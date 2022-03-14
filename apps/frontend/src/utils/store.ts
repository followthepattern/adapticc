import create, { GetState, SetState } from "zustand";
import { devtools, persist, StoreApiWithPersist } from "zustand/middleware";

const storeNames = {
  navbar: "navbar-store",
  userStore: "user-store",
};

interface NavbarExpandedSate {
  navbarExpanded: boolean;
  toggleNavbarExpanded: () => void;
  closeNavbar: () => void;
  openNavbar: () => void;
}

export const useNavbarStore = create<
  NavbarExpandedSate,
  SetState<NavbarExpandedSate>,
  GetState<NavbarExpandedSate>,
  StoreApiWithPersist<NavbarExpandedSate>
>(
  persist(
    devtools((set) => ({
      navbarExpanded: true,
      toggleNavbarExpanded: () =>
        set((state) => ({ navbarExpanded: !state.navbarExpanded })),
      closeNavbar: () => set(() => ({ navbarExpanded: false })),
      openNavbar: () => set(() => ({ navbarExpanded: true })),
    })),
    {
      name: storeNames.navbar,
      getStorage: () => localStorage,
    }
  )
);

interface UserState {
  token: string | null;
  setToken: (jwt: string) => void;
  removeToken: () => void;
}

export const useUserStore = create<
  UserState,
  SetState<UserState>,
  GetState<UserState>,
  StoreApiWithPersist<UserState>
>(
  persist(
    devtools((set) => ({
      token: null,
      setToken: (newToken: string) => set(() => ({token: newToken})),
      removeToken: () => set(() => ({token: null})),
    })),
    {
      name: storeNames.userStore,
      getStorage: () => localStorage,
    }
  )
);

interface StorageWrapper<T> {
  state: T;
  version: number;
}

export function GetTokenFromStorage(): string | null{
  let userSateStr = localStorage.getItem(storeNames.userStore);

  if (!userSateStr) {
    return null;
  }

  let userState = JSON.parse(userSateStr) as StorageWrapper<UserState>;

  if (!userState) {
    return null;
  }

  return userState.state.token;
}