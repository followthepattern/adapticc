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
  jwt: string | null;
  setJwt: (jwt: string) => void;
}

export const useUserStore = create<
  UserState,
  SetState<UserState>,
  GetState<UserState>,
  StoreApiWithPersist<UserState>
>(
  persist(
    devtools((set) => ({
      jwt: null,
      setJwt: (newJwt: string) => set(() => ({jwt: newJwt}))
    })),
    {
      name: storeNames.userStore,
      getStorage: () => localStorage,
    }
  )
);
